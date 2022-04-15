package middleware

import (
	"github.com/gin-gonic/gin"
	"ippool_api/utils/log"
	"time"
)

func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		end := time.Now()
		//执行时间
		latency := end.Sub(start)
		log.GlobalLog.Infof("%s | %s | %s | %d | %s | %d | %s", c.ClientIP(), c.Request.Method, c.Request.RequestURI, c.Writer.Status(), c.Request.Proto, latency, c.Request.UserAgent())
	}
}
