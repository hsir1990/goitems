package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //这是一个驱动,使用sql包时必须注入（至少）一个数据库驱动
)

var (
	Db  *sql.DB
	err error
)

func init() {
	// ?parseTime=true&loc=Local 设置本地时区
	Db, err = sql.Open("mysql",
		"root:password@tcp(0.0.0.0:3306)/zongheng_online_test_interface?parseTime=true&loc=Local")
	if err != nil {
		panic(err.Error())
	}
}
