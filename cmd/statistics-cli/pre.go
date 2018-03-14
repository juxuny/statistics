package main

import (
	stat "github.com/juxuny/statistics"
	"os"
	"encoding/csv"
	"time"
	"fmt"
)

var _STANDARDIZATION_HEADER = []string{
	"stock_code", "name", "date", "time", "current_price", "open_price","yesterday_price", "max", "min",
	"buy_price_1", "buy_price_2", "buy_price_3", "buy_price_4", "buy_price_5",
	"buy_1", "buy_2", "buy_3", "buy_4", "buy_5",
	"sell_price_1", "sell_price_2", "sell_price_3", "sell_price_4", "sell_price_5",
	"sell_1", "sell_2", "sell_3", "sell_4", "sell_5",
	"deal", "deal_price",
	"mean_price", "mean_sell", "mean_buy", "mean_deal_volume", "mean_deal_price",
	"sd_price", "sd_sell", "sd_buy", "sd_deal_volume", "sd_deal_price",
	"week_0", "week_1", "week_2", "week_3", "week_4", "week_5", "week_6",
}

var timeMap map[string]int

func generateTimeMap() (m map[string]int) {
	m = make(map[string]int)
	t, e := time.Parse("15:04:05", "09:25:00")
	if e != nil {
		log.Panic(e)
	}
	ii := 0
	for t.Format("15:04") <= "16:00" {
		tmp := t.Format("15:04")
		m[tmp] = ii
		ii ++
		t = t.Add(time.Minute)
	}
	return
}

func initTimeMap() {
	timeMap = generateTimeMap()
	var header []string
	iterator := 0
	for range timeMap {
		header = append(header, fmt.Sprintf("t_%d", iterator))
		iterator ++
	}
	_STANDARDIZATION_HEADER = stat.MergeStrings(_STANDARDIZATION_HEADER, header)
}

func timeToOneHotVector(timeStr string) (oneHot stat.OneHot, e error) {
	index, b := timeMap[timeStr[:5]]
	if b {
		oneHot = stat.NewOneHot(index, len(timeMap))
		return
	}
	e = fmt.Errorf("invalid time")
	return
}


//数据导出预处理
func pre() {
	initTimeMap()
	if !stat.CheckDate(start) || !stat.CheckDate(end) {
		log.Print("invalid date")
		return
	}
	log.Print("start export...")
	codeList, e := stat.LoadStockCode(dbConfig)
	if e != nil {
		log.Panic(e)
	}
	if _, err := os.Stat(out); os.IsNotExist(err) {
		e = os.MkdirAll(out, 0755)
		if e != nil {
			log.Panic(e)
		}
	}
	for _, stockCode := range codeList {
		if code != "" && stockCode.Code != code {
			continue
		}
		log.Print("fetch: ", stockCode.Code)
		data, e := stat.GetDurationStock(dbConfig, stockCode.Code, start, end, 0, 50000)
		if e != nil {
			log.Print(e)
			continue
		}
		dataStandardization, mean, sd, e := data.Standardization()
		if e != nil {
			log.Println(e)
			continue
		}

		//write csv
		f, e := os.Create(out + string(os.PathSeparator) + stockCode.Code + ".csv")
		if e != nil {
			log.Println(e)
			continue
		}
		w := csv.NewWriter(f)
		w.Write(_STANDARDIZATION_HEADER[2:])
		for _, v := range dataStandardization {
			t, e := time.Parse("2006-01-02 15:04:05", v.Date + " " + v.Time)
			if e != nil {
				log.Println(e)
				continue
			}
			weekOneHot := stat.NewOneHot(int(t.Weekday()), 7)
			timeOneHot, e := timeToOneHotVector(t.Format("15:04"))
			if e != nil {
				log.Println(e)
				continue
			}
			row := v.ToStrings()[2:]
			row = stat.MergeStrings(row, mean.ToStrings(), sd.ToStrings(), weekOneHot.ToStrings(), timeOneHot.ToStrings())
			e = w.Write(row)
			if e != nil {
				log.Println(e)
				continue
			}
		}
		f.Close()
	}
}
