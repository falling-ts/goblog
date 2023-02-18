package controllers

import (
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
	//
}
