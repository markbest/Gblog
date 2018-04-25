package controllers

import (
	"github.com/markbest/Gblog/models"
	"github.com/markbest/Gblog/utils"
)

type MainController struct {
	BaseController
}

// @router / [get]
func (c *MainController) Get() {
	pageSize := utils.StringToInt(c.config["web_perpage"])
	page, _ := c.GetInt("page", 1)
	articles, num := models.GetLatestArticles(pageSize, (page-1)*pageSize)
	pages := models.NewPage(page, pageSize, int(num), "/")

	c.Data["article"] = articles
	c.Data["page"] = pages.Show()
	c.Layout = "frontend/layout/2columns-right.tpl"
	c.TplName = "frontend/article/list.tpl"
}
