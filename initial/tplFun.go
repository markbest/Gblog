package initial

import (
	"github.com/astaxie/beego"
	"blog/utils"
)

func InitTplFunc() {
	beego.AddFuncMap("substring", utils.SubString)
	beego.AddFuncMap("is_active", utils.IsActive)
}
