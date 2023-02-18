package models

import (
	"goblog/pkg/logger"
	"goblog/pkg/types"
	"strconv"
)

type Article struct {
	BaseModel
	ID    uint64
	Title string
	Body  string
}

func NewArticle() *Article {
	return &Article{}
}

// GetAll 获取全部文章
func (*Article) GetAll() ([]Article, error) {
	var articles []Article
	if err := db.Find(&articles).Error; err != nil {
		return articles, err
	}
	return articles, nil
}

// Get 通过 ID 获取文章
func (article *Article) Get(idStr string) error {
	id := types.StringToUint64(idStr)
	if err := db.First(article, id).Error; err != nil {
		return err
	}

	return nil
}

// Create 创建文章，通过 article.ID 来判断是否创建成功
func (article *Article) Create() (err error) {
	if err = db.Create(article).Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}

// Update 更新文章
func (article *Article) Update() (rowsAffected int64, err error) {
	result := db.Save(&article)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}

	return result.RowsAffected, nil
}

// Link 获取链接
func (article *Article) Link() string {
	showURL, err := router.Get("articles.show").URL("id", strconv.FormatUint(article.ID, 10))
	if err != nil {
		logger.LogError(err)
		return ""
	}
	return showURL.String()
}

// Delete 删除文章
func (article *Article) Delete() (rowsAffected int64, err error) {
	result := db.Delete(&article)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}

	return result.RowsAffected, nil
}
