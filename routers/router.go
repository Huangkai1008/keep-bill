package routers

import (
	"github.com/gin-gonic/gin"
	"keep-bill/pkg/setting"
)

func InitRouter() *gin.Engine {
	// 初始化路由
	r := gin.New()
	r.Use(gin.Logger()) // 添加日志
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	userApi := r.Group("/user")
	{
		// 注册
		userApi.POST("/register")

	}

	return r
}
