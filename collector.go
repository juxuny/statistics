package statistics

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"time"
	"strings"
	"github.com/opesun/goquery"
)

const (
	//A股 - 股票列表
	STOCK_CODE_API = "http://www.ctxalgo.com/api/stocks"

	//股票类型
	STOCK_TYPE_A = "A"

	BATCH_SIZE = 300
)

const CREATE_STOCK_TABLE_TEMPLATE = `
CREATE TABLE IF NOT EXISTS %s
(
  date            VARCHAR(10) NOT NULL
  COMMENT '日期',
  time            VARCHAR(8)  NOT NULL
  COMMENT '时间',
  current_price   DOUBLE      NULL
  COMMENT '当前价',
  open_price      DOUBLE      NULL
  COMMENT '今日开盘价',
  yesterday_close DOUBLE      NULL
  COMMENT '昨日收盘价',
  max             DOUBLE      NULL
  COMMENT '今日最高价',
  min             DOUBLE      NULL
  COMMENT '今日最低价',
  buy_1           DOUBLE      NULL
  COMMENT '”4695″，“买一”申请4695股，即47手；',
  buy_2           DOUBLE      NULL,
  buy_3           DOUBLE      NULL,
  buy_4           DOUBLE      NULL,
  buy_5           DOUBLE      NULL,
  buy_price_1     DOUBLE      NULL
  COMMENT '“买一”报价',
  buy_price_2     DOUBLE      NULL,
  buy_price_3     DOUBLE      NULL,
  buy_price_4     DOUBLE      NULL,
  buy_price_5     DOUBLE      NULL,
  sell_1          DOUBLE      NULL
  COMMENT '”3100″，“卖一”申报3100股，即31手；',
  sell_2          DOUBLE      NULL,
  sell_3          DOUBLE      NULL,
  sell_4          DOUBLE      NULL,
  sell_5          DOUBLE      NULL,
  sell_price_1    DOUBLE      NULL,
  sell_price_2    DOUBLE      NULL,
  sell_price_3    DOUBLE      NULL,
  sell_price_4    DOUBLE      NULL,
  sell_price_5    DOUBLE      NULL,
  deal            DOUBLE      NULL
  COMMENT '成交的股票数，由于股票交易以一百股为基本单位，所以在使用时，通常把该值除以一百；',
  deal_price      DOUBLE      NULL
  COMMENT '成交金额，单位为“元”，为了一目了然，通常以“万元”为成交金额的单位，所以通常把该值除以一万；'
)
  COMMENT '数据表模板';
`

const CREATE_PRIMARY_KEY = `ALTER TABLE  %s ADD PRIMARY KEY (date, time);`


type CollectorImpl struct {
	Type string
	config DBConfig
	Prefix string
}

func NewCollector(dbConfig DBConfig) (r Collector) {
	c := &CollectorImpl{Type: STOCK_TYPE_A, Prefix: "sina_"}
	c.config = dbConfig
	r = c
	return
}

