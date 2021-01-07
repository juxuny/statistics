package statistics

type DBConfig struct {
	Host string
	Port int
	User string
	Password string
	DatabaseName string
	InitTable bool
}

var _DEBUG = false
//默认配置
var DEFAULT_DB_CONIG = NewDefaultDBConfig()

func SetDebug(f bool) {
	_DEBUG = f
}

//返回默认的数据库配置
func NewDefaultDBConfig() (r DBConfig) {
	r.User = "root"
	r.Password = "123456"
	r.Host = "127.0.0.1"
	r.Port = 3307
	r.DatabaseName = "stock"
	r.InitTable = true
	return
}