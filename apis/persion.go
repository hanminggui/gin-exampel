package apis

import (
	"github.com/gin-gonic/gin"
	. "github.com/hanminggui/gin-exampel/moudles"
	"log"
	"net/http"
	"strconv"
	)

func IndexApi(c *gin.Context) {
	u := User{}
	u.Id = 3
	u.NickName = "第三个用户"
	u.Update()
	c.String(http.StatusOK, "u name is " + u.NickName)
}

func AddPersion(c *gin.Context) {
	nickname := c.PostForm("nickname")
	password := c.PostForm("password")
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	p := Person{Nickname: nickname, Password: password, FirstName: first_name, LastName: last_name}
	id, err := p.AddPerson()
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    id,
	})
}

func GetOnePersion(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	p := Person{Id: id}
	p.GetDetail()
	c.JSON(http.StatusOK, gin.H{
		"message":    "success",
		"frist_name": p.FirstName,
		"last_name":  p.LastName,
	})
}
