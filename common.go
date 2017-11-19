package statistics

import (
	"bytes"
	"io/ioutil"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/simplifiedchinese"
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