package main // import "github.com/hanminggui/gin-exampel"

import (
	. "github.com/hanminggui/gin-exampel/config"
	db "github.com/hanminggui/gin-exampel/database"
	r "github.com/hanminggui/gin-exampel/routers"
	"log"
	//"time"
	//"net/http"
	"os"
)

func main() {
	log.SetFlags(19)
	log.SetOutput(os.Stderr)
	defer db.SqlDB.Close()
	router := r.InitRouter()
	router.Run(":" + Conf.Port)
	//s := &http.Server{
	//	Addr:           ":8080",
	//	Handler:        router,
	//	ReadTimeout:    2 * time.Second,
	//	WriteTimeout:   2 * time.Second,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//s.ListenAndServe()
}
