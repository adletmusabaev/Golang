package middleware

import (
	"Assignment1_AdletMusabaev/internal/pkg/logger"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	logger := logger.NewLogger()
	return func(c *gin.Context) {
		logger.Info("Request: %s %s", c.Request.Method, c.Request.URL.Path)
		c.Next()
	}
}
