package controllers

import (
	"fmt"
	"goblog/app/models"
	"goblog/app/requests"
	"goblog/pkg/flash"
	"goblog/pkg/route"
	"goblog/pkg/view"
	"net/http"
)

// Category 文章分类控制器
type Category struct {
	Controllers
}

// Create 文章分类创建页面
func (*Category) Create(w http.ResponseWriter, r *http.Request) {
	view.Render(w, view.D{}, "categories.create")
}

// Store 保存文章分类
func (*Category) Store(w http.ResponseWriter, r *http.Request) {
	// 1. 初始化数据
	category := &models.Category{
		Name: r.PostFormValue("name"),
	}

	// 2. 表单验证
	errors := requests.ValidateCategoryForm(*category)

	// 3. 检测错误
	if len(errors) == 0 {
		// 创建文章分类
		category.Create()
		if category.ID > 0 {
			flash.Success("分类创建成功")
			indexURL := route.Name2URL("home")
			http.Redirect(w, r, indexURL, http.StatusFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "创建文章分类失败，请联系管理员")
		}
	} else {
		view.Render(w, view.D{
			"Category": category,
			"Errors":   errors,
		}, "categories.create")
	}
}

// Show 显示分类下的文章列表
func (*Category) Show(w http.ResponseWriter, r *http.Request) {
	//
}
