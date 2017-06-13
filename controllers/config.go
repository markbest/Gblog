package controllers

import (
	"blog/models"
	"github.com/astaxie/beego"
)

type AdminConfigController struct {
	AdminBaseController
}

// @router /admin/config [get,post]
func (this *AdminConfigController) Setting() {
	if this.Ctx.Input.Method() == "GET" {
		//配置列表
		configs := models.GetListConfig()

		//模板变量
		this.Data["xsrf_token"] = this.XSRFToken()
		this.Data["configs"] = configs
		this.Layout = "layout/admin/2columns_left.tpl"
		this.TplName = "admin/config.tpl"
	} else {
		flash := beego.NewFlash()
		config := &models.Config{}
		if err := this.ParseForm(config); err != nil {
			return
		}

		beego.Info(config)
		if _, err := models.InsertConfig(config); err != nil {
			beego.Info(err)
			flash.Error("新增配置失败，请重试！")
		}
		this.Redirect("/admin/config", 302)
	}
}

// @router /admin/config/multiupdate [post]
func (this *AdminConfigController) Multiupdate() {
	ids := this.GetStrings("id")
	values := this.GetStrings("value")
	configs := make(map[string]string)
	for k, v := range values {
		configs[ids[k]] = v
	}
	models.MultiUpdateConfig(configs)
	this.Redirect("/admin/config", 302)
}
