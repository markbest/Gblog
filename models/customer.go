package models

import (
	"blog/utils"
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
)

type Customer struct {
	Id         int64     `orm:"auto" form:"-"`
	Name       string    `orm:"size(64)" form:"name" valid:"Required;"`
	Email      string    `orm:"size(64);unique" form:"email" valid:"Required;Email"`
	Password   string    `orm:"size(32)" form:"password" valid:"Required"`
	Repassword string    `orm:"-" form:"repassword" valid:"Required"`
	Icon       string    `orm:"size(128)" form:"-"`
	Created_at time.Time `orm:"auto_now_add;type(datetime)" form:"-"`
	Updated_at time.Time `orm:"auto_now;type(datetime)" form:"-"`
}

func (c *Customer) TableName() string {
	return "customers"
}

func init() {
	orm.RegisterModel(new(Customer))
}

func InsertCustomer(c *Customer) (id int64, err error) {
	o := orm.NewOrm()
	c.Password = utils.GetMd5String(c.Password)
	id, err = o.Insert(c)
	return id, err
}

func UpdateCustomer(id int64, params map[string]string) error {
	o := orm.NewOrm()
	customer := Customer{Id: id}
	if o.Read(&customer) == nil {
		for k, v := range params {
			if k == "name" {
				customer.Name = v
			}
			if k == "old_password" {
				if customer.Password != utils.GetMd5String(v) {
					return errors.New("旧密码不符合")
				}
			}
			if k == "icon" {
				customer.Icon = v
			}
			if k == "new_password" {
				customer.Password = utils.GetMd5String(v)
			}
		}
		_, err := o.Update(&customer)
		return err
	}
	return nil
}

func GetCustomerInfo(id interface{}) (c Customer) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Customer))

	qs.Filter("id", id).One(&c)
	return c
}

func CustomerLogin(email string, password string) (Customer, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Customer))

	var c Customer
	qs.Filter("email", email).Filter("password", utils.GetMd5String(password)).One(&c)
	if c.Name == "" {
		return c, errors.New("用户名或密码错误！")
	}
	return c, nil
}
