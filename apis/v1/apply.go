package v1

import (
	"github.com/gin-gonic/gin"
	. "github.com/hanminggui/gin-exampel/moudles"
	"net/http"
)

func GetOneApply(c *gin.Context) {
	id := getInt64(c, "id")
	apply := Apply{Id: id}
	apply.GetDetail()
	if apply.Id == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, apply)
	}
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

