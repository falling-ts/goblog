package controllers

import apply "goblog/app"

var (
	app    = apply.App
	router = app.Router
	db     = app.DB
	err    = app.Err
)
