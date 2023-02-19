package controllers

import (
	"fmt"
	apply "goblog/app"
	"goblog/app/models"
	"goblog/pkg/flash"
	"goblog/pkg/logger"
	"gorm.io/gorm"
	"net/http"
)

var (
	app      = apply.App
	err      = app.Err
	article  = new(models.Article)
	user     = new(models.User)
	category = new(models.Category)
)

type Controllers struct{}

// ResponseForSQLError 处理 SQL 错误并返回
func (*Controllers) ResponseForSQLError(w http.ResponseWriter, err error) {
	if err == gorm.ErrRecordNotFound {
		// 3.1 数据未找到
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "404 文章未找到")
	} else {
		// 3.2 数据库错误
		logger.LogError(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "500 服务器内部错误")
	}
}

// ResponseForUnauthorized 处理未授权的访问
func (*Controllers) ResponseForUnauthorized(w http.ResponseWriter, r *http.Request) {
	flash.Warning("未授权操作！")
	http.Redirect(w, r, "/", http.StatusFound)
}
