package apply

import (
	"embed"
	"github.com/gorilla/mux"
	"goblog/app/error"
	"goblog/pkg/db"
	"goblog/public"
	"goblog/resources"
	"gorm.io/gorm"
)

type Apply struct {
	Router   *mux.Router
	DB       *gorm.DB
	Err      *error.Error
	TplFS    embed.FS
	StaticFS embed.FS
}

var App *Apply

// init 初始化
func init() {
	App = &Apply{
		Router:   mux.NewRouter(),
		DB:       db.InitGormDB(),
		Err:      error.NewError(),
		TplFS:    resources.TplFS,
		StaticFS: public.StaticFS,
	}
}
