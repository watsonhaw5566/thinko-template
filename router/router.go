package router

import (
	"github.com/gin-gonic/gin"
	"think-go/app/controller"
	"think-go/middleware"
)

// BindController 注册路由,绑定控制器,绑定中间件
func BindController(router *gin.Engine) *gin.Engine {
	// 跨域中间件
	router.Use(middleware.Cors())
	// 日志中间件
	router.Use(middleware.Logger())

	hello := router.Group("/api/v1")
	{
		hello.GET("/hello", controller.SayHello)
	}

	return router
}
