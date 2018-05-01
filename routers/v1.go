package routers

import (
	"github.com/gin-gonic/gin"
	. "github.com/hanminggui/gin-exampel/apis/v1"
)

func v1(router *gin.RouterGroup) {
	v1 := router.Group("/v1")
	user(v1)
	apply(v1)
	share(v1)
}

func user(router *gin.RouterGroup)  {
	router.GET("/user/:id", GetOneUser)
	router.POST("/user", AddUser)
	router.DELETE("/user/:id", DeleteUser)
}

func apply(router *gin.RouterGroup)  {
	router.GET("/apply/:id", GetOneApply)
	router.POST("/apply", AddApply)
}

func share(router *gin.RouterGroup)  {
	router.GET("/share/:id", GetOneShare)
	router.POST("/share", AddShare)
	router.DELETE("/share/:id", DeleteShare)
}

/**
创建用户
修改用户
获取用户

用户发布的分享列表
用户的报名列表
用户的粉丝列表
用户的关注列表


分享列表
分享的详情
发布分享
修改分享
删除分享

报名
修改我的分享里的报名

关注
取消关注

 */