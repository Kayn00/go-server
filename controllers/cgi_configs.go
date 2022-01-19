package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	// 统一错误处理,无法recover处理逻辑有另起的goroutine,示例/panic
	//r.Use(middlewares.ErrorMiddleware())
	CGIConfig(r)
	return r
}

func CGIConfig(r *gin.Engine) {
	AddPprof(&r.RouterGroup)

	root := r.Group("/")
	{
		root.GET("/ping", Ping)
		root.GET("/panic", Panic)
	}
}
