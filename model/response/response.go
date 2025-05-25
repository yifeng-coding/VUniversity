package response

import (
	"github.com/SniperCoding/VUniversity/model/errs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(c *gin.Context, err *errs.BizError, data any) {
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

func Fail(c *gin.Context, err *errs.BizError) {
	c.JSON(http.StatusOK, gin.H{
		"code":    err.Code(),
		"message": err.Message(),
		"data":    nil,
	})
}
