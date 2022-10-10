package router

import (
	v1 "server/api/v1"
	_ "server/docs"
	"server/logger"
	"server/middleware"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	// 跨域配置
	r.Use(middleware.CORSMiddleware())
	// swage配置
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.POST("/register", v1.UserRegisterAPI)
	// 用户相关
	r.POST("/api/v1/users/login", v1.LoginAPI)
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.DELETE("/user", v1.UserDeleteAPI)
	// 配置路由规则
	// r.GET("/ping", service.Ping)

	r.Use(middleware.JWTAuth())
	r.POST("/api/v1/users/info", v1.UserInfoAPI)
	return r
}
