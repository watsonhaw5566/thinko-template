package router

import (
	"github.com/think-go/tg"
	"think-go/app/controller"
	"think-go/middleware"
)

// BindController 注册路由,绑定控制器,绑定中间件
func BindController(engine *tg.Engine) {
	// 跨域中间件
	engine.Use(middleware.Cors)
	// 日志中间件
	engine.Use(middleware.Logger)

	// 首页宣传页
	engine.GET("/", controller.HomeView)

	// 分组路由
	router := engine.Group("/api/v1")
	{
		router.GET("hello", controller.SayHello)
	}
}
