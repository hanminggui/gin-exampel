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

func Insert(tableName string, moudle struct{}) (id int64, err error) {

	return
}

func Update(tableName string, moudle struct{}) (err error) {

	return
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
	maps,err1 := QueryMap(query, args...)
	if err1 != nil {
		err = err1
		return
	}
	err = LoadMoudle(moudle, maps[0])
	return
}

func FindList2(moudles interface{}, query string, args ...interface{}) (err error) {
	maps,err1 := QueryMap(query, args...)
	if err1 != nil {
		err = err1
		return
	}
	var models []interface{}
	for _,m := range maps{
		name := reflect.TypeOf(moudles).Name()
		v := reflect.New(reflect.TypeOf(ModMap[name]))
		err = LoadMoudle(&v, m)
		if err != nil {
			return
		}
		models = append(models, v.Interface())
	}
	log.Println("models:", models)
	moudles = append(models)
	return
}

func FindList(moudles interface{}, moudle interface{}, query string, args ...interface{}) (err error) {
	maps,err1 := QueryMap(query, args...)
	if err1 != nil {
		err = err1
		return
	}
	a0 := reflect.ValueOf(moudles).Elem()
	e0 := make([]reflect.Value, 0)
	for _,m := range maps{
		//mi := moudle.(reflect.Type)
		mi := reflect.New(reflect.TypeOf(&moudle).Elem()).Elem()
		//mi := ModMap[a0.Type().Name()]
		err = LoadMoudle(&mi, m)
		if err != nil {
			return
		}
		e0 = append(e0, reflect.ValueOf(mi))
	}
	val_arr1 := reflect.Append(a0, e0...)
	a0.Set(val_arr1)
	//a1 := reflect.TypeOf(moudles).Elem()
	//e0 := make([]reflect.Value, 0)
	//fmt.Print(a0.Type())
	//fmt.Print("n2 is ", a1.Name())
	//fmt.Print(a0.Elem().Type())
	//fmt.Print(a0.Index(0).Type())
	//e0 = append(e0, reflect.ValueOf(100))
	return
}

/**
 * 查询结果组装成map
 */
func QueryMap(query string, args ...interface{}) (results []map[string]string, err error) {
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

func LoadMoudle(moudle interface{}, mp map[string]string) (err error) {
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
		//fmt.Print("	field type is :", field.Type)
		//fmt.Print("	field type name is :", field.Type.Name())
		//fmt.Print("	field name is :", field.Name)
		//fmt.Print("	field json tag is :", field.Tag.Get("json"))
		//fmt.Print("	field value is :", values.Field(i))
		//fmt.Println()
	}
	return
}