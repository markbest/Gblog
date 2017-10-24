package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/markbest/Gblog/routers"
	_ "github.com/markbest/Gblog/utils"
)

func main() {
	beego.Run()
}
