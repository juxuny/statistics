package statistics

import (
	"golang.org/x/text/encoding/charmap"
	//"bytes"
)

func toUtf8(src string) (s string) {
	return src
	d := charmap.ISO8859_1.NewDecoder()
	s, e := d.String(src)
	if e != nil {
		log.Print(e)
	}
	return
}
