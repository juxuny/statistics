//下载历史数据（数据源：搜狐）
//  api-link: http://q.stock.sohu.com/hisHq?code=cn_600000&start=19980930&end=20180409&stat=1&order=D&period=d&rt=json
package main

import (
	stat "github.com/juxuny/statistics"
	"flag"
	l "log"
)

var dbConfig stat.DBConfig
var (
	log *l.Logger
	debug bool
	logFileName string
	verbose bool
	start, end string
)

func init() {
	flag.StringVar(&dbConfig.DatabaseName, "name", "stock_sohu", "database name")
	flag.StringVar(&dbConfig.Host, "host", "127.0.0.1", "host address")
	flag.StringVar(&dbConfig.User, "u", "root", "user for database")
	flag.StringVar(&dbConfig.Password, "p", "123456", "password")
	flag.IntVar(&dbConfig.Port, "port", 3306, "port for database")

	flag.BoolVar(&debug, "d", true, "debug mode")
	flag.StringVar(&logFileName, "log", "1.log", "log file path")
	flag.BoolVar(&verbose, "v", true, "verbose")
	flag.StringVar(&start, "start", "19980101", "start time")
	flag.StringVar(&end, "end", "20180411", "end time")

	flag.Parse()
	stat.SetDebug(debug)
	stat.SetLogFile(logFileName)
	log = stat.GetLogger()
}

func main() {
	log.Println("start download ...")
	codeList, e := stat.LoadStockCode(dbConfig)
	if e != nil {
		log.Panicln(e)
	}
	log.Println(codeList)
	stat.SetDebug(debug)
	collector, e := stat.NewSohuCollector(dbConfig, true)
	for _, stockCode := range codeList {
		c := stat.ConvertSinaCodeToSohuCode(stockCode.Code)
		log.Println("fatch data: ", c)
		m, e := collector.FetchStockPriceDuration(stat.NumberFilter(start), stat.NumberFilter(end), c)
		if e != nil {
			log.Println(e)
			continue
		}
		e = collector.SaveStockPrice(m[c]...)
		if e != nil {
			log.Println(e)
		}
	}
}