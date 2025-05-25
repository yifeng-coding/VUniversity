package middlewire

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"time"
)

// 自定义响应写入器
type responseBodyWriter struct {
	gin.ResponseWriter               // 嵌入原始写入器
	body               *bytes.Buffer // 新增缓冲区保存响应内容
}

func (r *responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)                  // 保存到缓冲区
	return r.ResponseWriter.Write(b) // 调用原始写入器发送响应
}

func (r *responseBodyWriter) WriteString(s string) (int, error) {
	r.body.WriteString(s)
	return r.ResponseWriter.WriteString(s)
}

func ZapLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// 记录请求体（c.Request.Body 是一个io.ReadCloser，读取后会被消耗（指针移到末尾）。当中间件读取了请求体，后续的处理函数（如 c.ShouldBindJSON()）将无法再次读取，因此读取后需要恢复）
		var reqBody string
		if c.Request.Body != nil {
			bodyBytes, _ := io.ReadAll(c.Request.Body)
			reqBody = string(bodyBytes)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // 恢复请求体
		}
		// 记录请求信息
		zap.L().Info("Request",
			zap.String("path", c.Request.URL.Path),
			zap.String("method", c.Request.Method),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("body", reqBody),
		)

		// 使用自定义响应写入器替代原始写入器（调用c.JSON()时，响应内容会直接写入网络连接，无法事后获取，因此自定义写入器，将数据缓存一份，方便打印日志）
		writer := &responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer

		// 处理请求
		c.Next()

		// 记录响应信息
		zap.L().Info("Response",
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.String("body", writer.body.String()),
			zap.String("latency", time.Since(start).Round(time.Millisecond).String()),
		)
	}
}
