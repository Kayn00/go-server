package middlewares

import (
	"github.com/gin-gonic/gin"
	"runtime/debug"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				debug.PrintStack()
				c.JSON(500, err)
			}
		}()
		c.Next()
	}
}
