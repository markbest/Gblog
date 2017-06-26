package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"log"
)

func GetRedisClient() cache.Cache {
	redis_host := beego.AppConfig.String("redis_host")
	redis_password := beego.AppConfig.String("redis_password")
	redis_port := beego.AppConfig.String("redis_port")
	redis_db := beego.AppConfig.String("redis_db")
	conn := redis_host + ":" + redis_port
	redis, err := cache.NewCache("redis", `{"key":"beecacheRedis", "conn":"`+conn+`", "dbNum":"`+redis_db+`","password":"`+redis_password+`"}`)
	if err != nil {
		log.Println(err)
	}
	return redis
}
