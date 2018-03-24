package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	. "github.com/hanminggui/gin-exampel/config"
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