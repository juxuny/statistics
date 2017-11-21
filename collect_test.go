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
	c, e := NewCollector("A", DEFAULT_DB_CONIG)
	if e != nil {
		t.Log(e)
		t.Fail()
	}
	//stockCodes, e := c.FetchStockCode()
	//if e != nil {
	//	t.Log(e)
	//	t.Fail()
	//}
	//c.SaveStockCode(stockCodes)

	r, e := c.FetchStockPrices("sh601006", "sh600439")
	if e != nil {
		t.Log(e)
		t.Fail()
	}
	t.Log(r)
	e = c.SaveStockPrice(r["sh601006"])
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