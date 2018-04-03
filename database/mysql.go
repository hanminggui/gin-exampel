package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/hanminggui/gin-exampel/config"
	"log"
	"reflect"
	"strconv"
	"time"
)

var SqlDB *sql.DB
var ModMap = make(map[string]interface{})

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
type Map map[string]string
type Maps []Map
func Insert(tableName string, moudle interface{}) (id int64, err error) {
	columns, values := getKV(moudle)
	sqlCenter := ""
	sqlCenter2 := ""
	for i:=0; i<len(columns); i++ {
		sqlCenter += columns[i]
		sqlCenter2 += "?"
		if i<len(columns)-1 {
			sqlCenter += ", "
			sqlCenter2 += ", "
		}
	}
	r,e := SqlDB.Exec("INSERT "+tableName+"("+sqlCenter+") VALUE (" + sqlCenter2 + ")", values...)
	err = e
	log.Println(r,e)
	return
}

func Update(tableName string, moudle interface{}) (err error) {
	columns, values := getKV(moudle)
	moudleValues := reflect.ValueOf(moudle).Elem()
	if moudleValues.FieldByName("Id").String() == "0" {
		Insert(tableName, moudle)
		return
	}
	sqlCenter := ""
	for i:=0; i<len(columns); i++ {
		sqlCenter += columns[i] + "=?"
		if i<len(columns)-1 {
			sqlCenter += ", "
		}
	}
	values = append(values, moudleValues.FieldByName("Id"))
	r,e := SqlDB.Exec("UPDATE "+tableName+" SET " + sqlCenter + "WHERE id=?", values...)
	err = e
	log.Println(r,e)
	return
}

func getKV(moudle interface{}) ([]string, []interface{}) {
	columns := make([]string, 0)
	values := make([]interface{}, 0)
	moudleType := reflect.TypeOf(moudle).Elem()
	moudleValues := reflect.ValueOf(moudle).Elem()
	for i:=0; i< moudleValues.NumField(); i++ {
		field := moudleType.Field(i)
		kind := field.Type.Kind()
		if kind == reflect.String || kind == reflect.Int || kind == reflect.Int64 && (moudleValues.Field(
			i).String() != "0" && moudleValues.Field(i).String() != "") {
			columns = append(columns, field.Tag.Get("json"))
			values = append(values, moudleValues.Field(i))
		}
	}
	return columns, values
}

/**
 * 查询单个值并赋值
 * @value 可以是任意类型 被赋值目标
 * @query sql语句
 * @args 替换sql中的参数
 */
func QueryValue(value interface{}, query string, args ...interface{}) (err error) {
	err = SqlDB.QueryRow(query, args...).Scan(value)
	return
}

/**
 *
 */
func QueryOne(moudle interface{}, query string, args ...interface{}) (err error) {
	mp,err1 := QueryMap(query, args...)
	if err1 != nil {
		err = err1
		return
	}
	err = mp.Load(moudle)
	return
}

/**
 * @moudles 对象切片的指针，需要先填充好对应数量的空对象再放进来
 */
//func (maps Maps) Load(moudles interface{}) (err error) {
//	list := reflect.ValueOf(moudles).Elem()
//	for i:=0; i<len(maps); i++{
//		err = maps[i].Load(list.Index(i))
//		if err != nil {
//			return
//		}
//	}
//	return
//}

func QueryMap(query string, args ...interface{}) (result Map, err error) {
	maps,err1 := QueryMaps(query, args...)
	err = err1
	result = maps[0]
	return
}

/**
 * 查询结果组装成map
 */
func QueryMaps(query string, args ...interface{}) (results Maps, err error) {
	log.Println("querying: ", query, args)
	beginTime := time.Now()
	rows, err := SqlDB.Query(query, args...)
	endTime := time.Now()
	if err != nil{
		log.Fatalln(err)
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil{
		log.Fatalln(err)
	}
	vals := make([][]byte, len(cols))
	scans := make([]interface{}, len(cols))

	for i := range vals{
		scans[i] = &vals[i]
	}

	for rows.Next(){
		err = rows.Scan(scans...)
		if err != nil{
			log.Fatalln(err)
		}

		row := make(map[string]string)
		for k, v := range vals{
			key := cols[k]
			row[key] = string(v)
		}
		results = append(results, row)
	}
	if len(results) > 0 {
		log.Printf("used %d ms total rows: %d rows0: %s\n", endTime.UnixNano()/1000000 - beginTime.UnixNano()/1000000,
			len(results), results[0])
	} else {
		log.Printf("used %d ms total rows: 0\n", endTime.UnixNano()/1000000 - beginTime.UnixNano()/1000000)

	}
	return
}


func (mp Map) Load(moudle interface{}) (err error) {
	moudleType := reflect.TypeOf(moudle).Elem()
	values := reflect.ValueOf(moudle).Elem()
	for i:=0; i<values.NumField(); i++ {
		field := moudleType.Field(i)
		var val int64 = 0
		switch field.Type.Kind() {
		case reflect.String:
			values.Field(i).SetString(mp[field.Tag.Get("json")])
		case reflect.Int:
			if len(mp[field.Tag.Get("json")]) > 0 {
				val, err = strconv.ParseInt(mp[field.Tag.Get("json")], 10, 64)
				if err != nil {
					return
				}
			}
			values.Field(i).SetInt(val)
		case reflect.Int64:
			if len(mp[field.Tag.Get("json")]) > 0 {
				val, err = strconv.ParseInt(mp[field.Tag.Get("json")], 10, 64)
				if err != nil {
					return
				}
			}
			values.Field(i).SetInt(val)
		}
	}
	return
}