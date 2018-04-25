package controllers

import (
	"github.com/astaxie/beego"
	"github.com/markbest/Gblog/api"
	"github.com/markbest/Gblog/models"
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
	var data []api.Article
	perPage, _ := c.GetInt("per_page", 10)
	page, _ := c.GetInt("page", 1)
	Articles, total := models.GetLatestArticles(perPage, perPage*(page-1))
	for _, v := range Articles {
		var art api.Article
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
