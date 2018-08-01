package statistics

import (
	"bytes"
	"io/ioutil"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/simplifiedchinese"
	"regexp"
	"os"
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

func ConvertSinaCodeToSohuCode(code string) (s string) {
	for _, i := range code {
		if i >= '0' && i <= '9' {
			s += string(i)
		}
	}
	return "cn_" + s
}


//只返回数字
func NumberFilter(s string) (r string) {
	for _, i := range s {
		if i >= '0' && i <= '9' {
			r += string(i)
		}
	}
	return r
}


//判断路径是否存在
func PathExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}