package main

import (
	apply "goblog/app"
	_ "goblog/bootstrap"
	"net/http"
)

var (
	app = apply.App
	err = app.Err
)

func main() {
	err.Throw(http.ListenAndServe(":3000", app.Router))
}
