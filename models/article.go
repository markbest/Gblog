package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"strings"
	"strconv"
)

type Article struct {
	Id    		int64 `orm:"auto" form:"-"`
	Title  		string `orm:"size(128)" form:"title" valid:"Required;"`
	Slug  		string `orm:"size(128)" form:"slug" valid:"Required;"`
	Summary		string `orm:"size(256)" form:"summary" valid:"Required;"`
	Body		string `orm:"content;type(text);null" form:"body" valid:"Required;"`
	Image           string `orm:"size(128)" form:"-"`
	Views		int64  `form:"-"`
	User		*User  `orm:"rel(fk)" form:"-"`
	Cat		*Category `orm:"rel(fk)" form:"-"`
	Created_at      time.Time `orm:"auto_now_add;type(datetime)" form:"-"`
	Updated_at      time.Time `orm:"auto_now;type(datetime)" form:"-"`
}

func (a *Article) TableName() string{
	return "articles"
}

func init(){
	orm.RegisterModel(new(Article))
}

func GetLatestArticles(page int, offset int) (a []Article, count int64){
	o := orm.NewOrm()

	var articles []Article
	qs := o.QueryTable(new(Article))
	count, _ = qs.Count()
	qs.OrderBy("-id").RelatedSel().Limit(page, offset).All(&articles, "Id", "Title", "Summary", "Slug", "Views", "Created_at")
	for _, v := range articles {
		a = append(a, v)
	}
	return a, count
}

func GetTopViewArticles() (a []Article){
	o := orm.NewOrm()

	var articles []Article
	qs := o.QueryTable(new(Article))
	qs.OrderBy("-views").RelatedSel().Limit(8).All(&articles, "Id", "Title", "Views")
	for _, v := range articles {
		a = append(a, v)
	}
	return a
}

func GetSearchArticles(keyword string, page int, offset int) (a []Article, count int64){
	o := orm.NewOrm()

	var articles []Article
	qs := o.QueryTable(new(Article)).Filter("title__icontains", keyword)
	count, _ = qs.Count()
	qs.OrderBy("-created_at").Limit(page, offset).All(&articles)
	for _, v := range articles {
		a = append(a, v)
	}
	return a, count
}

func GetArticleTags() (t []map[string]int64){
	o := orm.NewOrm()

	var articles []Article
	qs := o.QueryTable(new(Article))
	qs.OrderBy("-views").All(&articles)

	for _, v := range articles {
		tags := make(map[string]int64)
		tags_list := strings.Split(v.Slug, "、")
		for _, value := range tags_list {
			tags[value] = v.Id
		}
		t = append(t, tags)
	}
	return t
}

func GetArticleInfo(id int64) (a Article){
	o := orm.NewOrm()

	qs := o.QueryTable(new(Article))
	qs.Filter("id", id).RelatedSel().One(&a)
	return a
}

func InsertArticle(a *Article) (id int64, err error){
	o := orm.NewOrm()
	id, err = o.Insert(a)
	return id, err
}

func UpdateArticle(id int64, params map[string]string){
	o := orm.NewOrm()

	article := Article{Id: id}
	if o.Read(&article) == nil {
		for k, v := range params {
			if k == "title" {
				article.Title = v
			}
			if k == "slug" {
				article.Slug = v
			}
			if k == "summary" {
				article.Summary = v
			}
			if k == "body" {
				article.Body = v
			}
			if k == "cat_id" {
				id, _ := strconv.ParseInt(v, 10, 64)
				category := GetCategoryInfo(id)
				article.Cat = &category
			}
			if k == "user_id" {
				id, _ := strconv.ParseInt(v, 10, 64)
				user := GetUserInfo(id)
				article.User = &user
			}
		}
		o.Update(&article)
	}
	return
}

func DeleteArticle(id int64) error {
	o := orm.NewOrm()

	article := Article{Id: id}
	if _, err := o.Delete(&article); err != nil {
		return err
	}
	return nil
}

func IncreaseViews(id int64) {
	o := orm.NewOrm()

	article := Article{Id: id}
	if o.Read(&article) == nil {
		article.Views = int64(article.Views) + 1
		o.Update(&article, "Views")
	}
	return
}