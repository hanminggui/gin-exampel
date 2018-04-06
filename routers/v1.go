package routers

import (
	"github.com/gin-gonic/gin"
	. "github.com/hanminggui/gin-exampel/apis/v1"
)

func v1(router *gin.RouterGroup) {
	v1 := router.Group("/v1")
	user(v1)
	apply(v1)
}

func user(router *gin.RouterGroup)  {
	router.GET("/user/:id", GetOneUser)
	router.POST("/user", AddUser)
	router.DELETE("/user/:id", DeleteUser)
}

func apply(router *gin.RouterGroup)  {
	router.GET("/apply/:id", GetOneApply)
	router.POST("/apply", AddApply)
	router.DELETE("/apply/:id", DeleteApply)
}