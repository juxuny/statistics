package statistics

import "fmt"

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
