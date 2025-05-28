package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/yifeng-coding/VUniversity/handler"
	"net/http"
)

func RegisterRouter(r *gin.Engine) {
	// 测试HelloWorld
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello World")
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 帖子相关
	post := r.Group("post")
	post.GET("/list", handler.GetPostList)
}
