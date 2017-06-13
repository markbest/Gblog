package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type Config struct {
	Id    		int64 `orm:"auto" form:"-"`
	Name  		string `orm:"size(255)" form:"name" valid:"Required;"`
	Path  		string `orm:"size(255)" form:"path" valid:"Required;"`
	Value		string `orm:"size(255)" form:"value" valid:"Required;"`
	Created_at      time.Time `orm:"auto_now_add;type(datetime)" form:"-"`
	Updated_at      time.Time `orm:"auto_now;type(datetime)" form:"-"`
}

func (c *Config) TableName() string {
	return "configs"
}

func init(){
	orm.RegisterModel(new(Config))
}


func GetListConfig() (c []Config) {
	o := orm.NewOrm()

	var configs []Config
	qs := o.QueryTable(new(Config))
	qs.All(&configs)
	for _, v := range configs {
		c = append(c, v)
	}
	return c
}

func InsertConfig(c *Config) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(c)
	return id, err
}

func MultiUpdateConfig(params map[string]string) {
	o := orm.NewOrm()

	for k, v := range params {
		id, _ := strconv.ParseInt(k, 10, 64)
		config := Config{Id: id}
		if o.Read(&config) == nil {
			config.Value = v
			o.Update(&config)
		}
	}
	return
}
