package controllers

import "github.com/gin-gonic/gin"

// @Tags default
// @Summary ping
// @Param tradeCode query string true "交易代码"
// @Success 200 {string} pong
// @Failure 400 "retcode：10002 参数错误；10003 校验错误"
// @Failure 500 "retcode：20001 服务错误；20002 接口错误；20003 无数据错误；20004 数据库异常；20005 缓存异常 重复: 20006"
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(200, "pong")
}
