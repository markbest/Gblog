package controllers

import (
	"github.com/astaxie/beego"
	"github.com/markbest/Gblog/models"
	"strconv"
)

type ApiController struct {
	beego.Controller
}

// @router /category/list [get]
func (c *ApiController) GetCategoryList() {
	data := models.GetApiCategoryJson()
	c.Data["json"] = data
	c.ServeJSON()
}

// @router /article/list [get]
func (c *ApiController) GetArticleList() {
	var data []models.ApiArticle
	perPage, _ := c.GetInt("per_page", 10)
	page, _ := c.GetInt("page", 1)
	Articles, total := models.GetLatestArticles(perPage, perPage*(page-1))
	for _, v := range Articles {
		var art models.ApiArticle
		art.Id = v.Id
		art.Title = v.Title
		art.Body = v.Body
		art.Slug = v.Slug
		art.Summary = v.Summary
		art.Views = v.Views
		art.User = v.User.Name
		art.CreatedAt = &v.CreatedAt
		art.UpdatedAt = &v.UpdatedAt
		data = append(data, art)
	}
	c.Data["json"] = map[string]interface{}{"Total": total, "Per_page": perPage, "Page": page, "Data": data}
	c.ServeJSON()
}

// @router /search [get]
func (c *ApiController) GetSearchResult() {
	var data []models.ApiSearch
	articles, _ := models.GetSearchArticles(c.GetString("s"), 0, 0)

	if len(articles) > 0 {
		for _, article := range articles {
			var search models.ApiSearch
			search.Id = article.Id
			search.Url = "http://www.markbest.site/article/" + strconv.Itoa(int(article.Id))
			search.Title = article.Title
			search.CreatedAt = article.CreatedAt
			search.UpdatedAt = article.UpdatedAt
			data = append(data, search)
		}
	}

	c.Data["json"] = data
	c.ServeJSON()
}
