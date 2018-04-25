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
func (c *ArticleController) GetInfo() {
	//redis cache client
	redis := utils.GetRedisClient()

	//文章详情
	var article models.Article
	articleId, _ := c.GetInt64(":id")
	cacheTag := "article-" + strconv.FormatInt(articleId, 10)
	if redis.IsExist(cacheTag) {
		cacheContent := string(redis.Get(cacheTag).([]uint8))
		json.Unmarshal([]byte(cacheContent), &article)
	} else {
		article = models.GetArticleInfo(articleId)
		if str, err := json.Marshal(article); err == nil {
			cacheTime := utils.StringToInt64(c.config["web_cache_time"])
			redis.Put(cacheTag, string(str), time.Duration(cacheTime)*time.Second)
		}
	}

	//增加article的views
	models.IncreaseViews(articleId)

	//模板变量
	c.Data["article"] = article
	c.Layout = "frontend/layout/2columns-right.tpl"
	c.TplName = "frontend/article/info.tpl"
}

type AdminArticleController struct {
	AdminBaseController
}

// @router /admin [get]
func (c *AdminArticleController) Index() {
	c.Redirect("/admin/article", 302)
}

// @router /admin/article [get]
func (c *AdminArticleController) ListArticles() {
	var pageSize = 30
	page, _ := c.GetInt("page", 1)
	articles, num := models.GetLatestArticles(pageSize, (page-1)*pageSize)
	pages := models.NewPage(page, pageSize, int(num), "/admin/article")

	c.Data["articles"] = articles
	c.Data["page"] = pages.Show()
	c.Layout = "admin/layout/2columns-left.tpl"
	c.TplName = "admin/article/list.tpl"
}

// @router /admin/article/:id [get,post,delete]
func (c *AdminArticleController) UpdateArticle() {
	id, _ := c.GetInt64(":id")
	if c.Ctx.Input.Method() == "GET" {
		c.Data["category"] = models.GetCategoryList()
		c.Data["article"] = models.GetArticleInfo(id)
		c.Layout = "admin/layout/2columns-left.tpl"
		c.TplName = "admin/article/edit.tpl"
	} else {
		if c.GetString("_method") == "DELETE" {
			models.DeleteArticle(id)
		} else {
			params := make(map[string]string)
			params["title"] = c.GetString("title")
			params["slug"] = c.GetString("slug")
			params["summary"] = c.GetString("summary")
			params["body"] = c.GetString("body")
			params["cat_id"] = c.GetString("cat_id")
			params["user_id"] = "1"
			models.UpdateArticle(id, params)

			//删除文章的缓存
			redis := utils.GetRedisClient()
			cacheTag := "article-" + strconv.FormatInt(id, 10)
			redis.Delete(cacheTag)
		}
		c.Redirect("/admin/article", 302)
	}
}

// @router /admin/article/create [get,post]
func (c *AdminArticleController) AddArticle() {
	if c.Ctx.Input.Method() == "GET" {
		c.Data["category"] = models.GetCategoryList()
		c.Layout = "admin/layout/2columns-left.tpl"
		c.TplName = "admin/article/add.tpl"
	} else {
		article := &models.Article{}
		if err := c.ParseForm(article); err != nil {
			return
		}

		//文章归属分类
		catId, _ := c.GetInt64("cat_id")
		category := models.GetCategoryInfo(catId)
		article.Cat = &category

		//文章创建者
		user := models.GetUserInfo(1)
		article.User = &user

		models.InsertArticle(article)
		c.Redirect("/admin/article", 302)
	}
}
