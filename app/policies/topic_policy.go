package policies

import (
	"goblog/app/models"
	"goblog/pkg/auth"
)

// CanModifyArticle 是否允许修改话题
func CanModifyArticle(_article models.Article) bool {
	return auth.User().ID == _article.UserID
}
