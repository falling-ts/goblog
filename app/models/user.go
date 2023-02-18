package models

import "goblog/pkg/logger"

// User 用户模型
type User struct {
	BaseModel

	Name     string `gorm:"column:name;type:varchar(255);not null;unique"`
	Email    string `gorm:"column:email;type:varchar(255);default:NULL;unique;"`
	Password string `gorm:"column:password;type:varchar(255)"`
}

// Create 创建用户，通过 User.ID 来判断是否创建成功
func (user *User) Create() (err error) {
	if err = db.Create(&user).Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}
