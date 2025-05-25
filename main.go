package main

import (
	"github.com/SniperCoding/VUniversity/config"
	"github.com/SniperCoding/VUniversity/dal/mysql"
	"github.com/SniperCoding/VUniversity/middlewire"
	"github.com/SniperCoding/VUniversity/router"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func init() {

}

func main() {
	// 初始化配置
	config.GetConfig()
	// 初始化DB
	mysql.GetDB()
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

	r := gin.New()
	r.Use(middlewire.ZapLogger()) // 使用zap日志中间件
	r.Use(gin.Recovery())         // 使用恢复中间件
	router.RegisterRouter(r)
	if err := r.Run("127.0.0.1:8080"); err != nil {
		panic(err)
	}
}
