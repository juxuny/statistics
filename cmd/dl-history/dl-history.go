package main

import "flag"
import l "log"
import (
	stat "github.com/juxuny/statistics"
)

var (
	log *l.Logger
	debug bool
	logFileName string
	verbose bool
	start, end string
	outDir string
)

func init() {
	flag.BoolVar(&debug, "d", true, "debug mode")
	flag.StringVar(&logFileName, "log", "1.log", "log file path")
	flag.BoolVar(&verbose, "v", true, "verbose")
	flag.StringVar(&start, "start", "19980101", "start time")
	flag.StringVar(&end, "end", "20180411", "end time")
	flag.StringVar(&outDir, "o", "./data", "output directory path")

	flag.Parse()
	stat.SetDebug(debug)
	_ = stat.SetLogFile(logFileName)
	log = stat.GetLogger()
}


func main() {
	if err := stat.TouchDir(outDir); err != nil {
		panic(err)
	}

	col := stat.NewHistoryCollector()
	stockList, err := col.FetchStockCode()
	if err != nil {
		panic(err)
	}
	for _, item := range stockList {
		log.Println(item.Code, item.Name)
	}
}