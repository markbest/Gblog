package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/markbest/Gblog/models"
	"github.com/markbest/Gblog/utils"
	"time"
)

type BaseController struct {
	beego.Controller
	config map[string]string
}

func (c *BaseController) Prepare() {
	//redis cache client
	redis := utils.GetRedisClient()

	//配置信息
	var cacheTime int64
	config := make(map[string]string)
	if redis.IsExist("configs") {
		cacheContent := string(redis.Get("configs").([]uint8))
		json.Unmarshal([]byte(cacheContent), &config)
		cacheTime = utils.StringToInt64(config["web_cache_time"])
	} else {
		config = models.GetConfigs()
		cacheTime = utils.StringToInt64(config["web_cache_time"])
		if str, err := json.Marshal(config); err == nil {
			redis.Put("configs", string(str), time.Duration(cacheTime)*time.Second)
		}
	}
	c.config = config

	//分类列表
	var allCategory []models.Category
	if redis.IsExist("allCategory") {
		cacheContent := string(redis.Get("allCategory").([]uint8))
		json.Unmarshal([]byte(cacheContent), &allCategory)
	} else {
		allCategory = models.GetCategoryList()
		if str, err := json.Marshal(allCategory); err == nil {
			redis.Put("allCategory", string(str), time.Duration(cacheTime)*time.Second)
		}
	}

	//侧边栏
	var latest, hot []models.Article
	if redis.IsExist("latest") {
		cacheContent := string(redis.Get("latest").([]uint8))
		json.Unmarshal([]byte(cacheContent), &latest)
	} else {
		latest, _ = models.GetLatestArticles(8, 0)
		if str, err := json.Marshal(latest); err == nil {
			redis.Put("latest", string(str), time.Duration(cacheTime)*time.Second)
		}
	}

	if redis.IsExist("hot") {
		cacheContent := string(redis.Get("hot").([]uint8))
		json.Unmarshal([]byte(cacheContent), &hot)
	} else {
		hot = models.GetTopViewArticles()
		if str, err := json.Marshal(hot); err == nil {
			redis.Put("hot", string(str), time.Duration(cacheTime)*time.Second)
		}
	}

	var tags []map[string]int64
	if redis.IsExist("tags") {
		cacheContent := string(redis.Get("tags").([]uint8))
		json.Unmarshal([]byte(cacheContent), &tags)
	} else {
		tags = models.GetArticleTags()
		if str, err := json.Marshal(tags); err == nil {
			redis.Put("tags", string(str), time.Duration(cacheTime)*time.Second)
		}
	}

	//模板变量
	c.Data["xsrf_token"] = c.XSRFToken()
	c.Data["current_url"] = c.Ctx.Request.RequestURI
	c.Data["category"] = allCategory
	c.Data["latest"] = latest
	c.Data["hot"] = hot
	c.Data["tags"] = tags
	c.Data["configs"] = config
}

type AdminBaseController struct {
	beego.Controller
	isAdminLogin bool
}

func (c *AdminBaseController) Prepare() {
	//后台登录信息
	var loginUser models.User
	adminUserLogin := c.GetSession("admin_userLogin")
	if adminUserLogin == nil {
		c.isAdminLogin = false
	} else {
		c.isAdminLogin = true
		loginUser = models.GetUserInfo(c.GetSession("admin_userId"))
	}

	//模板变量
	c.Data["xsrf_token"] = c.XSRFToken()
	c.Data["current_url"] = c.Ctx.Request.RequestURI
	c.Data["isAdminLogin"] = c.isAdminLogin
	c.Data["loginUser"] = loginUser
}
