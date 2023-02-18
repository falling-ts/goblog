package bootstrap

import (
	apply "goblog/app"
	"goblog/app/http/middlewares"
	"goblog/routes"
)

var (
	app    = apply.App
	router = app.Router
)

func init() {
	initRouter()
}

// initRouter 初始化路由
func initRouter() {
	// 前置中间件
	router.Use(middlewares.RemoveSlash)

	// 注册Web
	routes.RegisterWeb(router)

	// 后置中间件
	router.Use(middlewares.ForceHTML)
}
