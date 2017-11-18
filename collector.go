package statistics

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"database/sql"
)

const (
	//A股 - 股票列表
	STOCK_CODE_API = "http://www.ctxalgo.com/api/stocks"

	//股票类型
	STOCK_TYPE = "A"
)


type CollectorImpl struct {
	Type string
	db *sql.DB
	config DBConfig
}

func NewCollector(t string, dbConfig DBConfig) (r Collector, e error) {
	c := &CollectorImpl{Type: t}
	c.config = dbConfig
	r = c
	return
}

func (t *CollectorImpl) FetchStockCode()(r []StockCode, e error) {
	url := ""
	switch (t.Type) {
	case STOCK_TYPE:
		url = STOCK_CODE_API
		break
	default:
		return r, fmt.Errorf("unkonw type: %s", t.Type)
		break
	}
	resp, e := http.Get(url)
	if resp.StatusCode != 200 {
		return r, fmt.Errorf("HTTP ERROR: %d", resp.StatusCode)
	}
	log.Print(resp.Header)
	data, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return
	}
	m := make(map[string]string)
	e = json.Unmarshal(data, &m)
	if e != nil {
		return
	}
	for k, v := range m {
		r = append(r, StockCode{Type: t.Type, Code: toUtf8(k), Name: toUtf8(v)})
	}
	return
}

func (t *CollectorImpl) Save(r []StockCode) (e error) {
	db, e := NewConnection(t.config)
	if e != nil {
		log.Print(e)
		return
	}
	tx, e := db.Begin()
	if e != nil {
		log.Print(e)
		return
	}
	for _, v := range r {
		_, e = tx.Exec("INSERT INTO stock_code (code, name, type) SELECT ?, ?, ? FROM DUAL WHERE NOT EXISTS (SELECT id FROM stock_code WHERE code = ?)", v.Code, v.Name, v.Type, v.Code)
		if e != nil {
			log.Print(e)
			tx.Rollback()
			return
		}
	}
	e = tx.Commit()
	if e != nil {
		log.Print(e)
	}
	e = db.Close()
	if e != nil {
		log.Print(e)
	}
	return
}