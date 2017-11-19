package statistics


import (
	l "log"
	"os"
)

var log *l.Logger
var w = &Writer{}

func init() {
	log = l.New(w, "", l.LUTC|l.Ldate|l.Ltime|l.Lshortfile)
}

func SetLogFile(fileName string) (e error) {
	if fileName != "" {
		log = l.New(&FileWriter{fileName: fileName}, "", l.LUTC|l.Ldate|l.Ltime|l.Lshortfile)
		if log == nil {
			l.Panic("log is nil")
		}
	}
	return
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


func GetLogger()  *l.Logger {
	return log
}

type FileWriter struct {
	fileName string
	f *os.File
}

func (t *FileWriter) Write(data []byte) (n int, e error) {
	if _DEBUG {
		if t.f == nil {
			t.f, e = os.OpenFile(t.fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
			if e != nil {
				return
			}
		}
		n, e = t.f.Write(data)
		t.f.Sync()
	}
	n = len(data)
	return
}

func (t *FileWriter) Close() (e error) {
	if t.f != nil {
		t.f.Close()
		t.f = nil
	}
	return
}