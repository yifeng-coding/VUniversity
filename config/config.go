package config

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

type Config struct {
	MySQL struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Database string `json:"database"`
		Username string `json:"username"`
		Password string `json:"password"`
		Charset  string `json:"charset"`
		Loc      string `json:"loc"`
	}
}

// GetConfig 获取配置，使用sync.Once保证单例
var GetConfig = sync.OnceValue(func() *Config {
	// 设置配置文件位置
	viper.SetConfigFile("./config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("get config failed, errs:%v", err))
	}
	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		panic(fmt.Errorf("unmarshal to Conf failed, errs:%v", err))
	}
	return &c
})
