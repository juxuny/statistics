package statistics

type DBConfig struct {
	Host string
	Port int
	User string
	Password string
	DatabaseName string
}

var _DEBUG = false


func SetDebug(f bool) {
	_DEBUG = f
}