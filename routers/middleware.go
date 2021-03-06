package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"net/http"
	"time"
	"fmt"
	"runtime"
	"strconv"
)

func middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid4, _ := uuid.NewV4()
		c.Writer.Header().Set("X-Request-Id", uuid4.String())
		c.Next()
		if !c.Writer.Written() {
			c.JSON(http.StatusOK, gin.H{ "message": "ok"})
		}
	}
}

func timeOut() gin.HandlerFunc {
	return func(c *gin.Context) {
		ch := make(chan int)
		go func(c1 chan int) {
			time.Sleep(8 * time.Second)
			if !c.Writer.Written() && !c.IsAborted() {
				notTimeout,exit := c.Get("not_timeout")
				if !exit || notTimeout != true {
					c1 <- 1
				}
			}
		}(ch)
		go func(c1 chan int) {
			c.Next()
			//if c.IsAborted() { return }
			c1 <- 0
		}(ch)
		fmt.Println(runtime.NumGoroutine())
		if 1 == <- ch{
			fmt.Println(runtime.NumGoroutine())
			c.AbortWithStatus(http.StatusGatewayTimeout)
			return
		}
	}
}

func clearTimeOut() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("not_timeout", true)
		c.Next()
	}
}

func needPager() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "GET" {
			limit,_ := strconv.Atoi(c.Query("limit"))
			offset,_ := strconv.Atoi(c.Query("offset"))
			if limit == 0 {
				limit = 20
			}
			c.Set("limit", limit)
			c.Set("offset", offset)
		}
		c.Next()
	}
}