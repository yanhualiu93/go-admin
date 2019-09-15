package custom

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
)

func NewDb(username, password, host, port, database string) *xorm.Engine {
	url := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8", username, password, host, port, database)
	engine, err := xorm.NewMySQL(xorm.MYSQL_DRIVER, url)
	if err != nil {
		panic(err)
	}
	err = engine.Ping()
	if err != nil {
		panic(err)
	}
	return engine
}
