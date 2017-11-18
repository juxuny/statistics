package statistics

type StockCode struct {
	Type string
	//stock code
	Code string
	//stock name
	Name string
}

type Collector interface {
	FetchStockCode() (r []StockCode, e error)
	Save(r []StockCode) (e error)
}