package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Picture struct {
	Id    		int64 `orm:"auto"`
	Img_url  	string `orm:"size(128)"`
	Note            string `orm:"size(128)"`
	Created_at      time.Time `orm:"auto_now_add;type(datetime)"`
	Updated_at      time.Time `orm:"auto_now;type(datetime)"`
}

func (p *Picture) TableName() string {
	return "pictures"
}

func init(){
	orm.RegisterModel(new(Picture))
}

func GetPicturesList() (p []Picture){
	o := orm.NewOrm()

	var pictures []Picture
	qs := o.QueryTable(new(Picture))
	qs.All(&pictures)
	for _, v := range pictures {
		p = append(p, v)
	}
	return p
}

func InsertPicture(p *Picture) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(p)
	return id, err
}

func UpdatePicture(id int64, params map[string]string) {
	o := orm.NewOrm()
	picture := Picture{Id: id}
	if o.Read(&picture) == nil {
		for k, v := range params {
			if k == "img_url" {
				picture.Img_url = v
			}
			if k == "note" {
				picture.Note = v
			}
		}
		o.Update(&picture)
	}
	return
}

func DeletePicture(id int64) (err error){
	o := orm.NewOrm()
	picture := Picture{Id: id}
	if _, err := o.Delete(&picture); err != nil {
		return err
	}
	return
}