package main

import (
	_ "blog/initial"
	_ "blog/routers"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
)

func main() {
	beego.Run()
}
