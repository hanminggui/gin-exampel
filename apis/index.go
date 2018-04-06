package apis

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	)

func IndexApi(c *gin.Context) {
	log.Println("index")
	c.String(http.StatusOK, "hello")
}
