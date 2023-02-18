package routes

import (
	"github.com/gorilla/mux"
	"goblog/app/http/controllers"
	"net/http"
)

// RegisterWeb RegWeb 注册网页相关路由
func RegisterWeb(r *mux.Router) {

	// 静态页面
	pages := new(controllers.Pages)
	r.HandleFunc("/", pages.Home).Methods("GET").Name("home")
	r.HandleFunc("/about", pages.About).Methods("GET").Name("about")
	r.NotFoundHandler = http.HandlerFunc(pages.NotFound)

	articles := new(controllers.Articles)
	r.HandleFunc("/articles/{id:[0-9]+}", articles.Show).Methods("GET").Name("articles.show")
	r.HandleFunc("/articles", articles.Index).Methods("GET").Name("articles.index")
	r.HandleFunc("/articles", articles.Store).Methods("POST").Name("articles.store")
	r.HandleFunc("/articles/create", articles.Create).Methods("GET").Name("articles.create")
	r.HandleFunc("/articles/{id:[0-9]+}/edit", articles.Edit).Methods("GET").Name("articles.edit")
	r.HandleFunc("/articles/{id:[0-9]+}", articles.Update).Methods("POST").Name("articles.update")
	r.HandleFunc("/articles/{id:[0-9]+}/delete", articles.Delete).Methods("POST").Name("articles.delete")
}
