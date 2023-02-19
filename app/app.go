package apply

import (
	"embed"
	"github.com/gorilla/mux"
	"goblog/app/error"
	"goblog/pkg/db"
	"gorm.io/gorm"
)

type Apply struct {
	Router *mux.Router
	DB     *gorm.DB
	Err    *error.Error
	TplFS  embed.FS
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

// SetTplFS 设置数据
func (app *Apply) SetTplFS(tplFS embed.FS) {
	app.TplFS = tplFS
}
