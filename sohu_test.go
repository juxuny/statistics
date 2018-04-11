package statistics

import "testing"

func TestCodeTransform(t *testing.T) {
	c := StockCode{Code:"sh600000", Name:"浦发银行", Type:"A"}
	code := c.GetSohuCode()
	if code != "cn_600000" {
		t.Logf("invalid stock code: %s", code)
		t.Fail()
	}

	if c.GetSinaCode() != "sh600000" {
		t.Logf("invalid stock code: %s", c.GetSinaCode())
		t.Fail()
	}
	if c.GetNumberCode() != "600000" {
		t.Logf("invalid stock code: %s", c.GetNumberCode())
		t.Fail()
	}
}


func TestDownloadSohuData(t *testing.T) {
	var tmp = DEFAULT_DB_CONIG
	tmp.DatabaseName = "stock_sohu"
	SetDebug(true)
	collector, e := NewSohuCollector(DEFAULT_DB_CONIG, true)
	if e != nil {
		t.Error(e)
		t.Fail()
		return
	}
	r, e := collector.FetchStockPriceDuration("20180308", "20180411", "600000")
	if e != nil {
		t.Error(e)
		t.Fatal()
	}
	t.Log(r)
	for _, p := range r {
		e = collector.SaveStockPrice(p...)
		if e != nil {
			panic(e)
		}
	}
}
