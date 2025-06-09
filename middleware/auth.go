package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/yifeng-coding/VUniversity/model"
	"github.com/yifeng-coding/VUniversity/util"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取token
		token := c.GetHeader("Authorization")
		if token == "" {
			c.Abort()
			model.Fail(c, model.TokenInvalid.WithMessage("未认证"))
			return
		}
		// 验证 Token 格式
		tokenStr := ""
		fields := strings.Fields(token)
		if len(fields) >= 2 && fields[0] == "Bearer" {
			tokenStr = fields[1]
		} else {
			c.Abort()
			model.Fail(c, model.TokenInvalid.WithMessage("格式无效"))
			return
		}
		// 解析 Token
		claims, err := util.ParseToken(tokenStr)
		if err != nil {
			c.Abort()
			model.Fail(c, model.TokenInvalid.WithMessage("内容无效"))
			return
		}
		// 将用户信息存入上下文
		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
