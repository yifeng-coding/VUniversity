package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yifeng-coding/VUniversity/config"
	"github.com/yifeng-coding/VUniversity/dal/mysql"
	"github.com/yifeng-coding/VUniversity/dal/redis"
	_ "github.com/yifeng-coding/VUniversity/docs"
	"github.com/yifeng-coding/VUniversity/middleware"
	"github.com/yifeng-coding/VUniversity/router"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//	@title			微学堂
//	@version		1.0
//	@description	本项目是一个前后端分离的多用户论坛项目，实现了用户注册、登录、发帖、评论、私信、点赞、关注、搜索、记录日志、敏感词过滤等功能。
//	@contact.name	微信公众号：一枫说码

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description				使用格式: Bearer {token}
func main() {
	fmt.Println(`
		██╗   ██╗██╗   ██╗███╗   ██╗██╗████████╗███████╗██████╗ ███████╗██╗████████╗██╗   ██╗
		██║   ██║██║   ██║████╗  ██║██║╚══██╔══╝██╔════╝██╔══██╗██╔════╝██║╚══██╔══╝╚██╗ ██╔╝
		██║   ██║██║   ██║██╔██╗ ██║██║   ██║   █████╗  ██████╔╝███████╗██║   ██║    ╚████╔╝ 
		╚██╗ ██╔╝██║   ██║██║╚██╗██║██║   ██║   ██╔══╝  ██╔══██╗╚════██║██║   ██║     ╚██╔╝  
		 ╚████╔╝ ╚██████╔╝██║ ╚████║██║   ██║   ███████╗██║  ██║███████║██║   ██║      ██║   
		  ╚═══╝   ╚═════╝ ╚═╝  ╚═══╝╚═╝   ╚═╝   ╚══════╝╚═╝  ╚═╝╚══════╝╚═╝   ╚═╝      ╚═╝
	`)
	// 初始化配置
	config.GetConfig()
	// 初始化DB
	mysql.GetDB()
	// 初始化Redis
	redis.GetRedis()
	// 初始化日志（同时输出到控制台和文件）
	logger, _ := zap.Config{
		Level:    zap.NewAtomicLevelAt(zap.InfoLevel), // 设置日志级别
		Encoding: "json",
		OutputPaths: []string{
			"app.log", // 日志文件路径
			"stdout",  // 同时输出到控制台
		},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:     "time",
			LevelKey:    "level",
			MessageKey:  "msg",
			EncodeTime:  zapcore.ISO8601TimeEncoder, // 时间格式
			EncodeLevel: zapcore.LowercaseLevelEncoder,
		},
	}.Build()
	zap.ReplaceGlobals(logger) // 设置全局Logger
	defer zap.L().Sync()

	// 初始化gin框架
	r := gin.New()
	r.Use(middleware.CORSMiddleware()) // 使用跨域中间件
	r.Use(middleware.ZapLogger())      // 使用zap日志中间件
	r.Use(gin.Recovery())              // 使用恢复中间件
	router.RegisterRouter(r)
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
