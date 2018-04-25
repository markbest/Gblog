package controllers

type WorksController struct {
	BaseController
}

// @router /works [get]
func (c *WorksController) Get() {
	c.Layout = "frontend/layout/single.tpl"
	c.TplName = "frontend/works.tpl"
}
