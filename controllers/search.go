package controllers

import (
	"blog/models"
)

type SearchController struct {
	BaseController
}

// @router /search [get]
func (this *SearchController) Get() {
	//文章列表
	var pageSize int = 6
	page, err := this.GetInt("page")//获取页数
	if err != nil && page < 1 {
		page = 1
	}
	articles, num := models.GetSearchArticles(this.GetString("s"), pageSize, (page - 1) * pageSize)

	//分页
	var pages models.Page = models.NewPage(page, pageSize, int(num), "/search?s=" + this.GetString("s"))

	//侧边栏
	latest, _ := models.GetLatestArticles(8, 0)
	hot := models.GetTopViewArticles()
	tags := models.GetArticleTags()

	//模板变量
	this.Data["s"] = this.GetString("s")
	this.Data["article"] = articles
	this.Data["latest"] = latest
	this.Data["hot"] = hot
	this.Data["tags"] = tags
	this.Data["page"] = pages.Show()
	this.Layout = "layout/frontend/2columns_right.tpl"
	this.TplName = "article_list.tpl"
}