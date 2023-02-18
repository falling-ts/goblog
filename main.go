package main

import _ "goblog/config"
import (
	apply "goblog/app"
	"goblog/app/http/middlewares"
	_ "goblog/bootstrap"
	"net/http"
)

var (
	app = apply.App
	err = app.Err
)

func main() {
	err.Throw(http.ListenAndServe(":3000", middlewares.RemoveSlash(app.Router)))
}
