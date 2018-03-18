package routers

import (
	"github.com/gin-gonic/gin"
	. "../apis"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	v1(router)
	router.GET("/", IndexApi)
	return router
}

