package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/yifeng-coding/VUniversity/config"
	"sync"
)

// GetRedis 获取redis连接，使用sync.Once保证单例
var GetRedis = sync.OnceValue(func() *redis.Client {
	redisConf := config.GetConfig().Redis

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConf.Host, redisConf.Port),
		Password: redisConf.Password,
		DB:       redisConf.DB,
	})
	// 测试连接
	if _, err := client.Ping(client.Context()).Result(); err != nil {
		panic("Error to Redis connection, errs: " + err.Error())
	}
	return client
})
