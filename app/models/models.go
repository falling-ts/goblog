package models

import (
	apply "goblog/app"
	"goblog/pkg/types"
	"gorm.io/gorm"
	"time"
)

var (
	app    = apply.App
	router = app.Router
	db     = app.DB
)

func init() {
	migration(db)
}

// migration 自动迁移
func migration(db *gorm.DB) {

	// 自动迁移
	db.AutoMigrate(
		&User{},
		&Article{},
		&Category{},
	)
}

// BaseModel 模型基类
type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;not null"`

	CreatedAt time.Time `gorm:"column:created_at;index"`
	UpdatedAt time.Time `gorm:"column:updated_at;index"`
}

// GetStringID 获取 ID 的字符串格式
func (b BaseModel) GetStringID() string {
	return types.Uint64ToString(b.ID)
}
