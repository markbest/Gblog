package utils

import "github.com/astaxie/beego"

func init() {
	//Init mysql connection
	InitSql()

	//Init template function
	beego.AddFuncMap("substring", SubString)
	beego.AddFuncMap("is_active", IsActive)
	beego.AddFuncMap("version", GetStaticVersion)
	beego.AddFuncMap("base_url", GetBaseUrl)
}
