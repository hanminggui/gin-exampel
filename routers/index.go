package routers

import (
	"github.com/gin-gonic/gin"
	. "github.com/hanminggui/gin-exampel/apis"
	"fmt"
)

func InitRouter() *gin.Engine {
	fmt.Println("gin run mode ------------------- " + gin.Mode())
	router := gin.Default()
	router.GET("/", IndexApi)
	api := router.Group("/api")
	v1(api)
	return router
}

