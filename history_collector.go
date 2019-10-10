package statistics

type HistoryCollector struct {
	FCollector
	Type string
	Prefix string
}

// 初始化本地数据下载器
func NewHistoryCollector() (r *HistoryCollector) {
	c := &HistoryCollector{Type: STOCK_TYPE_A, Prefix: "sina_"}
	r = c
	return
}

// 获取一天历史数据
func (t *HistoryCollector) OneDayHistory(code string, date string) (data []StockPrice, err error) {
	return
}