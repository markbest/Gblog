package models

import (
	"time"
	"github.com/astaxie/beego/orm"
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

func UpdateConfig(id int64, params map[string]string) error{
	o := orm.NewOrm()

	config := Config{Id: id}
	if o.Read(&config) == nil {
		for k, v := range params {
			if k == "name" {
				config.Name = v
			}
			if k == "path" {
				config.Path = v
			}
			if k == "value" {
				config.Value = v
			}
		}
		_, err := o.Update(&config)
		return err
	}
	return nil
}

func DeleteConfig(id int64) error {
	o := orm.NewOrm()
	config := Config{Id: id}
	if _, err := o.Delete(&config); err != nil {
		return err
	}
	return nil
}

func GetConfigs() (map[string]string) {
	o := orm.NewOrm()

	var configs []Config
	l := make(map[string]string)
	qs := o.QueryTable(new(Config))
	qs.All(&configs)
	for _, v := range configs {
		l[v.Path] = v.Value
	}
	return l
}