func (t *CollectorImpl) FetchStockCode()(r []StockCode, e error) {
	url := STOCK_CODE_API
	resp, e := http.Get(url)
	if resp.StatusCode != 200 {
		return r, fmt.Errorf("HTTP ERROR: %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	log.Print(resp.Header)
	log.Print("fetch stock code finished.")
	data, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return
	}
	m := make(map[string]string)
	e = json.Unmarshal(data, &m)
	if e != nil {
		return
	}
	for k, v := range m {
		r = append(r, StockCode{Type: t.Type, Code: k, Name: v})
	}
	return
}

func (t *CollectorImpl) SaveStockCode(r []StockCode) (e error) {
	db, e := NewConnection(t.config)
	if e != nil {
		log.Print(e)
		return
	}
	defer db.Close()
	for _, v := range r {
		tx, e := db.Begin()
		if e != nil {
			log.Print(e)
			continue
		}
		_, e = tx.Exec("INSERT INTO stock_code (code, name, type) SELECT ?, ?, ? FROM DUAL WHERE NOT EXISTS (SELECT id FROM stock_code WHERE code = ?)", v.Code, v.Name, v.Type, v.Code)
		if e != nil {
			log.Print(e)
			tx.Rollback()
			continue
		}
		_, e = db.Exec(fmt.Sprintf(CREATE_STOCK_TABLE_TEMPLATE, t.Prefix + v.Code))
		if e != nil {
			log.Print(e)
			tx.Rollback()
			continue
		}
		_, e = db.Exec(fmt.Sprintf(CREATE_PRIMARY_KEY, t.Prefix + v.Code))
		if e != nil {
			tx.Rollback()
			continue
		}
		e = tx.Commit()
	}
	return
}

//按股票代码创建表
func (t *CollectorImpl) init(stockCode string) (e error) {
	return
}

func (t *CollectorImpl) FetchStockPrice(stockCode ...string) (r map[string]StockPrice, e error) {
	r = make(map[string]StockPrice)
	for _, code := range stockCode {
		//log.Print("fetch: ", code)
		client := http.DefaultClient
		client.Timeout = time.Duration(30 * time.Second)
		resp, e := client.Get("http://hq.sinajs.cn/list=" + code)
		//resp, e := http.Get()
		if e != nil {
			log.Print(e)
			continue
		}
		//log.Print("fetch success, ", code)
		if resp.StatusCode != 200 {
			log.Printf("fetch data failed, stock code: %s", code)
			continue
		}
		data, e := ioutil.ReadAll(resp.Body)
		if e != nil {
			log.Print(e)
			continue
		}
		if !resp.Close {
			resp.Body.Close()
		}
		s, e := GBK_UTF8(string(data))
		if e != nil {
			continue
		}
		tmp, e := ParseStockPrice(code, s)
		if e != nil {
			log.Print(e)
			continue
		}
		r[code] = tmp
	}
	return
}


func (t *CollectorImpl) SaveStockPrice(price ...StockPrice) (e error) {
	db, e := NewConnection(t.config)
	if e != nil {
		return
	}
	defer db.Close()
	for _, p := range price {
		table := t.Prefix + p.StockCode
		sql := `INSERT IGNORE INTO %s (date, time, current_price, open_price, yesterday_close, max, min, buy_1, buy_2, buy_3, buy_4, buy_5, buy_price_1, buy_price_2, buy_price_3, buy_price_4, buy_price_5, sell_1, sell_2, sell_3, sell_4, sell_5, sell_price_1, sell_price_2, sell_price_3, sell_price_4, sell_price_5, deal, deal_price)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`
		//log.Print("save stock price")
		_, e = db.Exec(fmt.Sprintf(sql, table), p.Date, p.Time, p.CurrentPrice, p.OpenPrice, p.YesterdayPrice, p.Max, p.Min, p.Buy[0], p.Buy[1], p.Buy[2], p.Buy[3], p.Buy[4], p.BuyPrice[0], p.BuyPrice[1], p.BuyPrice[2], p.BuyPrice[3], p.BuyPrice[4], p.Sell[0], p.Sell[1], p.Sell[2], p.Sell[3], p.Sell[4], p.SellPrice[0], p.SellPrice[1], p.SellPrice[2], p.SellPrice[3], p.SellPrice[4], p.Deal, p.DealPrice, )
			if e != nil {
				log.Print(e)
			}
	}
	return
}

func (t *CollectorImpl) FetchStockPrices(stockCode ...string) (r map[string]StockPrice, e error) {
	if len(stockCode) > BATCH_SIZE {
		e = fmt.Errorf("the number of code is too large: %d", len(stockCode))
		return
	}
	r = make(map[string]StockPrice)
	q := ""
	for _, c := range stockCode {
		if q != "" {
			q += ","
		}
		q += strings.Trim(c, " ")
	}
	resp, e := http.Get("http://hq.sinajs.cn/list=" + q)
	if e != nil {
		return
	}
	defer resp.Body.Close()
	data, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return
	}
	ss, e := GBK_UTF8(string(data))
	if e != nil {
		return
	}
	ss = strings.Trim(ss, "\n")
	vars := strings.Split(ss, "\n")
	for i, s := range vars {
		tmp, e := ParseStockPrice(stockCode[i], s)
		if e != nil {
			log.Print(e)
			continue
		}
		r[stockCode[i]] = tmp
	}
	e = nil
	return
}


//抓取凤凰网的股票代码
type FCollector struct {
	CollectorImpl
}

func NewFCollector(config DBConfig) (r Collector) {
	c := &FCollector{}
	c.config = config
	c.Prefix = "sina_"
	r = c
	return
}

func (t *FCollector) FetchStockCode() (r []StockCode, e error) {
	ts := []string{
		"A", "ha", "sh", // A 股, 沪A, (其中sh是对应新浪数据库里的前缀)
		"A", "sa", "sz", //A股，深A,(其中sz是对应新浪数据库里的前缀)
		"B", "hb", "sh",
		"B", "sb", "sz",
		"A", "gem", "sz",//A股，创业板
	}
	for j := 0; j < len(ts); j+=3 {
		ret, e := t.fetch(ts[j], ts[j+1], ts[j+2])
		if e != nil {
			log.Print(e)
			continue
		}
		for i := 0; i < len(ret); i++ {
			r = append(r, ret[i])
		}
	}
	return
}

// stockType: A 表示A股
// stockClass: sh
func (t *FCollector) fetch(stockType, stockClass, sinaPrefix string) (r []StockCode, e error) {
	log.Print("fetch ", stockType, stockClass)
	var tt string
	switch stockType {
	case "A":
		tt = "stock_a"
	case "B":
		tt = "stock_b"
	default:
		e = fmt.Errorf("unknown type: %s", stockType)
		return
	}
	resp, e := goquery.ParseUrl(fmt.Sprintf("http://app.finance.ifeng.com/hq/list.php?type=%s&class=%s", strings.ToLower(tt), strings.ToLower(stockClass)))
	if e != nil {
		return
	}
	ret := resp.Find(".result ul li")
	for i := 0; i < ret.Length(); i++ {
		tmp := strings.Trim(ret.Eq(i).Text(), " \n")
		start := strings.Index(tmp, "(")
		end := strings.Index(tmp, ")")
		if start > 0 && end > 0 && start < end {
			stockCode := StockCode{}
			stockCode.Type = stockType
			stockCode.Name = tmp[:start]
				stockCode.Code = sinaPrefix + tmp[start+1:end]
			r = append(r, stockCode)
		}
	}
	return
}