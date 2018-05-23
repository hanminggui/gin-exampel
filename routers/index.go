package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "github.com/hanminggui/gin-exampel/apis"
)

func InitRouter() *gin.Engine {
	fmt.Println("gin run mode ------------------- " + gin.Mode())
	router := gin.Default()
	router.Use(middleware(), timeOut())
	router.GET("/", IndexApi)
	api := router.Group("/api")
	v1(api)
	return router
}
