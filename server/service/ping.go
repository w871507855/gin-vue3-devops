package service

import "github.com/gin-gonic/gin"

// @Tags ping方法
// @Summary ping
// @accept application/json
// @Produce application/json
// @Success 200 {string} json "{"code": "200", "data": "", "message"}"
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    200,
		"message": "pong",
		"data":    "",
	})
}
