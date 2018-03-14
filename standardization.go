package statistics

import (
	"encoding/json"
	"fmt"
)

type Mean struct {
	//平均价格
	Price float64
	//平均卖出量
	Sell float64
	//平均买入量
	Buy float64
	//平均成交量
	DealVolume float64
	//平均成功总额
	DealPrice float64
}

type Sd struct {
	//价格标准差
	Price float64
	//卖出量标准差
	Sell float64
	//买入量标准差
	Buy float64
	//成交量标准差
	DealVolume float64
	//成交总额标准差
	DealPrice float64
}


func (t Mean) ToStrings() (ret []string) {
	ret = []string{
		fmt.Sprintf("%v", t.Price),
		fmt.Sprintf("%v", t.Sell),
		fmt.Sprintf("%v", t.Buy),
		fmt.Sprintf("%v", t.DealVolume),
		fmt.Sprintf("%v", t.DealPrice),
	}
	return
}

func (t Sd) ToStrings() (ret []string) {
	ret = []string{
		fmt.Sprintf("%v", t.Price),
		fmt.Sprintf("%v", t.Sell),
		fmt.Sprintf("%v", t.Buy),
		fmt.Sprintf("%v", t.DealVolume),
		fmt.Sprintf("%v", t.DealPrice),
	}
	return
}


type StockPriceStandardization StockPrice
type StockPriceStandardizationList []StockPriceStandardization

func (t StockPriceStandardization) String() (ret string) {
	data ,e := json.Marshal(t)
	if e != nil {
		panic(e)
		return ""
	}
	return string(data)
}

func (t StockPriceStandardization) ToStrings() (r []string) {
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

