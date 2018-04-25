package controllers

import (
	"github.com/markbest/Gblog/models"
	"github.com/markbest/Gblog/utils"
)

type CategoryController struct {
	BaseController
}

// @router /category/:title [get]
func (c *CategoryController) ListArticle() {
	pageSize := utils.StringToInt(c.config["web_perpage"])
	page, _ := c.GetInt("page", 1)
	category := models.GetCategoryInfoByTitle(c.GetString(":title"))
	articles, num := models.GetCategoryArticles(category.Id, pageSize, (page-1)*pageSize)
	pages := models.NewPage(page, pageSize, int(num), "/category/"+c.GetString(":title"))

	c.Data["article"] = articles
	c.Data["page"] = pages.Show()
	c.Layout = "frontend/layout/2columns-right.tpl"
	c.TplName = "frontend/article/list.tpl"
}

type AdminCategoryController struct {
	AdminBaseController
}

// @router /admin/category [get,post]
func (c *AdminCategoryController) ListCategory() {
	if c.Ctx.Input.Method() == "GET" {
		allCategory := models.GetCategoryList()
		c.Data["category"] = allCategory
		c.Layout = "admin/layout/2columns-left.tpl"
		c.TplName = "admin/category.tpl"
	} else {
		category := &models.Category{}
		if err := c.ParseForm(category); err != nil {
			return
		}
		models.InsertCategory(category)
		redis := utils.GetRedisClient()
		redis.Delete("allCategory")
		c.Redirect("/admin/category", 302)
	}
}

// @router /admin/category/:id [post,put,delete]
func (c *AdminCategoryController) UpdateCategory() {
	if c.GetString("_method") == "DELETE" {
		id, _ := c.GetInt64(":id")
		err := models.DeleteCategory(id)

		redis := utils.GetRedisClient()
		redis.Delete("allCategory")
		if err == nil {
			c.Redirect("/admin/category", 302)
		}
	} else {
		id, _ := c.GetInt64(":id")
		params := make(map[string]string)
		params["title"] = c.GetString("title")
		params["parent_id"] = c.GetString("parent_id")
		params["sort"] = c.GetString("sort")

		err := models.UpdateCategory(id, params)
		if err == nil {
			redis := utils.GetRedisClient()
			redis.Delete("allCategory")
			c.Redirect("/admin/category", 302)
		}
	}
}
