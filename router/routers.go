package router

import (
	"github.com/SniperCoding/VUniversity/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRouter(r *gin.Engine) {
	// 测试HelloWorld
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello World")
	})

	// 帖子相关
	post := r.Group("post")
	post.GET("/list", handler.GetPostList)
}
