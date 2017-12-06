package statistics

import (
	"encoding/csv"
	"os"
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