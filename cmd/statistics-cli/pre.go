package main

import (
	stat "github.com/juxuny/statistics"
	"os"
	"encoding/csv"
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
}


//数据导出预处理
func pre() {
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
			row := v.ToStrings()[2:]
			row = stat.MergeStrings(row, mean.ToStrings(), sd.ToStrings())
			e = w.Write(row)
			if e != nil {
				log.Println(e)
				continue
			}
		}
		f.Close()
	}
}
