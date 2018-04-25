package controllers

import (
	"github.com/markbest/Gblog/models"
)

type SearchController struct {
	BaseController
}

// @router /search [get]
func (c *SearchController) Get() {
	var pageSize = 6
	page, _ := c.GetInt("page", 1)
	articles, num := models.GetSearchArticles(c.GetString("s"), pageSize, (page-1)*pageSize)
	pages := models.NewPage(page, pageSize, int(num), "/search?s="+c.GetString("s"))

	//模板变量
	c.Data["s"] = c.GetString("s")
	c.Data["article"] = articles
	c.Data["page"] = pages.Show()
	c.Layout = "frontend/layout/2columns-right.tpl"
	c.TplName = "frontend/article/list.tpl"
}
