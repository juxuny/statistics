package main

import (
	stat "github.com/juxuny/statistics"
	"fmt"
	"flag"
	"time"
)

var dbConfig stat.DBConfig

var debug bool

func init() {

	flag.StringVar(&dbConfig.DatabaseName, "name", "stock", "database name")
	flag.StringVar(&dbConfig.Host, "host", "127.0.0.1", "host address")
	flag.StringVar(&dbConfig.User, "u", "root", "user for database")
	flag.StringVar(&dbConfig.Password, "p", "123456", "password")
	flag.IntVar(&dbConfig.Port, "port", 3306, "port for database")
	flag.BoolVar(&debug, "d", false, "debug mode")
	flag.Parse()
	stat.SetDebug(debug)
}

func main() {
	for {

		collector, e := stat.NewCollector(stat.STOCK_TYPE, dbConfig)
		if e != nil {
			panic(e)
		}
		r, e := collector.FetchStockCode()
		if e != nil {
			panic(e)
		}
		fmt.Print(r)
		e = collector.Save(r)
		if e != nil {
			panic(e)
		}
		fmt.Println("save success")
		time.Sleep(time.Minute * 10)
	}
}
