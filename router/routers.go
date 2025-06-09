package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/yifeng-coding/VUniversity/handler"
	"github.com/yifeng-coding/VUniversity/middleware"
	"net/http"
)

func RegisterRouter(r *gin.Engine) {
	// 测试HelloWorld
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello World")
	})
	// swagger接口文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 帖子相关
	post := r.Group("post")
	post.GET("/list", handler.GetPostList)

	// 用户相关
	user := r.Group("user")
	user.POST("/register", handler.CreateUser)
	user.GET("/register_code", handler.GetRegisterVerifyCode)
	user.POST("/login", handler.UserLogin)
	user.POST("/refresh_token", handler.GetNewTokenByRefreshToken)
	user.GET("/info", middleware.AuthMiddleware(), handler.GetUserInfo)                // 需要鉴权
	user.POST("/update/info", middleware.AuthMiddleware(), handler.UpdateUser)         // 需要鉴权
	user.POST("/update/password", middleware.AuthMiddleware(), handler.UpdatePassword) // 需要鉴权
	user.POST("/upload/avatar", middleware.AuthMiddleware(), handler.UploadAvatar)     // 需要鉴权

	// 注册静态文件服务(用于上传头像)
	r.Static("/static", "./static")
}
