package statistics

type OneHot []bool

func NewOneHot(k, length int) (ret OneHot) {
	ret = make([]bool, length)
	ret[k] = true
	return
}

func (t OneHot) ToStrings() (ret []string) {
	for _, v := range t {
		if v {
			ret = append(ret, "1")
		} else {
			ret = append(ret, "0")
		}
	}
	return
}