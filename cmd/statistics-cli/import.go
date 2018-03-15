package main

import (
	stat "github.com/juxuny/statistics"
	"os"
	"encoding/csv"
)

func loadCSV(code, path string) (ret [][]string, e error) {
	fileName := path + string(os.PathSeparator) + code + ".csv"
	f, e := os.Open(fileName)
	if e != nil {
		return
	}
	defer f.Close()
	r := csv.NewReader(f)
	ret, e = r.ReadAll()
	return
}

func importDir() {
	if dir == "" {
		log.Println("argument 'dir' cannot be empty")
		return
	}
	if !stat.CheckDate(date) {
		log.Println("invalid date: ", date)
		return
	}
	log.Println("start import")
	codeList, e := stat.LoadStockCode(dbConfig)
	if e != nil {
		log.Panic(e)
	}
	count := 0
	for _, stockCode := range codeList {
		if code != "" && stockCode.Code != code {
			continue
		}
		ret, e := loadCSV(stockCode.Code, dir + string(os.PathSeparator) + date)
		if e != nil {
			log.Println(e)
			continue
		}
		if len(ret) <= 1 {
			log.Println("no data, ", stockCode.Code)
			continue
		}
		ret = ret[1:] //remove the first row(header row)
		stockPriceList, e := stat.ParseStockPriceListFromStrings(stockCode.Code, stockCode.Name, ret)
		if e != nil {
			log.Println(e)
			continue
		}
		log.Println("save stock data,", stockCode.Code, date)
		collector := stat.NewCollector(dbConfig)
		e = collector.SaveStockPrice(stockPriceList...)
		if e != nil {
			log.Println(e)
			continue
		}
		count ++
	}
	log.Println("import number of stock:", count)
}
