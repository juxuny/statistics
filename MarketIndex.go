package statistics

import "fmt"

var CSV_HEADER_INDEX = []string{
	"code", "name", "date", "time", "point", "price", "change_rate", "deal", "deal_price",
}

type MarketIndex struct {
	//数据库里自增ID
	Id int
	Code string
	Name string
}


type MarketIndexInfo struct {
	Code string
	Name string
	Date string
	Time string
	//当前点数
	Point float64
	//当前价格
	Price float64
	//成交率
	ChangeRate float64
	//成交量（手）
	Deal float64
	//成交额（万元）
	DealPrice float64
}

func (t MarketIndexInfo) ToStrings() (r []string) {
	r = []string{
		t.Code,
		t.Name,
		t.Date,
		t.Time,
		fmt.Sprintf("%v", t.Price),
		fmt.Sprintf("%v", t.Point),
		fmt.Sprintf("%v", t.ChangeRate),
		fmt.Sprintf("%v", t.Deal),
		fmt.Sprintf("%v", t.DealPrice),
	}
	return
}
