package routes

import (
	"github.com/gorilla/mux"
	"goblog/app/http/controllers"
	"net/http"
)

// RegisterWeb RegWeb 注册网页相关路由
func RegisterWeb(router *mux.Router) {
	// 静态页面
	pages := new(controllers.Page)
	router.HandleFunc("/", pages.Home).Methods("GET").Name("home")
	router.HandleFunc("/about", pages.About).Methods("GET").Name("about")
	router.NotFoundHandler = http.HandlerFunc(pages.NotFound)

	articles := new(controllers.Article)
	router.HandleFunc("/articles/{id:[0-9]+}", articles.Show).Methods("GET").Name("articles.show")
	router.HandleFunc("/articles", articles.Index).Methods("GET").Name("articles.index")
	router.HandleFunc("/articles", articles.Store).Methods("POST").Name("articles.store")
	router.HandleFunc("/articles/create", articles.Create).Methods("GET").Name("articles.create")
	router.HandleFunc("/articles/{id:[0-9]+}/edit", articles.Edit).Methods("GET").Name("articles.edit")
	router.HandleFunc("/articles/{id:[0-9]+}", articles.Update).Methods("POST").Name("articles.update")
	router.HandleFunc("/articles/{id:[0-9]+}/delete", articles.Delete).Methods("POST").Name("articles.delete")

	// 静态资源
	router.PathPrefix("/css/").Handler(http.FileServer(http.Dir("./public")))
	router.PathPrefix("/js/").Handler(http.FileServer(http.Dir("./public")))
}
