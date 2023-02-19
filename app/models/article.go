package models

import (
	"goblog/pkg/logger"
	"goblog/pkg/pagination"
	"goblog/pkg/route"
	"goblog/pkg/types"
	"net/http"
	"strconv"
)

type Article struct {
	BaseModel
	Title  string `gorm:"type:varchar(255);not null;" valid:"title"`
	Body   string `gorm:"type:longtext;not null;" valid:"body"`
	UserID uint64 `gorm:"not null;index"`
	User   User
}

// GetAll 获取全部文章
func (*Article) GetAll(r *http.Request, perPage int) ([]Article, pagination.ViewData, error) {
	// 1. 初始化分页实例
	db := db.Model(Article{}).Order("created_at desc")
	_pager := pagination.New(r, db, route.Name2URL("articles.index"), perPage)

	// 2. 获取视图数据
	viewData := _pager.Paging()

	// 3. 获取数据
	var articles []Article
	_pager.Results(&articles)

	return articles, viewData, nil
}

// Get 通过 ID 获取文章
func (article *Article) Get(idStr string) error {
	id := types.StringToUint64(idStr)
	if err := db.Preload("User").First(article, id).Error; err != nil {
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

// CreatedAtDate 创建日期
func (article *Article) CreatedAtDate() string {
	return article.CreatedAt.Format("2006-01-02")
}

// GetByUserID 获取全部文章
func (*Article) GetByUserID(uid string) ([]Article, error) {
	var articles []Article
	if err := db.Where("user_id = ?", uid).Preload("User").Find(&articles).Error; err != nil {
		return articles, err
	}
	return articles, nil
}
