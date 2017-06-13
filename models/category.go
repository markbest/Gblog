package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"strconv"
)

type Category struct {
	Id    		int64 `orm:"auto" form:"-"`
	Title  		string `orm:"size(256)" form:"title" valid:"Required;"`
	Parent_id  	int64  `form:"parent_id" valid:"Required;"`
	Sort       	int64  `form:"sort" valid:"Required;"`
	Created_at      time.Time `orm:"auto_now_add;type(datetime)" form:"-"`
	Updated_at      time.Time `orm:"auto_now;type(datetime)" form:"-"`
	Count_articles  int64 `orm:"-" form:"-"`
	Sub_category    []Category `orm:"-" form:"-"`
}

func (c *Category) TableName() string{
	return "categories"
}

func init(){
	orm.RegisterModel(new(Category))
}

func GetSubCategory(parent_id int64) (c []Category) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Category))

	var l []Category
	qs.Filter("parent_id", parent_id).OrderBy("sort").All(&l)
	for _, v := range l {
		_, count_articles := GetCategoryArticles(v.Id, 10, 1)
		allCategory := Category{v.Id, v.Title, v.Parent_id, v.Sort, v.Created_at, v.Updated_at, count_articles, nil}
		c = append(c, allCategory)
	}
	return c
}

func GetCategoryList() (c []Category){
	o := orm.NewOrm()
	qs := o.QueryTable(new(Category))

	var l []Category
	qs.Filter("parent_id", 0).OrderBy("sort").All(&l)
	for _, v := range l {
		var count int64
		subCategory := GetSubCategory(v.Id)
		if subCategory != nil {
			for _, sub := range subCategory {
				count += sub.Count_articles
			}
		} else {
			_, count = GetCategoryArticles(v.Id, 10, 1)
		}

		allCategory := Category{v.Id, v.Title, v.Parent_id, v.Sort, v.Created_at, v.Updated_at, count, GetSubCategory(v.Id)}
		c = append(c, allCategory)
	}
	return c
}

func GetCategoryArticles(id int64, page int, offset int) (a []Article, count int64) {
	o := orm.NewOrm()

	var articles []Article
	aqs := o.QueryTable(new(Article)).Filter("cat_id", id)
	count, _ = aqs.Count()
	aqs.OrderBy("-created_at").Limit(page, offset).All(&articles)
	for _, v := range articles {
		a = append(a, v)
	}
	return a, count
}

func InsertCategory(c *Category) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(c)
	return id, err
}

func DeleteCategory(id int64) error{
	o := orm.NewOrm()
	category := Category{Id: id}
	if _, err := o.Delete(&category); err != nil {
		return err
	}
	return nil
}

func UpdateCategory(id int64, params map[string]string) error{
	o := orm.NewOrm()
	category := Category{Id: id}
	if o.Read(&category) == nil {
		for k, v := range params {
			if k == "title" {
				category.Title = v
			}
			if k == "parent_id" {
				id, _ := strconv.ParseInt(v, 10, 64)
				category.Parent_id = id
			}
			if k == "sort" {
				id, _ := strconv.ParseInt(v, 10, 64)
				category.Sort = id
			}
		}
		_, err := o.Update(&category)
		return err
	}
	return nil
}

func GetCategoryInfo(id int64) (c Category){
	o := orm.NewOrm()

	qs := o.QueryTable(new(Category))
	qs.Filter("id", id).One(&c)
	return c
}

func GetCategoryInfoByTitle(title string) (c Category){
	o := orm.NewOrm()

	qs := o.QueryTable(new(Category))
	qs.Filter("title", title).One(&c)
	return c
}