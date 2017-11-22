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
	start, end string
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
	flag.StringVar(&start, "start", "09:00", "time to start")
	flag.StringVar(&end, "end", "16:00", "time to exit")
	flag.Parse()
	stat.SetDebug(debug)
	stat.SetLogFile(logFileName)
	log = stat.GetLogger()
}

func main() {
	for {
		current := time.Now().Format("15:04")
		log.Print("current: ", current)
		if current >= start && current <= end {
			go collectPrice()
			go collectMarketIndexes()
		}
		time.Sleep(time.Minute)
	}
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
	collector := stat.NewCollector(dbConfig)
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
		if verbose {
			log.Print("saved: ", stockPrice.StockCode)
		}
	}
}

func collectMarketIndexes() {
	collector := stat.NewCollector(dbConfig)
	marketIndexes, e := stat.LoadMarketIndexes(dbConfig)
	if e != nil {
		log.Print(e)
		return
	}
	codes := make([]string, 0)
	for _, v := range marketIndexes {
		codes = append(codes, v.Code)
	}
	r, e := collector.FetchMarketIndexes(codes...)
	if e != nil {
		log.Print(e)
		return
	}
	var tmp []stat.MarketIndexInfo
	for _, v := range r {
		tmp = append(tmp, v)
	}
	e = collector.SaveMarketIndexesData(tmp...)
	if e != nil {
		log.Print(e)
		return
	}
	log.Print("collect market index data finished.")
}