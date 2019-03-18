package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"log"
)

func GetRedisClient() cache.Cache {
	redisHost := beego.AppConfig.String("redis_host")
	redisPwd := beego.AppConfig.String("redis_password")
	redisPort := beego.AppConfig.String("redis_port")
	redisDb := beego.AppConfig.String("redis_db")
	conn := redisHost + ":" + redisPort
	redis, err := cache.NewCache("redis", `{"key":"beecacheRedis", "conn":"`+conn+`", "dbNum":"`+redisDb+`","password":"`+redisPwd+`"}`)
	if err != nil {
		log.Println(err)
	}
	return redis
}
