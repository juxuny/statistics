package statistics

import (
	"strings"
	"fmt"
	"strconv"
)

type StockCode struct {
	Type string
	//stock code
	Code string
	//stock name
	Name string
}

type StockPrice struct {
	StockCode string
	Name string
	Date string
	Time string
	CurrentPrice float64
	OpenPrice float64
	YesterdayPrice float64
	Max float64
	Min float64
	BuyPrice [5]float64
	Buy [5]float64
	SellPrice [5]float64
	Sell [5]float64
	Deal float64
	DealPrice float64
}


//解释从新浪获取到的原始数据
//  e.g var hq_str_sh601006="大秦铁路,8.800,8.790,9.030,9.060,8.700,9.030,9.040,109367311,971205411.000,73530,9.030,895919,9.020,555200,9.010,29700,9.000,137700,8.990,256000,9.040,772531,9.050,1108457,9.060,269100,9.070,248000,9.080,2017-11-17,15:00:00,00";
//  0：”大秦铁路”，股票名字；
//	1：”27.55″，今日开盘价；
//	2：”27.25″，昨日收盘价；
//	3：”26.91″，当前价格；
//	4：”27.55″，今日最高价；
//	5：”26.20″，今日最低价；
//	6：”26.91″，竞买价，即“买一”报价；
//	7：”26.92″，竞卖价，即“卖一”报价；
//	8：”22114263″，成交的股票数，由于股票交易以一百股为基本单位，所以在使用时，通常把该值除以一百；
//	9：”589824680″，成交金额，单位为“元”，为了一目了然，通常以“万元”为成交金额的单位，所以通常把该值除以一万；
//	10：”4695″，“买一”申请4695股，即47手；
//	11：”26.91″，“买一”报价；
//	12：”57590″，“买二”
//	13：”26.90″，“买二”
//	14：”14700″，“买三”
//	15：”26.89″，“买三”
//	16：”14300″，“买四”
//	17：”26.88″，“买四”
//	18：”15100″，“买五”
//	19：”26.87″，“买五”
//	20：”3100″，“卖一”申报3100股，即31手；
//	21：”26.92″，“卖一”报价
//	(22, 23), (24, 25), (26,27), (28, 29)分别为“卖二”至“卖四的情况”
//	30：”2008-01-11″，日期；
//	31：”15:05:32″，时间；
//  32: "00", 未知
func ParseStockPrice(stockCode string, resp string) (r StockPrice, e error) {
	ii := strings.Index(resp, "\"")
	if ii == -1 {
		e = fmt.Errorf("invalid response data")
		return
	}
	resp = resp[ii+1:]
	ii = strings.Index(resp, "\"")
	if ii == -1 {
		e = fmt.Errorf("invalid response data")
		return
	}
	resp = resp[: ii]
	s := strings.Split(resp, ",")
	if len(s) != 33 {
		e = fmt.Errorf("invalid response data, filed length: %d", len(s))
		return
	}
	r.StockCode = stockCode
	r.Name = s[0]
	r.OpenPrice, e = strconv.ParseFloat(s[1], 64)
	r.YesterdayPrice, e = strconv.ParseFloat(s[2], 64)
	r.CurrentPrice, e = strconv.ParseFloat(s[3], 64)
	r.Max, e = strconv.ParseFloat(s[4], 64)
	r.Min, e = strconv.ParseFloat(s[5], 64)
	//6
	//7
	r.Deal, e = strconv.ParseFloat(s[8], 64)
	r.DealPrice, e = strconv.ParseFloat(s[9], 64)
	r.Buy[0], e = strconv.ParseFloat(s[10], 64)
	r.BuyPrice[0], e = strconv.ParseFloat(s[11], 64)
	r.Buy[1], e = strconv.ParseFloat(s[12], 64)
	r.BuyPrice[1], e = strconv.ParseFloat(s[13], 64)
	r.Buy[2], e = strconv.ParseFloat(s[14], 64)
	r.BuyPrice[2], e = strconv.ParseFloat(s[15], 64)
	r.Buy[3], e = strconv.ParseFloat(s[16], 64)
	r.BuyPrice[3], e = strconv.ParseFloat(s[17], 64)
	r.Buy[4], e = strconv.ParseFloat(s[18], 64)
	r.BuyPrice[4], e = strconv.ParseFloat(s[19], 64)
	r.Sell[0], e = strconv.ParseFloat(s[20], 64)
	r.SellPrice[0], e = strconv.ParseFloat(s[21], 64)
	r.Sell[1], e = strconv.ParseFloat(s[22], 64)
	r.SellPrice[2], e = strconv.ParseFloat(s[23], 64)
	r.Sell[2], e = strconv.ParseFloat(s[24], 64)
	r.SellPrice[1], e = strconv.ParseFloat(s[25], 64)
	r.Sell[3], e = strconv.ParseFloat(s[26], 64)
	r.SellPrice[3], e = strconv.ParseFloat(s[27], 64)
	r.Sell[4], e = strconv.ParseFloat(s[28], 64)
	r.SellPrice[4], e = strconv.ParseFloat(s[29], 64)
	r.Date = s[30]
	r.Time = s[31]
	return
}

type Collector interface {
	FetchStockCode() (r []StockCode, e error)
	SaveStockCode(r []StockCode) (e error)
	//一个接一个地获取
	FetchStockPrice(stockCode ...string) (r map[string]StockPrice, e error)
	FetchStockPrices(stockCode ...string) (r map[string]StockPrice, e error)
	SaveStockPrice(price ...StockPrice) (e error)
	init(stockCode string) (e error)
}


//从数据库加载股票代码
func LoadStockCode(config DBConfig) (r []StockCode, e error) {
	db, e := NewConnection(config)
	if e != nil {
		return
	}
	rs, e := db.Query("SELECT code, name, type FROM stock_code")
	if e != nil {
		return
	}
	for rs.Next() {
		t := StockCode{}
		e := rs.Scan(&t.Code, &t.Name, &t.Type)
		if e != nil {
			log.Print(e)
			continue
		}
		r = append(r, t)
	}
	return
}