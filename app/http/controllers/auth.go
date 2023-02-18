package controllers

import (
	"fmt"
	"goblog/app/models"
	"goblog/pkg/view"
	"net/http"
)

// Auth 处理用户认证
type Auth struct{}

// Register 注册页面
func (*Auth) Register(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.register")
}

// DoRegister 处理注册逻辑
func (*Auth) DoRegister(w http.ResponseWriter, r *http.Request) {
	// 0. 初始化变量
	name := r.PostFormValue("name")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	// 1. 表单验证
	// 2. 验证通过 —— 入库，并跳转到首页
	_user := models.User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	_user.Create()

	if _user.ID > 0 {
		fmt.Fprint(w, "插入成功，ID 为"+_user.GetStringID())
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "创建用户失败，请联系管理员")
	}
}
