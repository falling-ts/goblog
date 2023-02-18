package main

import (
	application "goblog/app"
	_ "goblog/bootstrap"
	"net/http"
)

var app = application.App

func main() {
	err := http.ListenAndServe(":3000", app.Router)
	if err != nil {
		return
	}
}
