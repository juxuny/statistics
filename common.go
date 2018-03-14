package statistics

import (
	"bytes"
	"io/ioutil"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/simplifiedchinese"
	"regexp"
)

func GBK_UTF8(src string) (string, error) {
	s := []byte(src)
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return "", e
	}
	return string(d), nil
}

func CheckDate(str string) (r bool) {
	r, e := regexp.MatchString("([\\d]{4})-([\\d]{2})-[\\d]{2}", str)
	if e != nil {
		panic(e)
	}
	return
}

//对数组求和
func Sum(data []float64) (ret float64) {
	for _, v := range data {
		ret += v
	}
	return
}

func MergeStrings(str ... []string) (ret []string) {
	for _, ss := range str {
		for _, s := range ss {
			ret = append(ret, s)
		}
	}
	return
}