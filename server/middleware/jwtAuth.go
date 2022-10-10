package middleware

import (
	"fmt"
	"net/http"
	"server/helper"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		fmt.Println("token", authHeader)
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "无权限访问，请求未携带token",
			})
			c.Abort() //结束后续操作
			return
		}
		// log.Print("token:", authHeader)

		//按空格拆分
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "请求头中auth格式有误",
			})
			c.Abort()
			return
		}
		fmt.Println("分割后的token", parts[1])

		//解析token包含的信息
		claims, err := helper.AnalyseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}
		fmt.Println("解析后的token", claims)

		// if err := CheckUserInfo(claims); err != nil{
		// 	c.JSON(http.StatusOK, gin.H{
		// 		"code": -1,
		// 		"msg":  "用户名或密码错误",
		// 	})
		// 	c.Abort()
		// 	return
		// }

		// 将当前请求的claims信息保存到请求的上下文c上
		c.Set("claims", claims)
		c.Next() // 后续的处理函数可以用过c.Get("claims")来获取当前请求的用户信息
	}
}
