package main

import (
	stat "github.com/juxuny/statistics"
)

func clear() {
	if !stat.CheckDate(date) {
		log.Print("invaild date")
		return
	}
	log.Print("start clear...")
	log.Print("date: ", date)

	//清理股价数据
	codeList, e := stat.LoadStockCode(dbConfig)
	if e != nil {
		log.Panic(e)
	}
	for _, stockCode := range codeList {
		log.Print("clear: ", stockCode.Code)
		e = stat.ClearOneDay(dbConfig, stockCode.Code, date)
		if e != nil {
			log.Print(e)
			continue
		}
	}

	//清理大盘指数
	indexes, e := stat.LoadMarketIndexes(dbConfig)
	if e != nil {
		log.Panic(e)
	}
	for _, index := range indexes {
		e = stat.ClearOneDayIndex(dbConfig, index.Code, date)
		if e != nil {
			log.Print(e)
			continue
		}
	}
}
