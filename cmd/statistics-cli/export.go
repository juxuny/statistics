package main

import (
	stat "github.com/juxuny/statistics"
	"os"
)

func export() {
	if !stat.CheckDate(date) {
		log.Print("invaild date")
		return
	}
	log.Print("start export...")
	codeList, e := stat.LoadStockCode(dbConfig)
	if e != nil {
		log.Panic(e)
	}
	if _, err := os.Stat(out + string(os.PathSeparator) + date); os.IsNotExist(err) {
		e = os.MkdirAll(out + string(os.PathSeparator) + date, 0755)
		if e != nil {
			log.Panic(e)
		}
	}
	for _, stockCode := range codeList {
		log.Print("fetch: ", stockCode.Code)
		data, e := stat.GetOneDay(dbConfig, stockCode.Code, date)
		if e != nil {
			log.Print(e)
			continue
		}
		e = data.SaveToCSV(out + string(os.PathSeparator) + date + string(os.PathSeparator) + stockCode.Code + ".csv")
		if e != nil {
			log.Print(e)
			continue
		}
	}

	//导出大盘指数
	indexOutDir := out + string(os.PathSeparator) + date + string(os.PathSeparator) + "index"
	if _, err := os.Stat(indexOutDir); os.IsNotExist(err) {
		e = os.MkdirAll(indexOutDir, 0755)
		if e != nil {
			log.Panic(e)
		}
	}
	indexes, e := stat.LoadMarketIndexes(dbConfig)
	if e != nil {
		log.Panic(e)
	}
	for _, index := range indexes {
		indexData, e := stat.GetOneDayIndex(dbConfig, index.Code, date)
		if e != nil {
			log.Print(e)
			continue
		}
		e = indexData.SaveToCSV(indexOutDir + string(os.PathSeparator) + index.Code + ".csv")
		if e != nil {
			log.Print(e)
			continue
		}
	}
}