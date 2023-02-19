package main

import (
	_ "goblog/config"
	"goblog/pkg/view"
)
import (
	"embed"
	apply "goblog/app"
	"goblog/app/http/middlewares"
	_ "goblog/bootstrap"
	"net/http"
)

var (
	app = apply.App
	err = app.Err
)

//go:embed resources/views/articles/*
//go:embed resources/views/auth/*
//go:embed resources/views/categories/*
//go:embed resources/views/layouts/*
var tplFS embed.FS

func main() {
	app.SetTplFS(tplFS)
	view.InitTplFS()
	err.Throw(http.ListenAndServe(":3000", middlewares.RemoveSlash(app.Router)))
}
