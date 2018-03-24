package main // import "github.com/hanminggui/gin-exampel"

import (
	db "github.com/hanminggui/gin-exampel/database"
	r "github.com/hanminggui/gin-exampel/routers"
	. "github.com/hanminggui/gin-exampel/config"
)
func main() {

	defer db.SqlDB.Close()
	router := r.InitRouter()
	router.Run(":" + Conf.Port)
}
