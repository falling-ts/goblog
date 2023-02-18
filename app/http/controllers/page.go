package controllers

import (
	"fmt"
	"net/http"
)

// Page 处理静态页面
type Page struct{}

// Home 首页
func (*Page) Home(res http.ResponseWriter, _ *http.Request) {
	err.Throw(fmt.Fprint(res, "<h1>Hello, 欢迎来到 goblog！</h1>"))
}

// About 关于我们页面
func (*Page) About(res http.ResponseWriter, _ *http.Request) {
	err.Throw(fmt.Fprint(res, "此博客是用以记录编程笔记，如您有反馈或建议，请联系 "+
		"<a href=\"mailto:summer@example.com\">summer@example.com</a>"))
}

// NotFound 404 页面
func (*Page) NotFound(res http.ResponseWriter, _ *http.Request) {
	res.WriteHeader(http.StatusNotFound)
	err.Throw(fmt.Fprint(res, "<h1>请求页面未找到 :(</h1><p>如有疑惑，请联系我们。</p>"))
}
