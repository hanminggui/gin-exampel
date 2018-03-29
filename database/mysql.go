package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/hanminggui/gin-exampel/config"
	"log"
	"fmt"
)

var SqlDB *sql.DB

func init() {
	var err error
	conf := &Conf.Mysql

	SqlDB, err = sql.Open(Conf.Db, conf.User+":"+conf.Pass+"@tcp("+conf.Host+":"+conf.Port+")/"+conf.Database+"?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("mysql conn ok")
}

func Insert(tableName string, moudle struct{}) (id int64, err error) {

	return
}

func Update(tableName string, moudle struct{}) (err error) {

	return
}
func QueryValue(sql string, args ...interface{}) (value interface{}, err error) {
	row := SqlDB.QueryRow(sql, args)
	err = row.Scan(value)
	return
}
//func QueryValue2(value *interface{}, sql string, args ...interface{}) (err error) {
//	row := SqlDB.QueryRow(sql, args)
//	err = row.Scan(value)
//	return
//}
func QueryInt64(value *int64, sql string, args ...interface{}) (error) {
	row := SqlDB.QueryRow(sql, args)
	err := row.Scan(value)
	return err
}
func QueryInt(value *int, sql string, args ...interface{}) (err error) {
	interfs,err := QueryValue(sql, args)
	val,ok := interfs.(int)
	if !ok {
		fmt.Println("类型转换错误")
	}
	value = &val
	return
}
func QueryString(sql string, args ...interface{}) (value string, err error) {

	return
}
func FindOne(moudle *struct{}, sql string, args ...interface{}) (err error) {

	return
}
func FindList(moudle []*struct{}, sql string, args ...interface{}) (err error) {

	return
}

