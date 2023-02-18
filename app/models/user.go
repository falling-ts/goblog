package models

import (
	"goblog/pkg/logger"
	"goblog/pkg/types"
)

// User 用户模型
type User struct {
	BaseModel

	Name     string `gorm:"type:varchar(255);not null;unique" valid:"name"`
	Email    string `gorm:"type:varchar(255);unique;" valid:"email"`
	Password string `gorm:"type:varchar(255)" valid:"password"`

	// gorm:"-" —— 设置 GORM 在读写时略过此字段，仅用于表单验证
	PasswordConfirm string `gorm:"-" valid:"password_confirm"`
}

func NewUser() *User {
	return &User{}
}

// Create 创建用户，通过 User.ID 来判断是否创建成功
func (user *User) Create() (err error) {
	if err = db.Create(&user).Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}

func (user *User) Get(idstr string) error {
	id := types.StringToUint64(idstr)
	if err := db.First(user, id).Error; err != nil {
		return err
	}
	return nil
}

// GetByEmail 通过 Email 来获取用户
func (user *User) GetByEmail(email string) error {
	if err := db.Where("email = ?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

// ComparePassword 对比密码是否匹配
func (user *User) ComparePassword(password string) bool {
	return user.Password == password
}
