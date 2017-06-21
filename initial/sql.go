package initial

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func InitSql() {
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	db := beego.AppConfig.String("dbdatabase")
	conn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + db + "?charset=utf8&loc=Asia%2FShanghai"

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", conn)
}
