package v1

import (
	"github.com/gin-gonic/gin"
	"strconv"
	. "github.com/hanminggui/gin-exampel/moudles"
)

func getInt64(c *gin.Context, name string) (int64) {
	number, err := strconv.ParseInt(c.Param(name), 10, 64)
	Check(err)
	return number
}