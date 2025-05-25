package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRouter(r *gin.Engine) {
	// 测试HelloWorld
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello World")
	})
}
