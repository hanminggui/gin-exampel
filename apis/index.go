package apis

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"github.com/huandu/goroutine"
	"time"
)

func IndexApi(c *gin.Context) {
	log.Println(c.Writer.Header().Get("X-Request-Id"), "index")
	time.Sleep(3 * time.Second)
	c.JSON(http.StatusOK, gin.H{"id": goroutine.GoroutineId()})
	//c.String(http.StatusOK, "hello")
}
