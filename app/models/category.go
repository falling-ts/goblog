package models

import "goblog/pkg/logger"

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
