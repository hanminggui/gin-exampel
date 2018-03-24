package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "github.com/hanminggui/gin-exampel/moudles"
	"log"
	"strconv"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

func AddPersion(c *gin.Context) {
	nickname := c.PostForm("nickname")
	password := c.PostForm("password")
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	p := Person{Nickname: nickname, Password: password, FirstName: first_name, LastName: last_name}
	id,err :=  p.AddPerson()
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": id,
	})
}

func GetOnePersion(c *gin.Context)  {
	id,_ := strconv.Atoi(c.Param("id"))
	p := Person{Id: id}
	p.GetDetail()
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"frist_name": p.FirstName,
		"last_name": p.LastName,
	})
}