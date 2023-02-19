package models

import (
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"goblog/pkg/types"
)

// Category 文章分类
type Category struct {
	BaseModel

	Name string `gorm:"type:varchar(255);not null;" valid:"name"`
}

// Create 创建分类，通过 category.ID 来判断是否创建成功
func (category *Category) Create() (err error) {
	if err = db.Create(category).Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}

// All 获取分类数据
func (*Category) All() ([]Category, error) {
	var categories []Category
	if err := db.Find(&categories).Error; err != nil {
		return categories, err
	}
	return categories, nil
}

// Link 方法用来生成文章链接
func (category *Category) Link() string {
	return route.Name2URL("categories.show", "id", category.GetStringID())
}

// Get 通过 ID 获取分类
func (category *Category) Get(idstr string) error {
	id := types.StringToUint64(idstr)
	if err := db.First(category, id).Error; err != nil {
		return err
	}

	return nil
}
