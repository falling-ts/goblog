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
	// 注册Web
	routes.RegisterWeb(router)

	// 中间件
	router.Use(middlewares.StartSession)
}
