package utils

import (
	"github.com/astaxie/beego/cache"
)

func GetCacheClient() cache.Cache{
	cache, _ := cache.NewCache("file", `{"CachePath":"./cache","FileSuffix":".cache","DirectoryLevel":2,"EmbedExpiry":120}`)
	return cache
}
