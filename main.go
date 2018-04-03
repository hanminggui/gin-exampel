package main // import "github.com/hanminggui/gin-exampel"

import (
	. "github.com/hanminggui/gin-exampel/config"
	db "github.com/hanminggui/gin-exampel/database"
	r "github.com/hanminggui/gin-exampel/routers"
	"log"
)

func main() {
	log.SetFlags(19)
	defer db.SqlDB.Close()
	router := r.InitRouter()
	router.Run(":" + Conf.Port)
}
