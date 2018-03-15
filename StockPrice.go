package statistics

import (
	"fmt"
	"strconv"
)

var CSV_HEADER = []string{
	"stock_code", "name", "date", "time", "current_price", "open_price","yesterday_price", "max", "min",
	"buy_price_1", "buy_price_2", "buy_price_3", "buy_price_4", "buy_price_5",
	"buy_1", "buy_2", "buy_3", "buy_4", "buy_5",
	"sell_price_1", "sell_price_2", "sell_price_3", "sell_price_4", "sell_price_5",
	"sell_1", "sell_2", "sell_3", "sell_4", "sell_5",
	"deal", "deal_price",
}

type StockPrice struct {
	StockCode string
	Name string
	Date string
	Time string
	CurrentPrice float64
	OpenPrice float64
	YesterdayPrice float64
	Max float64
	Min float64
	BuyPrice [5]float64
	Buy [5]float64
	SellPrice [5]float64
	Sell [5]float64
	Deal float64
	DealPrice float64
}


func (t StockPrice) ToStrings() (r []string) {
	r = []string{
		t.StockCode,
		t.Name,
		t.Date,
		t.Time,
		fmt.Sprintf("%v", t.CurrentPrice),
		fmt.Sprintf("%v", t.OpenPrice),
		fmt.Sprintf("%v", t.YesterdayPrice),
		fmt.Sprintf("%v", t.Max),
		fmt.Sprintf("%v", t.Min),
		fmt.Sprintf("%v", t.BuyPrice[0]),
		fmt.Sprintf("%v", t.BuyPrice[1]),
		fmt.Sprintf("%v", t.BuyPrice[2]),
		fmt.Sprintf("%v", t.BuyPrice[3]),
		fmt.Sprintf("%v", t.BuyPrice[4]),
		fmt.Sprintf("%v", t.Buy[0]),
		fmt.Sprintf("%v", t.Buy[1]),
		fmt.Sprintf("%v", t.Buy[2]),
		fmt.Sprintf("%v", t.Buy[3]),
		fmt.Sprintf("%v", t.Buy[4]),
		fmt.Sprintf("%v", t.SellPrice[0]),
		fmt.Sprintf("%v", t.SellPrice[1]),
		fmt.Sprintf("%v", t.SellPrice[2]),
		fmt.Sprintf("%v", t.SellPrice[3]),
		fmt.Sprintf("%v", t.SellPrice[4]),
		fmt.Sprintf("%v", t.Sell[0]),
		fmt.Sprintf("%v", t.Sell[1]),
		fmt.Sprintf("%v", t.Sell[2]),
		fmt.Sprintf("%v", t.Sell[3]),
		fmt.Sprintf("%v", t.Sell[4]),
		fmt.Sprintf("%v", t.Deal),
		fmt.Sprintf("%v", t.DealPrice),
	}
	return
}

//与ToStrings方法相反
func ParseStockPriceFromStrings(row []string) (ret StockPrice, e error) {
	if len(row) != 31 {
		e = fmt.Errorf("invalid row len: %v", len(row))
		return
	}
	if row[0] == "" {
		e = fmt.Errorf("invalid data: first element is the stock code, it can't be empty")
		return
	}
	ret.StockCode = row[0]
	ret.Name = row[1]
	ret.Date = row[2]
	ret.Time = row[3]
	ret.CurrentPrice, e = strconv.ParseFloat(row[4], 64)
	if e != nil { return ret, e }
	ret.OpenPrice, e = strconv.ParseFloat(row[5], 64)
	if e != nil { return ret, e }
	ret.YesterdayPrice, e = strconv.ParseFloat(row[6], 64)
	if e != nil { return ret, e }
	ret.Max, e = strconv.ParseFloat(row[7], 64)
	if e != nil { return ret, e }
	ret.Min, e = strconv.ParseFloat(row[8], 64)
	if e != nil { return ret, e }

	ret.BuyPrice[0], e = strconv.ParseFloat(row[9], 64)
	if e != nil { return ret, e }
	ret.BuyPrice[1], e = strconv.ParseFloat(row[10], 64)
	if e != nil { return ret, e }
	ret.BuyPrice[2], e = strconv.ParseFloat(row[11], 64)
	if e != nil { return ret, e }
	ret.BuyPrice[3], e = strconv.ParseFloat(row[12], 64)
	if e != nil { return ret, e }
	ret.BuyPrice[4], e = strconv.ParseFloat(row[13], 64)
	if e != nil { return ret, e }

	ret.Buy[0], e = strconv.ParseFloat(row[14], 64)
	if e != nil { return ret, e }
	ret.Buy[1], e = strconv.ParseFloat(row[15], 64)
	if e != nil { return ret, e }
	ret.Buy[2], e = strconv.ParseFloat(row[16], 64)
	if e != nil { return ret, e }
	ret.Buy[3], e = strconv.ParseFloat(row[17], 64)
	if e != nil { return ret, e }
	ret.Buy[4], e = strconv.ParseFloat(row[18], 64)
	if e != nil { return ret, e }

	ret.SellPrice[0], e = strconv.ParseFloat(row[19], 64)
	if e != nil { return ret, e }
	ret.SellPrice[1], e = strconv.ParseFloat(row[20], 64)
	if e != nil { return ret, e }
	ret.SellPrice[2], e = strconv.ParseFloat(row[21], 64)
	if e != nil { return ret, e }
	ret.SellPrice[3], e = strconv.ParseFloat(row[22], 64)
	if e != nil { return ret, e }
	ret.SellPrice[4], e = strconv.ParseFloat(row[23], 64)
	if e != nil { return ret, e }

	ret.Sell[0], e = strconv.ParseFloat(row[24], 64)
	if e != nil { return ret, e }
	ret.Sell[1], e = strconv.ParseFloat(row[25], 64)
	if e != nil { return ret, e }
	ret.Sell[2], e = strconv.ParseFloat(row[26], 64)
	if e != nil { return ret, e }
	ret.Sell[3], e = strconv.ParseFloat(row[27], 64)
	if e != nil { return ret, e }
	ret.Sell[4], e = strconv.ParseFloat(row[28], 64)
	if e != nil { return ret, e }

	ret.Deal, e = strconv.ParseFloat(row[29], 64)
	if e != nil { return ret, e }
	ret.DealPrice, e = strconv.ParseFloat(row[30], 64)
	if e != nil { return ret, e }
	return
}