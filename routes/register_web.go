package routes

import (
	"github.com/gorilla/mux"
	"goblog/app/http/controllers"
	"net/http"
)

// RegisterWeb RegWeb 注册网页相关路由
func RegisterWeb(router *mux.Router) {
	// 静态页面
	page := new(controllers.Page)
	router.HandleFunc("/", page.Home).Methods("GET").Name("home")
	router.HandleFunc("/about", page.About).Methods("GET").Name("about")
	router.NotFoundHandler = http.HandlerFunc(page.NotFound)

	article := new(controllers.Article)
	router.HandleFunc("/articles/{id:[0-9]+}", article.Show).Methods("GET").Name("articles.show")
	router.HandleFunc("/articles", article.Index).Methods("GET").Name("articles.index")
	router.HandleFunc("/articles", article.Store).Methods("POST").Name("articles.store")
	router.HandleFunc("/articles/create", article.Create).Methods("GET").Name("articles.create")
	router.HandleFunc("/articles/{id:[0-9]+}/edit", article.Edit).Methods("GET").Name("articles.edit")
	router.HandleFunc("/articles/{id:[0-9]+}", article.Update).Methods("POST").Name("articles.update")
	router.HandleFunc("/articles/{id:[0-9]+}/delete", article.Delete).Methods("POST").Name("articles.delete")

	// 静态资源
	router.PathPrefix("/css/").Handler(http.FileServer(http.Dir("./public")))
	router.PathPrefix("/js/").Handler(http.FileServer(http.Dir("./public")))

	// 用户认证
	auth := new(controllers.Auth)
	router.HandleFunc("/auth/register", auth.Register).Methods("GET").Name("auth.register")
	router.HandleFunc("/auth/do-register", auth.DoRegister).Methods("POST").Name("auth.doregister")
	router.HandleFunc("/auth/login", auth.Login).Methods("GET").Name("auth.login")
	router.HandleFunc("/auth/dologin", auth.DoLogin).Methods("POST").Name("auth.dologin")
}
