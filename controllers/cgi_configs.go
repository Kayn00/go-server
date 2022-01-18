package controllers

import "github.com/gin-gonic/gin"

func GetRouter() *gin.Engine {
	r := gin.Default()
	CGIConfig(r)
	return r
}

func CGIConfig(r *gin.Engine) {
	AddPprof(&r.RouterGroup)

	root := r.Group("/")
	{
		root.GET("/ping", Ping)
	}
}
