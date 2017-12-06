package main

import (
	stat "github.com/juxuny/statistics"
	"flag"
	l "log"
)

var (
	dbConfig = stat.NewDefaultDBConfig()
	log *l.Logger
	debug bool
	logFileName string
	verbose bool
	mode string

	//要导出的日期
	date string
	//股票代码, e.g sz300715
	code string
	//csv目录
	out string
)

func init() {
	dbConfig.InitTable = false

	//指令
	flag.StringVar(&mode, "m", "help", "command")

	//数据库
	flag.StringVar(&dbConfig.DatabaseName, "name", "stock", "database name")
	flag.StringVar(&dbConfig.Host, "host", "127.0.0.1", "host address")
	flag.StringVar(&dbConfig.User, "u", "root", "user for database")
	flag.StringVar(&dbConfig.Password, "p", "123456", "password")
	flag.IntVar(&dbConfig.Port, "port", 3306, "port for database")

	//调试
	flag.BoolVar(&debug, "d", false, "debug mode")
	flag.StringVar(&logFileName, "log", "", "log file path")
	flag.BoolVar(&verbose, "v", false, "verbose")

	//数据参数
	flag.StringVar(&date, "date", "", "YYYY-MM-DD")
	flag.StringVar(&code, "code", "", "stock code, e.g sz300715")
	flag.StringVar(&out, "out", ".", "output dir")

	flag.Parse()
	stat.SetDebug(debug)
	stat.SetLogFile(logFileName)
	log = stat.GetLogger()
}

func main() {
	switch mode {
	case "export": {
		export()
	}
	default:
		flag.PrintDefaults()
	}
}
