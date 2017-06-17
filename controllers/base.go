package controllers

import (
	"github.com/astaxie/beego"
	"blog/models"
	"blog/utils"
	"encoding/json"
	"time"
)

type BaseController struct {
	beego.Controller
	isLogin bool
}

func (this *BaseController) Prepare() {
	//redis cache client
	redis := utils.GetRedisClient()

	//前台登陆信息
	var loginCustomer models.Customer
	userLogin := this.GetSession("userLogin")
	if userLogin == nil {
		this.isLogin = false
	} else {
		this.isLogin = true
		loginCustomer = models.GetCustomerInfo(this.GetSession("userId"))
	}

	//分类列表
	var allCategory []models.Category
	if redis.IsExist("allCategory") {
		cache_content := string(redis.Get("allCategory").([]uint8))
		json.Unmarshal([]byte(cache_content), &allCategory)
	} else {
		allCategory = models.GetCategoryList()
		if str, err := json.Marshal(allCategory); err == nil {
			redis.Put("allCategory", string(str), 24 * time.Hour)
		}
	}

	//侧边栏
	var latest, hot []models.Article
	if redis.IsExist("latest") {
		cache_content := string(redis.Get("latest").([]uint8))
		json.Unmarshal([]byte(cache_content), &latest)
	} else {
		latest, _ = models.GetLatestArticles(8, 0)
		if str, err := json.Marshal(latest); err == nil {
			redis.Put("latest", string(str), 24 * time.Hour)
		}
	}

	if redis.IsExist("hot") {
		cache_content := string(redis.Get("hot").([]uint8))
		json.Unmarshal([]byte(cache_content), &hot)
	} else {
		hot = models.GetTopViewArticles()
		if str, err := json.Marshal(hot); err == nil {
			redis.Put("hot", string(str), 24 * time.Hour)
		}
	}

	var tags []map[string]int64
	if redis.IsExist("tags") {
		cache_content := string(redis.Get("tags").([]uint8))
		json.Unmarshal([]byte(cache_content), &tags)
	} else {
		tags = models.GetArticleTags()
		if str, err := json.Marshal(tags); err == nil {
			redis.Put("tags", string(str), 24 * time.Hour)
		}
	}

	//模板变量
	this.Data["xsrf_token"] = this.XSRFToken()
	this.Data["current_url"] = this.Ctx.Request.RequestURI
	this.Data["isLogin"] = this.isLogin
	this.Data["loginCustomer"] = loginCustomer
	this.Data["category"] = allCategory
	this.Data["latest"] = latest
	this.Data["hot"] = hot
	this.Data["tags"] = tags
}

type AdminBaseController struct {
	beego.Controller
	isAdminLogin bool
}

func (this *AdminBaseController) Prepare() {
	//后台登录信息
	var loginUser models.User
	admin_userLogin := this.GetSession("admin_userLogin")
	if admin_userLogin == nil {
		this.isAdminLogin = false
	} else {
		this.isAdminLogin = true
		loginUser = models.GetUserInfo(this.GetSession("admin_userId"))
	}

	//模板变量
	this.Data["xsrf_token"] = this.XSRFToken()
	this.Data["current_url"] = this.Ctx.Request.RequestURI
	this.Data["isAdminLogin"] = this.isAdminLogin
	this.Data["loginUser"] = loginUser
}

