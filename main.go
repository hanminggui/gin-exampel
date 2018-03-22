package main

import (
	db "./database"
	r "./routers"
	. "./config"
)
func main() {

	defer db.SqlDB.Close()
	router := r.InitRouter()
	router.Run(":" + Conf.Port)
}
