//用于收集每分钟的股票数据
package main

import (
	stat "github.com/juxuny/statistics"
	"flag"
	"time"
)

var dbConfig stat.DBConfig

var (
	debug bool
	logFileName string
	log = stat.GetLogger()
	verbose bool
	initStockCode bool
)

func init() {

	flag.StringVar(&dbConfig.DatabaseName, "name", "stock", "database name")
	flag.StringVar(&dbConfig.Host, "host", "127.0.0.1", "host address")
	flag.StringVar(&dbConfig.User, "u", "root", "user for database")
	flag.StringVar(&dbConfig.Password, "p", "123456", "password")
	flag.IntVar(&dbConfig.Port, "port", 3306, "port for database")
	flag.BoolVar(&debug, "d", false, "debug mode")
	flag.StringVar(&logFileName, "log", "1.log", "log file path")
	flag.BoolVar(&verbose, "v", false, "verbose")
	flag.BoolVar(&initStockCode, "i", true, "update table named stock_code")
	flag.Parse()
	stat.SetDebug(debug)
	stat.SetLogFile(logFileName)
	log = stat.GetLogger()
}

func main() {
	if initStockCode {
		collectCode()
	}
	go func () {
		for {
			time.Sleep(12 * time.Hour)
			go collectCode()
		}
	}()
	for {
		go collectPrice()
		time.Sleep(time.Minute)
	}
}


func collectCode() {
	collector, e := stat.NewCollector(stat.STOCK_TYPE, dbConfig)
	if e != nil {
		log.Print(e)
	}
	r, e := collector.FetchStockCode()
	if e != nil {
		log.Print(e)
	}
	e = collector.SaveStockCode(r)
	if e != nil {
		log.Print(e)
		return
	}
	log.Println("save stock code success")
}

func collectPrice() {
	log.Print("start collect price data")
	stockCodeList, e := stat.LoadStockCode(dbConfig)
	if e != nil {
		log.Println(e)
		return
	}
	log.Printf("load stock code list success, size: %d", len(stockCodeList))

	var codeList []string
	for _, v := range stockCodeList {
		if len(codeList) >= stat.BATCH_SIZE {
			handle(codeList)
			codeList = make([]string, 0)
		}
		codeList = append(codeList, v.Code)
	}
	handle(codeList)
	log.Print("finished")
}

func handle(codeList []string) {
	collector, e := stat.NewCollector(stat.STOCK_TYPE, dbConfig)
	if e != nil {
		log.Print(e)
	}
	r, e := collector.FetchStockPrices(codeList...)
	if e != nil {
		log.Print(e)
		return
	}
	for _, stockPrice := range r {
		e = collector.SaveStockPrice(stockPrice)
		if e != nil {
			log.Print(e)
			continue
		}
		log.Print("saved: ", stockPrice.StockCode)
	}
}