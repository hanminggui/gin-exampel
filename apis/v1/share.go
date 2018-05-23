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

func GetShareList(c *gin.Context)  {
	//tp := c.Query("type")

	limit,_ := strconv.Atoi(c.Query("limit"))
	offset,_ := strconv.Atoi(c.Query("offset"))
	c.JSON(http.StatusOK, GetShares(limit, offset))
}

func ApplyAuth(c *gin.Context)  {
	shareId, _ := strconv.ParseInt(c.Param("shareId"), 10, 64)
	applyId, _ := strconv.ParseInt(c.Param("applyId"), 10, 64)
	// 鉴权
	if shareId > 0 && applyId > 0{

	}
	apply := Apply{Id: applyId}
	tp := c.PostForm("type")
	if "pass" == tp {
		apply.Pass()
	} else if "down" == tp {
		apply.Down()
	}
	apply.GetDetail()
	c.JSON(http.StatusOK, apply)
}