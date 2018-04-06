package v1

import (
	"github.com/gin-gonic/gin"
	. "github.com/hanminggui/gin-exampel/moudles"
	"net/http"
	"strconv"
)

func GetOneShare(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	share := Share{Id: id}
	share.GetDetail()
	c.JSON(http.StatusOK, share)
}


func AddShare(c *gin.Context) {
	share := new(Share)
	err := c.BindJSON(share)
	Check(err)
	id, err := share.Add()
	Check(err)
	share.Id = id
	share.GetDetail()
	c.JSON(http.StatusOK, share)
}

func DeleteShare(c *gin.Context)  {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	share := Share{Id: id}
	share.Delete()
	c.JSON(http.StatusOK, share)
}
