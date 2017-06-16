package controllers

import (
	"blog/models"
)

type ArticleController struct {
	BaseController
}

// @router /article/:id [get]
func (this *ArticleController) GetInfo() {
	//文章详情
	article_id, _ := this.GetInt64(":id")
	article := models.GetArticleInfo(article_id)

	//增加article的views
	models.IncreaseViews(article_id)

	//侧边栏
	latest, _ := models.GetLatestArticles(8, 0)
	hot := models.GetTopViewArticles()
	tags := models.GetArticleTags()

	//模板变量
	this.Data["latest"] = latest
	this.Data["hot"] = hot
	this.Data["tags"] = tags
	this.Data["article"] = article
	this.Layout = "layout/frontend/2columns_right.tpl"
	this.TplName = "article_info.tpl"
}

type AdminArticleController struct {
	AdminBaseController
}

// @router /admin [get]
func (this *AdminArticleController) Index() {
	this.Redirect("/admin/article", 302)
}

// @router /admin/article [get]
func (this *AdminArticleController) ListArticles() {
	//文章列表
	var pageSize int = 30
	page, err := this.GetInt("page")//获取页数
	if err != nil && page < 1 {
		page = 1
	}
	articles, num := models.GetLatestArticles(pageSize, (page - 1) * pageSize)

	//分页
	var pages models.Page = models.NewPage(page, pageSize, int(num), "/admin/article")

	//模板变量
	this.Data["articles"] = articles
	this.Data["page"] = pages.Show()
	this.Layout = "layout/admin/2columns_left.tpl"
	this.TplName = "admin/article/list.tpl"
}

// @router /admin/article/:id [get,post,delete]
func (this *AdminArticleController) UpdateArticle() {
	id, _ := this.GetInt64(":id")
	if this.Ctx.Input.Method() == "GET" {
		this.Data["category"] = models.GetCategoryList()
		this.Data["article"] = models.GetArticleInfo(id)
		this.Layout = "layout/admin/2columns_left.tpl"
		this.TplName = "admin/article/edit.tpl"
	} else {
		if this.GetString("_method") == "DELETE" {
			models.DeleteArticle(id)
		} else {
			params := make(map[string]string)
			params["title"] = this.GetString("title")
			params["slug"] = this.GetString("slug")
			params["summary"] = this.GetString("summary")
			params["body"] = this.GetString("body")
			params["cat_id"] = this.GetString("cat_id")
			params["user_id"] = "1"
			models.UpdateArticle(id, params)
		}
		this.Redirect("/admin/article", 302)
	}

}

// @router /admin/article/create [get,post]
func (this *AdminArticleController) AddArticle() {
	if this.Ctx.Input.Method() == "GET" {
		this.Data["category"] = models.GetCategoryList()
		this.Layout = "layout/admin/2columns_left.tpl"
		this.TplName = "admin/article/add.tpl"
	} else {
		article := &models.Article{}
		if err := this.ParseForm(article); err != nil {
			return
		}

		//文章归属分类
		cat_id, _ := this.GetInt64("cat_id")
		category := models.GetCategoryInfo(cat_id)
		article.Cat = &category

		//文章创建者
		user := models.GetUserInfo(1)
		article.User = &user

		models.InsertArticle(article)
		this.Redirect("/admin/article", 302)
	}
}
