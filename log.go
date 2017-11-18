package statistics


import (
	l "log"
	"os"
)

var log *l.Logger
var w = &Writer{}

func init() {
	log = l.New(w, "", l.LUTC|l.Ldate|l.Ltime)
}


type Writer struct {

}

func (t *Writer) Write(data []byte) (n int, e error) {
	if _DEBUG {
		return os.Stdout.Write(data)
	}
	n = len(data)
	return
}

func (t *Writer) Close() (e error) {
	return
}

