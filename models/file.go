package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"strconv"
)

type File struct {
	Id    		int64 `orm:"auto"`
	Cat		*Category  `orm:"rel(fk)"`
	Title  		string `orm:"size(64)"`
	Name 		string `orm:"size(64)"`
	Size            int64
	Link            string `orm:"size(128)"`
	Type            string `orm:"size(64)"`
	Created_at      time.Time `orm:"auto_now_add;type(datetime)"`
}

func (f *File) TableName() string{
	return "files"
}

func init(){
	orm.RegisterModel(new(File))
}

func GetFilesList(page int, offset int) (f []File, count int64){
	o := orm.NewOrm()

	var files []File
	qs := o.QueryTable(new(File))
	count, _ = qs.Count()
	qs.OrderBy("-created_at").RelatedSel().Limit(page, offset).All(&files)
	for _, v := range files {
		f = append(f, v)
	}
	return f, count
}

func GetFileInfo(id int64) (f File){
	o := orm.NewOrm()
	qs := o.QueryTable(new(File))
	qs.Filter("id", id).One(&f)
	return f
}

func InsertFile(f *File) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(f)
	return id, err
}

func UpdateFile(id int64, params map[string]string) error {
	o := orm.NewOrm()
	file := File{Id: id}
	if o.Read(&file) == nil {
		for k, v := range params {
			if k == "title" {
				file.Title = v
			}
			if k == "cat_id" {
				id, _ := strconv.ParseInt(v, 10, 64)
				category := GetCategoryInfo(id)
				file.Cat = &category
			}
		}
		_, err := o.Update(&file)
		return err
	}
	return nil
}

func DeleteFile(id int64) error {
	o := orm.NewOrm()
	file := File{Id: id}
	if _, err := o.Delete(&file); err != nil {
		return err
	}
	return nil
}