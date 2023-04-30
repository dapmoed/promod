package pm_logger

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Logger interface {
	Info(msg string, fields ...interface{})
	Error(msg string, fields ...interface{})
}

func GinMiddleware(logger Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("request received", zap.String("path", c.Request.URL.Path))
		c.Next()
	}
}
