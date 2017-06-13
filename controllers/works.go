package controllers

type WorksController struct {
	BaseController
}

// @router /works [get]
func (this *WorksController) Get() {
	//模板变量
	this.Layout = "layout/frontend/single.tpl"
	this.TplName = "works.tpl"
}
