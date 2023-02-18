package bootstrap

import (
	application "goblog/app"
	"goblog/app/http/middlewares"
	"goblog/routes"
)

var (
	app    = application.App
	router = app.Router
)

func init() {
	// 前置中间件
	router.Use(middlewares.RemoveSlash)

	// 注册Web
	routes.RegisterWeb(router)

	// 后置中间件
	router.Use(middlewares.ForceHTML)
}
