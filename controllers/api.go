package controllers

import (
	"github.com/astaxie/beego"
	"blog/models"
	"blog/api"
)

type ApiController struct {
	beego.Controller
}

// @router /category/list [get]
func (this *ApiController) GetCategoryList() {
	var data = []api.Category{}
	for _, v := range models.GetLayerCategoryList() {
		var cat api.Category
		cat.Id = v.Id
		cat.Title = v.Title
		cat.Parent_id = v.Parent_id
		cat.Sort = v.Sort
		cat.Created_at = v.Created_at
		cat.Updated_at = v.Updated_at
		data = append(data, cat)
	}
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

	Articles, total := models.GetLatestArticles(per_page, per_page * (page - 1))
	for _, v := range Articles {
		var cat api.Category
		cat.Id = v.Cat.Id
		cat.Title = v.Cat.Title
		cat.Parent_id = v.Cat.Parent_id
		cat.Sort = v.Cat.Sort
		cat.Created_at = v.Cat.Created_at
		cat.Updated_at = v.Cat.Updated_at

		var art api.Article
		art.Id = v.Id
		art.Title = v.Title
		art.Body = v.Body
		art.Slug = v.Slug
		art.Summary = v.Summary
		art.Image = v.Image
		art.Category = cat
		art.User_id = v.User.Id
		art.Created_at = v.Created_at
		art.Updated_at = v.Updated_at
		data = append(data, art)
	}
	this.Data["json"] = map[string]interface{}{"total": total, "per_page": per_page, "page": page, "data": data}
	this.ServeJSON()
}