package middlewares

import (
	"goblog/app/http/controllers"
	"net/http"
)

// Initial 开启 session 会话控制
func Initial(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		controllers.Initial()

		// 2. . 继续处理接下去的请求
		next.ServeHTTP(w, r)
	})
}
