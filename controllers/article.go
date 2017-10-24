package controllers

import (
	"encoding/json"
	"github.com/markbest/Gblog/models"
	"github.com/markbest/Gblog/utils"
	"strconv"
	"time"
)

type ArticleController struct {
	BaseController
}

// @router /article/:id [get]
func (this *ArticleController) GetInfo() {
	//redis cache client
	redis := utils.GetRedisClient()

	//文章详情（如果有缓存，先从缓存中取数据）
	var article models.Article
	article_id, _ := this.GetInt64(":id")
	cache_tag := "article-" + strconv.FormatInt(article_id, 10)
	if redis.IsExist(cache_tag) {
		cache_content := string(redis.Get(cache_tag).([]uint8))
		json.Unmarshal([]byte(cache_content), &article)
	} else {
		article = models.GetArticleInfo(article_id)
		if str, err := json.Marshal(article); err == nil {
			cache_time := utils.StringToInt64(this.config["web_cache_time"])
			redis.Put(cache_tag, string(str), time.Duration(cache_time)*time.Second)
		}
	}

	//增加article的views
	models.IncreaseViews(article_id)

	//模板变量
	this.Data["article"] = article
	this.Layout = "frontend/layout/2columns-right.tpl"
	this.TplName = "frontend/article/info.tpl"
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
	page, err := this.GetInt("page") //获取页数
	if err != nil && page < 1 {
		page = 1
	}
	articles, num := models.GetLatestArticles(pageSize, (page-1)*pageSize)

	//分页
	var pages models.Page = models.NewPage(page, pageSize, int(num), "/admin/article")

	//模板变量
	this.Data["articles"] = articles
	this.Data["page"] = pages.Show()
	this.Layout = "admin/layout/2columns-left.tpl"
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

			//删除文章的缓存
			redis := utils.GetRedisClient()
			cache_tag := "article-" + strconv.FormatInt(id, 10)
			redis.Delete(cache_tag)
		}
		this.Redirect("/admin/article", 302)
	}

}

// @router /admin/article/create [get,post]
func (this *AdminArticleController) AddArticle() {
	if this.Ctx.Input.Method() == "GET" {
		this.Data["category"] = models.GetCategoryList()
		this.Layout = "admin/layout/2columns-left.tpl"
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
