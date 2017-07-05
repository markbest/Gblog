package controllers

import (
	"blog/api"
	"blog/models"
	"github.com/astaxie/beego"
)

type ApiController struct {
	beego.Controller
}

// @router /category/list [get]
func (this *ApiController) GetCategoryList() {
	var data = []api.Category{}
	data = models.GetApiCategoryJson()
	this.Data["json"] = data
	this.ServeJSON()
}

// @router /article/list [get]
func (this *ApiController) GetArticleList() {
	var data = []api.Article{}
	per_page, _ := this.GetInt("per_page")
	page, _ := this.GetInt("page")

	if per_page <= 0 {
		per_page = 10
	}

	if page <= 0 {
		page = 1
	}

	Articles, total := models.GetLatestArticles(per_page, per_page*(page-1))
	for _, v := range Articles {
		var art api.Article
		art.Id = v.Id
		art.Title = v.Title
		art.Body = v.Body
		art.Slug = v.Slug
		art.Summary = v.Summary
		art.Views = v.Views
		art.User = v.User.Name
		art.Created_at = &v.Created_at
		art.Updated_at = &v.Updated_at
		data = append(data, art)
	}
	this.Data["json"] = map[string]interface{}{"Total": total, "Per_page": per_page, "Page": page, "Data": data}
	this.ServeJSON()
}
