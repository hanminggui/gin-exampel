package main

import (
	db "./database"
	r "./routers"
)
func main() {
	defer db.SqlDB.Close()
	router := r.InitRouter()
	router.Run(":8000")
}
