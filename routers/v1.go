package routers

import (
	"github.com/gin-gonic/gin"
	. "../apis"
	)


func v1(router *gin.RouterGroup) {
	v1 := router.Group("/v1")
	v1.GET("/persion/:id", GetOnePersion)
	v1.POST("/persion", AddPersion)
}