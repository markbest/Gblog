package controllers

import (
	"blog/models"
	"blog/utils"
	"github.com/astaxie/beego"
)

type AdminConfigController struct {
	AdminBaseController
}

// @router /admin/config [get,post]
func (this *AdminConfigController) ConfigList() {
	if this.Ctx.Input.Method() == "GET" {
		//配置列表
		configs := models.GetListConfig()

		//模板变量
		this.Data["configs"] = configs
		this.Layout = "layout/admin/2columns_left.tpl"
		this.TplName = "admin/config.tpl"
	} else {
		config := &models.Config{}
		if err := this.ParseForm(config); err != nil {
			return
		}

		beego.Info(config)
		if _, err := models.InsertConfig(config); err == nil {
			//删除配置的缓存
			redis := utils.GetRedisClient()
			redis.Delete("configs")
		}
		this.Redirect("/admin/config", 302)
	}
}

// @router /admin/config/:id [post,put,delete]
func (this *AdminConfigController) ConfigUpdate() {
	if this.GetString("_method") == "DELETE" {
		id, _ := this.GetInt64(":id")
		err := models.DeleteConfig(id)
		if err == nil {
			//删除配置的缓存
			redis := utils.GetRedisClient()
			redis.Delete("configs")

			this.Redirect("/admin/config", 302)
		}
	} else {
		id, _ := this.GetInt64(":id")
		params := make(map[string]string)
		params["name"] = this.GetString("name")
		params["path"] = this.GetString("path")
		params["value"] = this.GetString("value")
		err := models.UpdateConfig(id, params)
		if err == nil {
			//删除配置的缓存
			redis := utils.GetRedisClient()
			redis.Delete("configs")

			this.Redirect("/admin/config", 302)
		}
	}
}
