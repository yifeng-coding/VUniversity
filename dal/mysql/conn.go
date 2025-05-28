package mysql

import (
	"fmt"
	"github.com/yifeng-coding/VUniversity/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"net/url"
	"sync"
)

// GetDB 获取数据库连接，使用sync.Once保证单例
var GetDB = sync.OnceValue(func() *gorm.DB {
	mySQLConf := config.GetConfig().MySQL
	// 用户名:密码@tcp(IP:Port)/库名?参数
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=%s",
		mySQLConf.Username, mySQLConf.Password, mySQLConf.Host, mySQLConf.Port, mySQLConf.Database, mySQLConf.Charset, url.QueryEscape(mySQLConf.Loc))
	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 设置表名前缀及禁用表名复数
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
		},
	})
	if err != nil {
		panic("Error to Db connection, errs: " + err.Error())
	}
	return db
})
