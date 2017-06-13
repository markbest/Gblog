package main

import (
	_"blog/routers"
	"github.com/astaxie/beego"
	_"blog/initial"
	_ "github.com/astaxie/beego/session/redis"
)

func main() {
	beego.Run()
}

