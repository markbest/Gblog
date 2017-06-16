package controllers

import (
	"github.com/astaxie/beego"
	"blog/models"
)

type AdminUserController struct {
	AdminBaseController
}

// @router /admin/login [get,post]
func (this *AdminUserController) Login() {
	if this.Ctx.Input.Method() == "GET" {
		beego.ReadFromRequest(&this.Controller)
		this.Layout = "layout/admin/single.tpl"
		this.TplName = "admin/login.tpl"
	} else {
		flash := beego.NewFlash()
		email := this.GetString("email")
		password := this.GetString("password")
		user, err := models.AdminAuth(email, password)
		if err == nil {
			this.SetSession("admin_userLogin", 1)
			this.SetSession("admin_userId", user.Id)
			this.Redirect("/admin/article", 302)
		} else {
			flash.Error("登陆失败，请重试!")
			flash.Store(&this.Controller)
			this.Redirect("/admin/login", 302)
		}
	}
}

// @router /admin/logout [get]
func (this *AdminUserController) Logout() {
	this.DelSession("admin_userLogin")
	this.DelSession("admin_userInfo")
	this.Redirect("/", 302)
}

// @router /admin/user [get]
func (this *AdminUserController) ListUsers() {
	//用户列表
	var pageSize int = 30
	page, err := this.GetInt("page")//获取页数
	if err != nil && page < 1 {
		page = 1
	}
	users, num := models.GetAllUserList(pageSize, (page - 1) * pageSize)

	//分页
	var pages models.Page = models.NewPage(page, pageSize, int(num), "/admin/user")

	//模板变量
	this.Data["users"] = users
	this.Data["page"] = pages.Show()
	this.Layout = "layout/admin/2columns_left.tpl"
	this.TplName = "admin/user/list.tpl"
}

// @router /admin/user/:id [get,post]
func (this *AdminUserController) UpdateUser() {
	id, _ := this.GetInt64(":id")
	if this.Ctx.Input.Method() == "GET" {
		this.Data["user"] = models.GetUserInfo(id)
		this.Layout = "layout/admin/2columns_left.tpl"
		this.TplName = "admin/user/edit.tpl"
	} else {
		params := make(map[string]string)
		if this.GetString("name") != "" {
			params["name"] = this.GetString("name")
		}
		if this.GetString("password") != "" {
			params["password"] = this.GetString("password")
		}
		models.UpdateUser(id, params)
		this.Redirect("/admin/user", 302)
	}
}
