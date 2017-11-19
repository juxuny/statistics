package statistics

import (
	"testing"
)

func TestCollectStockData (t *testing.T) {
	SetDebug(true)
	c, e := NewCollector("A", DEFAULT_DB_CONIG)
	if e != nil {
		t.Log(e)
		t.Fail()
	}
	e = c.init("sh601006")
	if e != nil {
		t.Log(e)
		t.Fail()
	}
	r, e := c.FetchStockPrice("sh601006")
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

func TestLoadCode(t *testing.T) {
	SetDebug(true)
	r, e := LoadStockCode(DEFAULT_DB_CONIG)
	if e != nil {
		t.Log(e)
		t.Fail()
	}
	t.Log(r)
}

func TestLogger(t *testing.T) {
	SetLogFile("1.log")
	l := GetLogger()
	if l == nil {
		t.Log("l is nil")
		t.Fail()
	}
}