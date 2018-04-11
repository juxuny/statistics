package statistics

type StockCode struct {
	Type string
	//stock code
	Code string
	//stock name
	Name string
}


func (t StockCode) GetNumberCode() (r string) {
	for i := 0; i < len(t.Code); i++ {
		if t.Code[i] >= '0' && t.Code[i] <= '9' {
			r += string(t.Code[i])
		}
	}
	return r
}

func (t StockCode) GetSohuCode() (r string) {
	return "cn_" + t.GetNumberCode()
}

//默认就是新浪数据的编码，会有sh, sz前缀
func (t StockCode) GetSinaCode() (r string) {
	return t.Code
}
