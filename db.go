package statistics


import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

func NewConnection(c DBConfig) (db *sql.DB, e error) {
	db, e = sql.Open("mysql",fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", c.User, c.Password, c.Host, c.Port, c.DatabaseName))
	return
}
