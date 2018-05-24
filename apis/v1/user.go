package v1

import (
	"github.com/gin-gonic/gin"
	. "github.com/hanminggui/gin-exampel/moudles"
	"net/http"
)

func GetOneUser(c *gin.Context) {
	id := getInt64(c, "id")
	user := User{Id: id}
	user.GetDetail()
	if user.Id == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}


func AddUser(c *gin.Context) {
	var user User
	err := c.BindJSON(&user)
	Check(err)
	id, err := user.Add()
	Check(err)
	user.Id = id
	user.GetDetail()
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context)  {
	id := getInt64(c, "id")
	user := User{Id: id}
	user.Delete()
	c.JSON(http.StatusOK, user)
}

func GetUserShares(c *gin.Context)  {
	id := getInt64(c, "id")
	user := User{Id: id}
	user.GetShares(c.GetInt("offset"), c.GetInt("limit"))
	c.JSON(http.StatusOK, user.Shares)
}

func GetUserApplys(c *gin.Context)  {
	id := getInt64(c, "id")
	user := User{Id: id}
	user.GetApplys(c.GetInt("offset"), c.GetInt("limit"))
	c.JSON(http.StatusOK, user.Applys)
}

func GetUserFanss(c *gin.Context)  {
	id := getInt64(c, "id")
	user := User{Id: id}
	user.GetFanss(c.GetInt("offset"), c.GetInt("limit"))
	c.JSON(http.StatusOK, user.Fanss)
}

func GetUserFollows(c *gin.Context)  {
	id := getInt64(c, "id")
	user := User{Id: id}
	user.GetFollows(c.GetInt("offset"), c.GetInt("limit"))
	c.JSON(http.StatusOK, user.Follows)
}

func Follow(c *gin.Context)  {
	userId := getInt64(c, "userId")
	toUserId := getInt64(c, "toUserId")
	attention := Attention{UserId: userId, ToUserId: toUserId}
	attention.Follow()
	c.JSON(http.StatusOK, attention)
}

func UnFollow(c *gin.Context)  {
	userId := getInt64(c, "userId")
	toUserId := getInt64(c, "toUserId")
	attention := Attention{UserId: userId, ToUserId: toUserId}
	attention.UnFollow()
	c.JSON(http.StatusOK, attention)
}