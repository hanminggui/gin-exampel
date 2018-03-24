package config

import (
	"os"
	"encoding/json"
	"fmt"
)

type mysql struct {
	Host     string
	User     string
	Pass     string
	Port	 string
	Database string
}

type conf struct {
	Mysql	mysql
	Db 		string
	Port	string
}


var Conf conf

func init()  {
	env := os.Getenv("GIN_MODE")
	if env == "" {
		env = "debug"
	}
	file, _ := os.Open("./config/"+env+".json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&Conf)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(Conf.Mysql.Host)
}