package controllers

import (
	"github.com/markbest/Gblog/models"
	"github.com/markbest/Gblog/utils"
)

type CategoryController struct {
	BaseController
}

// @router /category/:title [get]
func (this *CategoryController) ListArticle() {
	//文章列表
	pageSize := utils.StringToInt(this.config["web_perpage"])
	page, err := this.GetInt("page") //获取页数
	if err != nil && page < 1 {
		page = 1
	}
	category := models.GetCategoryInfoByTitle(this.GetString(":title"))
	articles, num := models.GetCategoryArticles(category.Id, pageSize, (page-1)*pageSize)

	//分页
	var pages models.Page = models.NewPage(page, pageSize, int(num), "/category/"+this.GetString(":title"))

	//模板变量
	this.Data["article"] = articles
	this.Data["page"] = pages.Show()
	this.Layout = "frontend/layout/2columns-right.tpl"
	this.TplName = "frontend/article/list.tpl"
}

type AdminCategoryController struct {
	AdminBaseController
}

// @router /admin/category [get,post]
func (this *AdminCategoryController) ListCategory() {
	if this.Ctx.Input.Method() == "GET" {
		//分类列表
		allCategory := models.GetCategoryList()

		//模板变量
		this.Data["category"] = allCategory
		this.Layout = "admin/layout/2columns-left.tpl"
		this.TplName = "admin/category.tpl"
	} else {
		category := &models.Category{}
		if err := this.ParseForm(category); err != nil {
			return
		}
		models.InsertCategory(category)

		//删除分类的缓存
		redis := utils.GetRedisClient()
		redis.Delete("allCategory")

		this.Redirect("/admin/category", 302)
	}
}

// @router /admin/category/:id [post,put,delete]
func (this *AdminCategoryController) UpdateCategory() {
	if this.GetString("_method") == "DELETE" {
		id, _ := this.GetInt64(":id")
		err := models.DeleteCategory(id)

		//删除分类的缓存
		redis := utils.GetRedisClient()
		redis.Delete("allCategory")

		if err == nil {
			this.Redirect("/admin/category", 302)
		}
	} else {
		id, _ := this.GetInt64(":id")
		params := make(map[string]string)
		params["title"] = this.GetString("title")
		params["parent_id"] = this.GetString("parent_id")
		params["sort"] = this.GetString("sort")
		err := models.UpdateCategory(id, params)

		//删除分类的缓存
		redis := utils.GetRedisClient()
		redis.Delete("allCategory")

		if err == nil {
			this.Redirect("/admin/category", 302)
		}
	}
}
