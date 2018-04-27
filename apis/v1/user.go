package v1

import (
	"github.com/gin-gonic/gin"
	. "github.com/hanminggui/gin-exampel/moudles"
	"net/http"
	"strconv"
)

func GetOneUser(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	user := User{Id: id}
	user.GetDetail()
	c.JSON(http.StatusOK, user)
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
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	user := User{Id: id}
	user.Delete()
	c.JSON(http.StatusOK, user)
}

func GetUserShares(c *gin.Context)  {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	user := User{Id: id}
	user.GetShares()
	c.JSON(http.StatusOK, gin.H{"shares": user.Shares})
}

func GetUserApplys(c *gin.Context)  {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	user := User{Id: id}
	user.GetApplys()
	c.JSON(http.StatusOK, gin.H{"applys": user.Applys})
}

func GetUserFanss(c *gin.Context)  {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	user := User{Id: id}
	user.GetFanss()
	c.JSON(http.StatusOK, gin.H{"fanss": user.Fanss})
}

func GetUserFollows(c *gin.Context)  {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	user := User{Id: id}
	user.GetFollows()
	c.JSON(http.StatusOK, gin.H{"follows": user.Follows})
}