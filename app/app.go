package application

import (
	"database/sql"
	"github.com/gorilla/mux"
	"goblog/pkg/db"
)

type Application struct {
	Router *mux.Router
	DB     *sql.DB
}

var App *Application

// init 初始化
func init() {
	App = &Application{
		Router: mux.NewRouter(),
		DB:     db.InitDB(),
	}
}
