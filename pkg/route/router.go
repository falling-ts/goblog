package route

import (
	"github.com/gorilla/mux"
	apply "goblog/app"
	"goblog/pkg/config"
	"goblog/pkg/logger"
	"net/http"
)

var (
	app    = apply.App
	router = app.Router
)

// Name2URL 通过路由名称来获取 URL
func Name2URL(routeName string, pairs ...string) string {
	url, err := router.Get(routeName).URL(pairs...)
	if err != nil {
		logger.LogError(err)
		return ""
	}

	return config.GetString("app.url") + url.String()
}

// GetRouteVariable 获取 URI 路由参数
func GetRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}
