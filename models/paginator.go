package models

import (
	"bytes"
	"strconv"
	"strings"
)

type Page struct {
	PageNo     int    //当前页
	PageSize   int    //每页多少数据
	TotalPage  int    //总共多少页
	TotalCount int    //总共多少条数据
	FirstPage  int    //第一页
	LastPage   int    //最后一页
	Url        string //链接
}

func NewPage(PageNo int, PageSize int, TotalCount int, Url string) Page {
	return Page{PageNo: PageNo, PageSize: PageSize, TotalCount: TotalCount, Url: Url}
}

//计算总页数
func (this *Page) getPageCount() {
	var tp float32 = float32(this.TotalCount) / float32(this.PageSize)
	if tp < 1 {
		this.TotalPage = 1
	}
	var tpint float32 = float32(int(tp))

	if tp > tpint {
		tpint += 1
	}
	this.TotalPage = int(tpint)
	this.LastPage = int(tpint)
	this.FirstPage = 1
	this.execUrl()
}

//格式化URL地址
func (this *Page) execUrl() {
	if strings.Contains(this.Url, "?") {
		this.Url = strings.Join([]string{this.Url, "&page="}, "")
	} else {
		this.Url = strings.Join([]string{this.Url, "?page="}, "")
	}
}

//获取URL组织
func (this *Page) getUrl(page int) string {
	return strings.Join([]string{this.Url, strconv.Itoa(page)}, "")
}

//
func (this *Page) Show() string {
	this.getPageCount()
	var buf bytes.Buffer
	buf.WriteString("<ul class=\"pagination\">")
	if this.PageNo > 1 {
		buf.WriteString("<li><a href=\"")
		buf.WriteString(this.getUrl(1))
		buf.WriteString("\">上一页</a></li>")
	}
	for i := 1; i <= this.TotalPage; i++ {
		if i == this.PageNo {
			buf.WriteString("<li class=\"active\"><a href=\"javascript:void(0);\">")
			buf.WriteString(strconv.Itoa(i))
		} else {
			buf.WriteString("<li><a href=\"")
			buf.WriteString(this.getUrl(i))
			buf.WriteString("\">")
			buf.WriteString(strconv.Itoa(i))
		}
		buf.WriteString("</a></li>")
	}

	if this.PageNo < this.TotalPage {
		buf.WriteString("<li><a href=\"")
		var nextPage int = this.PageNo + 1
		buf.WriteString(this.getUrl(nextPage))
		buf.WriteString("\">下一页</a></li>")
	}
	buf.WriteString("<li><a href=\"javascript:void(0);\">")
	buf.WriteString(strconv.Itoa(this.PageNo))
	buf.WriteString("/")
	buf.WriteString(strconv.Itoa(this.TotalPage))
	buf.WriteString("</a></li></ul>")
	return buf.String()
}
