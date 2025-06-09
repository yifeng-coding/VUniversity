package model

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ResponseWrapper 通用响应格式
type ResponseWrapper struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func Response(c *gin.Context, err *BizError, data any) {
	if err != nil {
		Fail(c, err)
	} else {
		Success(c, data)
	}
}

// Success 成功响应
func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "成功！",
		"data":    data,
	})
}

func Fail(c *gin.Context, err *BizError) {
	c.JSON(http.StatusOK, gin.H{
		"code":    err.Code(),
		"message": err.Message(),
		"data":    nil,
	})
}
