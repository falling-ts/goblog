package routes

import (
	"github.com/gorilla/mux"
	"goblog/app/http/controllers"
	"goblog/app/http/middlewares"
	"net/http"
)

// RegisterWeb RegWeb 注册网页相关路由
func RegisterWeb(router *mux.Router) {
	// 静态页面
	page := new(controllers.Page)
	router.HandleFunc("/about", page.About).Methods("GET").Name("about")
	router.NotFoundHandler = http.HandlerFunc(page.NotFound)

	article := new(controllers.Article)
	router.HandleFunc("/", article.Index).Methods("GET").Name("articles.index")
	router.HandleFunc("/articles/{id:[0-9]+}", middlewares.Auth(article.Show)).Methods("GET").Name("articles.show")
	router.HandleFunc("/articles", middlewares.Auth(article.Store)).Methods("POST").Name("articles.store")
	router.HandleFunc("/articles/create", middlewares.Auth(article.Create)).Methods("GET").Name("articles.create")
	router.HandleFunc("/articles/{id:[0-9]+}/edit", middlewares.Auth(article.Edit)).Methods("GET").Name("articles.edit")
	router.HandleFunc("/articles/{id:[0-9]+}", middlewares.Auth(article.Update)).Methods("POST").Name("articles.update")
	router.HandleFunc("/articles/{id:[0-9]+}/delete", middlewares.Auth(article.Delete)).Methods("POST").Name("articles.delete")

	// 静态资源
	router.PathPrefix("/css/").Handler(http.FileServer(http.Dir("./public")))
	router.PathPrefix("/js/").Handler(http.FileServer(http.Dir("./public")))

	// 用户认证
	auth := new(controllers.Auth)
	router.HandleFunc("/auth/register", middlewares.Guest(auth.Register)).Methods("GET").Name("auth.register")
	router.HandleFunc("/auth/do-register", middlewares.Guest(auth.DoRegister)).Methods("POST").Name("auth.doregister")
	router.HandleFunc("/auth/login", middlewares.Guest(auth.Login)).Methods("GET").Name("auth.login")
	router.HandleFunc("/auth/dologin", middlewares.Guest(auth.DoLogin)).Methods("POST").Name("auth.dologin")
	router.HandleFunc("/auth/logout", middlewares.Auth(auth.Logout)).Methods("POST").Name("auth.logout")

	// 用户相关
	user := new(controllers.User)
	router.HandleFunc("/users/{id:[0-9]+}", user.Show).Methods("GET").Name("users.show")

	cate := new(controllers.Category)
	router.HandleFunc("/categories/create", middlewares.Auth(cate.Create)).Methods("GET").Name("categories.create")
	router.HandleFunc("/categories", middlewares.Auth(cate.Store)).Methods("POST").Name("categories.store")
	router.HandleFunc("/categories", middlewares.Auth(cate.Store)).Methods("POST").Name("categories.store")
	router.HandleFunc("/categories/{id:[0-9]+}", cate.Show).Methods("GET").Name("categories.show")
}
