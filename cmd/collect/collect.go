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
)

func init() {

	flag.StringVar(&dbConfig.DatabaseName, "name", "stock", "database name")
	flag.StringVar(&dbConfig.Host, "host", "127.0.0.1", "host address")
	flag.StringVar(&dbConfig.User, "u", "root", "user for database")
	flag.StringVar(&dbConfig.Password, "p", "123456", "password")
	flag.IntVar(&dbConfig.Port, "port", 3306, "port for database")
	flag.BoolVar(&debug, "d", false, "debug mode")
	flag.StringVar(&logFileName, "log", "1.log", "log file path")
	flag.Parse()
	stat.SetDebug(debug)
	stat.SetLogFile(logFileName)
	log = stat.GetLogger()
}

func main() {
	collectCode()
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
	log.Print(r)
	e = collector.SaveStockCode(r)
	if e != nil {
		log.Print(e)
		return
	}
	log.Println("save success")
}

func collectPrice() {
	log.Print("start collect price data")
	stockCodeList, e := stat.LoadStockCode(dbConfig)
	if e != nil {
		log.Println(e)
		return
	}
	log.Printf("load stock code list success, size: %d", len(stockCodeList))
	collector, e := stat.NewCollector(stat.STOCK_TYPE, dbConfig)
	if e != nil {
		log.Println(e)
		return
	}
	log.Println("create collector success")
	codes := make([]string, 0)
	for _, v := range stockCodeList {
		codes = append(codes, v.Code)
	}
	r, e := collector.FetchStockPrice(codes...)
	if e != nil {
		log.Println(e)
		return
	}
	for _, stockPrice := range r {
		e = collector.SaveStockPrice(stockPrice)
		if e != nil {
			log.Println(e)
		}
	}
	log.Print("finished!")
}