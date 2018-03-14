package statistics

import (
	"encoding/csv"
	"os"
	"fmt"
	"math/big"
	"math"
)

type StockPriceList []StockPrice

func (t StockPriceList) SaveToCSV(fileName string) (e error) {
	f, e := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0655)
	if e != nil {
		return
	}
	defer f.Close()
	r := csv.NewWriter(f)
	var data [][]string
	for i := 0; i < len(t); i++ {
		data = append(data, t[i].ToStrings()[2:])
	}
	r.Write(CSV_HEADER[2:])
	r.WriteAll(data)
	return
}

//数据标准化
func (t StockPriceList) Standardization() (ret StockPriceStandardizationList, mean Mean, sd Sd, e error) {
	num := len(t)
	if num == 0 {
		e = fmt.Errorf("no data")
		return
	}
	const SCALE = 1e9 // 放大原始数据，用于高精度除法

	//calculate the mean
	priceSum := big.NewInt(0)
	sellVolumeSum := big.NewInt(0)
	buyVolumeSum := big.NewInt(0)
	dealVolumeSum := big.NewInt(0)
	dealPriceSum := big.NewInt(0)
	for _, d := range t {
		cp := big.NewInt(int64(d.CurrentPrice * SCALE))
		sv := big.NewInt(int64(Sum(d.Sell[:])) * SCALE)
		bv := big.NewInt(int64(Sum(d.Buy[:])) * SCALE)
		dv := big.NewInt(int64(d.Deal * SCALE))
		dp := big.NewInt(int64(d.DealPrice * SCALE))


		priceSum = priceSum.Add(priceSum, cp)
		sellVolumeSum = sellVolumeSum.Add(sellVolumeSum, sv)
		buyVolumeSum = buyVolumeSum.Add(buyVolumeSum, bv)
		dealVolumeSum = dealVolumeSum.Add(dealVolumeSum, dv)
		dealPriceSum = dealPriceSum.Add(dealPriceSum, dp)

	}
	mean.Price = float64(priceSum.Div(priceSum, big.NewInt(int64(num))).Int64()) / SCALE
	mean.Sell = float64(sellVolumeSum.Div(sellVolumeSum, big.NewInt(int64(num) * 5)).Int64()) / SCALE
	mean.Buy = float64(buyVolumeSum.Div(buyVolumeSum, big.NewInt(int64(num) * 5)).Int64()) / SCALE
	mean.DealVolume = float64(dealVolumeSum.Div(dealVolumeSum, big.NewInt(int64(num))).Int64()) / SCALE
	mean.DealPrice = float64(dealPriceSum.Div(dealPriceSum, big.NewInt(int64(num))).Int64()) / SCALE

	//calculate the standard deviation
	for _, v := range t {
		sd.Price += math.Pow(v.CurrentPrice - mean.Price, 2)
		sd.Sell += math.Pow(v.Sell[0] - mean.Sell, 2)
		sd.Sell += math.Pow(v.Sell[1] - mean.Sell, 2)
		sd.Sell += math.Pow(v.Sell[2] - mean.Sell, 2)
		sd.Sell += math.Pow(v.Sell[3] - mean.Sell, 2)
		sd.Sell += math.Pow(v.Sell[4] - mean.Sell, 2)

		sd.Buy += math.Pow(v.Buy[0] - mean.Buy, 2)
		sd.Buy += math.Pow(v.Buy[1] - mean.Buy, 2)
		sd.Buy += math.Pow(v.Buy[2] - mean.Buy, 2)
		sd.Buy += math.Pow(v.Buy[3] - mean.Buy, 2)
		sd.Buy += math.Pow(v.Buy[4] - mean.Buy, 2)

		sd.DealVolume += math.Pow(v.Deal - mean.DealVolume, 2)
		sd.DealPrice += math.Pow(v.DealPrice - mean.DealPrice, 2)

	}
	sd.Price = math.Sqrt(sd.Price/float64(num-1))
	sd.Sell = math.Sqrt(sd.Sell/float64(num*5-1))
	sd.Buy = math.Sqrt(sd.Buy/float64(num*5-1))
	sd.DealVolume = math.Sqrt(sd.DealVolume/float64(num-1))
	sd.DealPrice = math.Sqrt(sd.DealPrice/float64(num-1))

	//standardization
	for _, v := range t {
		var item StockPriceStandardization
		item.StockCode = v.StockCode
		item.Name = v.Name
		item.Date = v.Date
		item.Time = v.Time
		item.CurrentPrice = (v.CurrentPrice - mean.Price) / sd.Price
		for i := 0; i < len(v.Sell); i++ {
			item.Sell[i] = (v.Sell[i] - mean.Sell) / sd.Sell
			item.Buy[i] = (v.Buy[i] - mean.Buy) / sd.Buy
		}
		item.Deal = (v.Deal - mean.DealVolume) / sd.DealVolume
		item.DealPrice = (v.DealPrice - mean.DealPrice) / sd.DealPrice

		//other price
		item.OpenPrice = (v.OpenPrice - mean.Price) / sd.Price
		item.YesterdayPrice = (v.YesterdayPrice - mean.Price) / sd.Price
		item.Max = (v.Max - mean.Price) / sd.Price
		item.Min = (v.Min - mean.Price) / sd.Price
		for i := 0; i < len(v.Sell); i++ {
			item.SellPrice[i] = (v.SellPrice[i] - mean.Price) / sd.Price
			item.BuyPrice[i] = (v.BuyPrice[i] - mean.Price) / sd.Price
		}
		ret = append(ret, item)
	}
	return
}