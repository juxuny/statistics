package statistics

import (
	"testing"
)

func TestLoadCode(t *testing.T) {
	SetDebug(true)
	r, e := LoadStockCode(DEFAULT_DB_CONIG)
	if e != nil {
		t.Log(e)
		t.Fail()
	}
	t.Log(r)
}

func TestCollectStockData (t *testing.T) {
	SetDebug(true)
	c := NewCollector(DEFAULT_DB_CONIG)

	//stockCodes, e := c.FetchStockCode()
	//if e != nil {
	//	t.Log(e)
	//	t.Fail()
	//}
	//c.SaveStockCode(stockCodes)

	r, e := c.FetchStockPrices("sh600832", "sh600439")
	if e != nil {
		t.Log(e)
		t.Fail()
	}
	t.Log(r)
	e = c.SaveStockPrice(r["sh600832"])
	if e != nil {
		t.Log(e)
		t.Fail()
	}
}

func TestLogger(t *testing.T) {
	SetLogFile("1.log")
	l := GetLogger()
	if l == nil {
		t.Log("l is nil")
		t.Fail()
	}
}

func TestMarketIndex(t *testing.T) {
	marketIndexes, e := LoadMarketIndexes(DEFAULT_DB_CONIG)
	if e != nil {
		t.Log(e)
		t.Fail()
		return
	}
	t.Log(marketIndexes)
	codes := make([]string, 0)
	for _, v := range marketIndexes {
		codes = append(codes, v.Code)
	}
	collector := NewCollector(DEFAULT_DB_CONIG)
	r, e := collector.FetchMarketIndexes(codes...)
	if e != nil {
		t.Log(e)
		t.Fail()
		return
	}
	t.Log(r)
	var tmp []MarketIndexInfo
	for _, v := range r {
		tmp = append(tmp, v)
	}
	e = collector.SaveMarketIndexesData(tmp...)
	if e != nil {
		t.Log(e)
		t.Fail()
	}
}