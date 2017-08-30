package models

import (
	"github.com/markbest/Gblog/utils"
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id         int64     `orm:"auto" form:"-"`
	Name       string    `orm:"size(64)" form:"name" valid:"Required;"`
	Email      string    `orm:"size(64);unique" form:"email" valid:"Required;Email"`
	Password   string    `orm:"size(32)" form:"password" valid:"Required"`
	Created_at time.Time `orm:"auto_now_add;type(datetime)" form:"-"`
	Updated_at time.Time `orm:"auto_now;type(datetime)" form:"-"`
}

func (u *User) TableName() string {
	return "users"
}

func init() {
	orm.RegisterModel(new(User))
}

func AdminAuth(email string, password string) (User, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))

	var u User
	qs.Filter("email", email).Filter("password", utils.GetMd5String(password)).One(&u)
	if u.Name == "" {
		return u, errors.New("用户名或密码错误！")
	}
	return u, nil
}

func GetUserInfo(id interface{}) (u User) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))
	qs.Filter("id", id).One(&u)
	return u
}

func GetAllUserList(page int, offset int) (u []User, count int64) {
	o := orm.NewOrm()

	var users []User
	qs := o.QueryTable(new(User))
	count, _ = qs.Count()
	qs.OrderBy("-created_at").RelatedSel().Limit(page, offset).All(&users)
	for _, v := range users {
		u = append(u, v)
	}
	return u, count
}

func UpdateUser(id int64, params map[string]string) {
	o := orm.NewOrm()

	user := User{Id: id}
	if o.Read(&user) == nil {
		for k, v := range params {
			if k == "name" {
				user.Name = v
			}
			if k == "password" {
				user.Password = utils.GetMd5String(v)
			}
		}
		o.Update(&user)
	}
	return
}
