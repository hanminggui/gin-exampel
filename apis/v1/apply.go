package v1

import (
	"github.com/gin-gonic/gin"
	. "github.com/hanminggui/gin-exampel/moudles"
	"net/http"
	"strconv"
)

func GetOneApply(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	apply := Apply{Id: id}
	apply.GetDetail()
	c.JSON(http.StatusOK, apply)
}

func AddApply(c *gin.Context) {
	var apply Apply
	err := c.BindJSON(&apply)
	Check(err)
	id, err := apply.Add()
	Check(err)
	apply.Id = id
	apply.GetDetail()
	c.JSON(http.StatusOK, apply)
}
