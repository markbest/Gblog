package main

import (
	_ "github.com/markbest/Gblog/initial"
	_ "github.com/markbest/Gblog/routers"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
)

func main() {
	beego.Run()
}
