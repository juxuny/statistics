//从凤凰网获取所有A,B股的股票代码
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

	flag.Parse()
	stat.SetDebug(debug)
	stat.SetLogFile(logFileName)
	log = stat.GetLogger()
}

func main() {
	collector := stat.NewFCollector(dbConfig)
	r ,e := collector.FetchStockCode()
	if e != nil {
		log.Panic(e)
		return
	}
	e = collector.SaveStockCode(r)
	if e != nil {
		log.Print(e)
		return
	}
	log.Print("finished")
}
