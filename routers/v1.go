package routers

import (
	"github.com/gin-gonic/gin"
	. "../apis"
	)


func v1(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	v1.GET("/persion/:id", GetOnePersion)
	v1.POST("/persion", AddPersion)
}