package auth

import (
	"errors"
	"goblog/app/models"
	"goblog/pkg/session"

	"gorm.io/gorm"
)

var user = models.NewUser()

func _getUID() string {
	_uid := session.Get("uid")
	uid, ok := _uid.(string)
	if ok && len(uid) > 0 {
		return uid
	}
	return ""
}

// User 获取登录用户信息
func User() *models.User {
	uid := _getUID()
	if len(uid) > 0 {
		err := user.Get(uid)
		if err == nil {
			return user
		}
	}
	return &models.User{}
}

// Attempt 尝试登录
func Attempt(email string, password string) error {
	// 1. 根据 Email 获取用户
	err := user.GetByEmail(email)

	// 2. 如果出现错误
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("账号不存在或密码错误")
		} else {
			return errors.New("内部错误，请稍后尝试")
		}
	}

	// 3. 匹配密码
	if !user.ComparePassword(password) {
		return errors.New("账号不存在或密码错误")
	}

	// 4. 登录用户，保存会话
	session.Put("uid", user.GetStringID())

	return nil
}

// Login 登录指定用户
func Login(user *models.User) {
	session.Put("uid", user.GetStringID())
}

// Logout 退出用户
func Logout() {
	session.Forget("uid")
}

// Check 检测是否登录
func Check() bool {
	return len(_getUID()) > 0
}
