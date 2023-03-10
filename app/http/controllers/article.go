package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"goblog/app/models"
	"goblog/app/policies"
	"goblog/app/requests"
	"goblog/pkg/auth"
	"goblog/pkg/flash"
	"goblog/pkg/route"
	"goblog/pkg/view"
	"net/http"
	"unicode/utf8"
)

// Article 文章相关页面
type Article struct {
	Controllers
}

type ArticleFormData struct {
	Title, Body string
	Article     *models.Article
	Errors      map[string]string
}

// Index 获取文件列表
func (art *Article) Index(w http.ResponseWriter, r *http.Request) {
	// 1. 获取结果集
	articles, pagerData, err := article.GetAll(r, 2)

	if err != nil {
		art.ResponseForSQLError(w, err)
	} else {
		// 2. 加载模板
		// ---  2. 加载模板 ---
		view.Render(w, view.D{
			"Articles":  articles,
			"PagerData": pagerData,
		}, "articles.index", "articles._article_meta")
	}
}

// Show 文章详情页面
func (art *Article) Show(w http.ResponseWriter, r *http.Request) {
	// 1. 获取 URL 参数
	vars := mux.Vars(r)
	id := vars["id"]

	// 2. 读取对应的文章数据
	err := article.Get(id)

	// 3. 如果出现错误
	if err != nil {
		art.ResponseForSQLError(w, err)
	} else {
		// 4. 读取成功，显示文章
		view.Render(w, view.D{
			"Article":          article,
			"CanModifyArticle": policies.CanModifyArticle(*article),
		}, "articles.show", "articles._article_meta")
	}
}

// Create 创建文章页面
func (*Article) Create(w http.ResponseWriter, r *http.Request) {
	view.Render(w, view.D{}, "articles.create", "articles._form_field")
}

// Store 保存文章
func (*Article) Store(w http.ResponseWriter, r *http.Request) {
	currentUser := auth.User()
	article := &models.Article{
		Title:  r.PostFormValue("title"),
		Body:   r.PostFormValue("body"),
		UserID: currentUser.ID,
	}

	errors := requests.ValidateArticleForm(*article)

	// 检查是否有错误
	if len(errors) == 0 {
		// 创建文章
		article.Create()
		if article.ID > 0 {
			indexURL := route.Name2URL("articles.show", "id", article.GetStringID())
			http.Redirect(w, r, indexURL, http.StatusFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "创建文章失败，请联系管理员")
		}
	} else {
		view.Render(w, view.D{
			"Article": article,
			"Errors":  errors,
		}, "articles.create", "articles._form_field")
	}
}

// Edit 修改文章页面
func (art *Article) Edit(w http.ResponseWriter, r *http.Request) {
	// 1. 获取 URL 参数
	id := route.GetRouteVariable("id", r)

	// 2. 读取对应的文章数据
	err := article.Get(id)

	// 3. 如果出现错误
	if err != nil {
		art.ResponseForSQLError(w, err)
	} else {
		// 检查权限
		if !policies.CanModifyArticle(*article) {
			art.ResponseForUnauthorized(w, r)
		} else {
			// 4. 读取成功，显示编辑文章表单
			view.Render(w, view.D{
				"Article": article,
				"Errors":  view.D{},
			}, "articles.edit", "articles._form_field")
		}
	}
}

// Update 更新文章
func (art *Article) Update(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)

	err := article.Get(id)

	// 3. 如果出现错误
	if err != nil {
		art.ResponseForSQLError(w, err)
	} else {
		// 4. 未出现错误

		// 检查权限
		if !policies.CanModifyArticle(*article) {
			art.ResponseForUnauthorized(w, r)
		} else {

			// 4.1 表单验证
			article.Title = r.PostFormValue("title")
			article.Body = r.PostFormValue("body")

			errors := requests.ValidateArticleForm(*article)

			if len(errors) == 0 {

				// 4.2 表单验证通过，更新数据
				rowsAffected, err := article.Update()

				if err != nil {
					// 数据库错误
					w.WriteHeader(http.StatusInternalServerError)
					fmt.Fprint(w, "500 服务器内部错误")
					return
				}

				// √ 更新成功，跳转到文章详情页
				if rowsAffected > 0 {
					showURL := route.Name2URL("articles.show", "id", id)
					http.Redirect(w, r, showURL, http.StatusFound)
				} else {
					fmt.Fprint(w, "您没有做任何更改！")
				}
			} else {

				// 4.3 表单验证不通过，显示理由
				view.Render(w, view.D{
					"Article": article,
					"Errors":  errors,
				}, "articles.edit", "articles._form_field")
			}
		}
	}
}

// Delete 删除文章
func (art *Article) Delete(w http.ResponseWriter, r *http.Request) {
	// 1. 获取 URL 参数
	id := route.GetRouteVariable("id", r)

	// 2. 读取对应的文章数据
	err := article.Get(id)

	// 3. 如果出现错误
	if err != nil {
		art.ResponseForSQLError(w, err)
	} else {
		// 4. 未出现错误，执行删除操作

		// 检查权限
		if !policies.CanModifyArticle(*article) {
			flash.Warning("您没有权限执行此操作！")
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			// 4. 未出现错误，执行删除操作
			rowsAffected, err := article.Delete()

			// 4.1 发生错误
			if err != nil {
				// 应该是 SQL 报错了
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "500 服务器内部错误")
			} else {
				// 4.2 未发生错误
				if rowsAffected > 0 {
					// 重定向到文章列表页
					indexURL := route.Name2URL("articles.index")
					http.Redirect(w, r, indexURL, http.StatusFound)
				} else {
					// Edge case
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprint(w, "404 文章未找到")
				}
			}
		}
	}
}

func validateArticleFormData(title string, body string) map[string]string {
	errors := make(map[string]string)
	// 验证标题
	if title == "" {
		errors["title"] = "标题不能为空"
	} else if utf8.RuneCountInString(title) < 3 || utf8.RuneCountInString(title) > 40 {
		errors["title"] = "标题长度需介于 3-40"
	}

	// 验证内容
	if body == "" {
		errors["body"] = "内容不能为空"
	} else if utf8.RuneCountInString(body) < 10 {
		errors["body"] = "内容长度需大于或等于 10 个字节"
	}

	return errors
}
