//导出搜狐数据为CSV文件
package main

import (
	stat "github.com/juxuny/statistics"
	"flag"
	l "log"
	"database/sql"
	"os"
	"encoding/csv"
	"fmt"
)

var dbConfig stat.DBConfig
var (
	log *l.Logger
	debug bool
	logFileName string
	verbose bool
	start, end string
	out string
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
	flag.StringVar(&out, "out", "out", "output directory")

	flag.Parse()
	stat.SetDebug(debug)
	stat.SetLogFile(logFileName)
	log = stat.GetLogger()
}

func main() {
	if !stat.PathExist(out) {
		l.Printf("directory not found, '%s'\n", out)
		l.Println("creating directory")
		os.Mkdir(out, 0777)
	}
	l.Println("connecting ..")
	db, e := stat.NewConnection(dbConfig)
	if e != nil {
		l.Println(e)
		return
	}
	defer db.Close()
	l.Println("loading stock code")
	rs, e := db.Query("SELECT code, name, type FROM stock_code")
	if e != nil {
		l.Println(e)
		return
	}
	defer rs.Close()
	var item struct {
		Code sql.NullString
		Name sql.NullString
		Type sql.NullString
	}
	var stockCodeList []stat.StockCode
	for rs.Next() {
		e = rs.Scan(&item.Code, &item.Name, &item.Type)
		if e != nil {
			l.Println(e)
			break
		}
		stockCodeList = append(stockCodeList, stat.StockCode{Type: item.Type.String, Name: item.Name.String, Code: item.Code.String})
	}
	l.Println(stockCodeList)

	l.Println("export stock code")
	f, e := os.Create(out + string(os.PathSeparator) + "stock_code.csv")
	if e != nil {
		l.Println(e)
		return
	}
	defer f.Close()
	csvFile := csv.NewWriter(f)
	csvFile.Write([]string{"code", "name", "type"})
	for _, v := range stockCodeList {
		r := []string{v.Code, v.Name, v.Type}
		l.Println(r)
		csvFile.Write(r)
	}
	csvFile.Flush()
	l.Println("export stock code finished")
	l.Println("export stock data")
	for _, stockCode := range stockCodeList {
		code := stockCode.GetNumberCode()
		l.Println("fetch: ", code)
		rs, e := db.Query("SELECT date, current_price, open_price, max, min, deal, deal_price FROM sohu_cn_" + code + " ORDER BY date")
		if e != nil {
			continue
		}
		var item struct {
			Date sql.NullString
			CurrentPrice sql.NullFloat64
			OpenPrice sql.NullFloat64
			Max sql.NullFloat64
			Min sql.NullFloat64
			Deal sql.NullFloat64
			DealPrice sql.NullFloat64
		}
		var prices []stat.StockPrice
		for rs.Next() {
			rs.Scan(&item.Date, &item.CurrentPrice, &item.OpenPrice, &item.Max, &item.Min, &item.Deal, &item.DealPrice)
			prices = append(prices, stat.StockPrice{StockCode: code,
				Date: item.Date.String,
				CurrentPrice: item.CurrentPrice.Float64,
				OpenPrice: item.OpenPrice.Float64,
				Max: item.Max.Float64,
				Min: item.Min.Float64,
				Deal: item.Deal.Float64,
				DealPrice: item.DealPrice.Float64,
				})
		}
		rs.Close()
		l.Println("writing " + code + ".csv")
		of, e := os.Create(out + string(os.PathSeparator) + code + ".csv")
		if e != nil {
			continue
		}
		csvFile = csv.NewWriter(of)
		//write header
		csvFile.Write([]string{"date", "current_price", "open_price", "max", "min", "deal", "deal_price"})
		for _, p := range prices {
			r := []string{
				p.Date,
				fmt.Sprintf("%v", p.CurrentPrice),
				fmt.Sprintf("%v", p.OpenPrice),
				fmt.Sprintf("%v", p.Max),
				fmt.Sprintf("%v", p.Min),
				fmt.Sprintf("%v", p.Deal),
				fmt.Sprintf("%v", p.DealPrice),
			}
			csvFile.Write(r)
		}
		csvFile.Flush()
		of.Close()
	}
}


