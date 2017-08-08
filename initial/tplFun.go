package initial

import (
	"blog/utils"
	"github.com/astaxie/beego"
)

func InitTplFunc() {
	beego.AddFuncMap("substring", utils.SubString)
	beego.AddFuncMap("is_active", utils.IsActive)
	beego.AddFuncMap("version", utils.GetStaticVersion)
	beego.AddFuncMap("fsize", utils.FileSizeUnitConversion)
	beego.AddFuncMap("base_url", utils.GetBaseUrl)
}
