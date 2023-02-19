package controllers

import (
	apply "goblog/app"
	"goblog/app/models"
)

var (
	app     = apply.App
	err     = app.Err
	article = new(models.Article)
	user    = new(models.User)
)
