package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/markbest/Gblog/controllers"
)

func init() {
	/* 验证后台客户登录 */
	var UserAuth = func(ctx *context.Context) {
		_, ok := ctx.Input.Session("admin_userLogin").(int)
		if !ok && ctx.Request.RequestURI != "/admin/login" && ctx.Request.RequestURI != "/admin/logout" {
			ctx.Redirect(302, "/admin/login")
		}
	}
	beego.InsertFilter("/admin/*", beego.BeforeRouter, UserAuth)

	beego.Include(&controllers.MainController{})
	beego.Include(&controllers.CategoryController{})
	beego.Include(&controllers.ArticleController{})
	beego.Include(&controllers.SearchController{})
	beego.Include(&controllers.WorksController{})

	//后台路由
	beego.Include(&controllers.AdminArticleController{})
	beego.Include(&controllers.AdminCategoryController{})
	beego.Include(&controllers.AdminUserController{})
	beego.Include(&controllers.AdminPictureController{})
	beego.Include(&controllers.AdminConfigController{})

	//API路由
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/v1",
			beego.NSInclude(
				&controllers.ApiController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
