package apply

import (
	"github.com/gorilla/mux"
	"goblog/app/error"
	"goblog/pkg/db"
	"gorm.io/gorm"
)

type Apply struct {
	Router *mux.Router
	DB     *gorm.DB
	Err    *error.Error
}

var App *Apply

// init 初始化
func init() {
	App = &Apply{
		Router: mux.NewRouter(),
		DB:     db.InitGormDB(),
		Err:    error.NewError(),
	}
}
