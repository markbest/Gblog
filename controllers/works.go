package controllers

type WorksController struct {
	BaseController
}

// @router /works [get]
func (this *WorksController) Get() {
	//模板变量
	this.Layout = "frontend/layout/single.tpl"
	this.TplName = "frontend/works.tpl"
}
