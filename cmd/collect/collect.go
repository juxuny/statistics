//用于收集每分钟的股票数据
package main

import (
	stat "github.com/juxuny/statistics"
	"flag"
	"time"
	"log"
)

var dbConfig stat.DBConfig

var debug bool

func init() {

	flag.StringVar(&dbConfig.DatabaseName, "name", "stock", "database name")
	flag.StringVar(&dbConfig.Host, "host", "127.0.0.1", "host address")
	flag.StringVar(&dbConfig.User, "u", "root", "user for database")
	flag.StringVar(&dbConfig.Password, "p", "123456", "password")
	flag.IntVar(&dbConfig.Port, "port", 3306, "port for database")
	flag.BoolVar(&debug, "d", true, "debug mode")
	flag.Parse()
	stat.SetDebug(debug)
}

func main() {
	go func () {
		for {
			go collectCode()
			time.Sleep(10 * time.Minute)
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
		panic(e)
	}
	r, e := collector.FetchStockCode()
	if e != nil {
		panic(e)
	}
	log.Print(r)
	e = collector.SaveStockCode(r)
	if e != nil {
		panic(e)
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