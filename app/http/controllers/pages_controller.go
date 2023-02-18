package controllers

import (
	"fmt"
	"net/http"
)

// PagesController 处理静态页面
type Pages struct {
}

// Home 首页
func (*Pages) Home(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprint(w, "<h1>Hello, 欢迎来到 goblog！</h1>")
	if err != nil {
		return
	}
}

// About 关于我们页面
func (*Pages) About(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprint(w, "此博客是用以记录编程笔记，如您有反馈或建议，请联系 "+
		"<a href=\"mailto:summer@example.com\">summer@example.com</a>")
	if err != nil {
		return
	}
}

// NotFound 404 页面
func (*Pages) NotFound(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	_, err := fmt.Fprint(w, "<h1>请求页面未找到 :(</h1><p>如有疑惑，请联系我们。</p>")
	if err != nil {
		return
	}
}
