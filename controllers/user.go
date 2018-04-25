package controllers

import (
	"github.com/astaxie/beego"
	"github.com/markbest/Gblog/models"
)

type AdminUserController struct {
	AdminBaseController
}

// @router /admin/login [get,post]
func (c *AdminUserController) Login() {
	if c.Ctx.Input.Method() == "GET" {
		beego.ReadFromRequest(&c.Controller)
		c.Layout = "admin/layout/single.tpl"
		c.TplName = "admin/login.tpl"
	} else {
		flash := beego.NewFlash()
		email := c.GetString("email")
		password := c.GetString("password")
		user, err := models.AdminAuth(email, password)
		if err == nil {
			c.SetSession("admin_userLogin", 1)
			c.SetSession("admin_userId", user.Id)
			c.Redirect("/admin/article", 302)
		} else {
			flash.Error("登陆失败，请重试!")
			flash.Store(&c.Controller)
			c.Redirect("/admin/login", 302)
		}
	}
}

// @router /admin/logout [get]
func (c *AdminUserController) Logout() {
	c.DelSession("admin_userLogin")
	c.DelSession("admin_userInfo")
	c.Redirect("/", 302)
}

// @router /admin/user [get]
func (c *AdminUserController) ListUsers() {
	var pageSize = 30
	page, _ := c.GetInt("page", 1)
	users, num := models.GetAllUserList(pageSize, (page-1)*pageSize)
	pages := models.NewPage(page, pageSize, int(num), "/admin/user")

	c.Data["users"] = users
	c.Data["page"] = pages.Show()
	c.Layout = "admin/layout/2columns-left.tpl"
	c.TplName = "admin/user/list.tpl"
}

// @router /admin/user/:id [get,post]
func (c *AdminUserController) UpdateUser() {
	id, _ := c.GetInt64(":id")
	if c.Ctx.Input.Method() == "GET" {
		c.Data["user"] = models.GetUserInfo(id)
		c.Layout = "admin/layout/2columns-left.tpl"
		c.TplName = "admin/user/edit.tpl"
	} else {
		params := make(map[string]string)
		if c.GetString("name") != "" {
			params["name"] = c.GetString("name")
		}
		if c.GetString("password") != "" {
			params["password"] = c.GetString("password")
		}
		models.UpdateUser(id, params)
		c.Redirect("/admin/user", 302)
	}
}
