package controllers

import (
	"github.com/astaxie/beego"
	"github.com/markbest/Gblog/models"
	"github.com/markbest/Gblog/utils"
)

type AdminConfigController struct {
	AdminBaseController
}

// @router /admin/config [get,post]
func (c *AdminConfigController) ConfigList() {
	if c.Ctx.Input.Method() == "GET" {
		configs := models.GetListConfig()
		c.Data["configs"] = configs
		c.Layout = "admin/layout/2columns-left.tpl"
		c.TplName = "admin/config.tpl"
	} else {
		config := &models.Config{}
		if err := c.ParseForm(config); err != nil {
			return
		}

		beego.Info(config)
		if _, err := models.InsertConfig(config); err == nil {
			redis := utils.GetRedisClient()
			redis.Delete("configs")
		}
		c.Redirect("/admin/config", 302)
	}
}

// @router /admin/config/:id [post,put,delete]
func (c *AdminConfigController) ConfigUpdate() {
	if c.GetString("_method") == "DELETE" {
		id, _ := c.GetInt64(":id")
		err := models.DeleteConfig(id)
		if err == nil {
			redis := utils.GetRedisClient()
			redis.Delete("configs")
			c.Redirect("/admin/config", 302)
		}
	} else {
		id, _ := c.GetInt64(":id")
		params := make(map[string]string)
		params["name"] = c.GetString("name")
		params["path"] = c.GetString("path")
		params["value"] = c.GetString("value")
		err := models.UpdateConfig(id, params)
		if err == nil {
			redis := utils.GetRedisClient()
			redis.Delete("configs")
			c.Redirect("/admin/config", 302)
		}
	}
}
