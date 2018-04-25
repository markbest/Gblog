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

func (p *Page) getPageCount() {
	var tp = float32(p.TotalCount) / float32(p.PageSize)
	if tp < 1 {
		p.TotalPage = 1
	}
	var tpint = float32(int(tp))
	if tp > tpint {
		tpint += 1
	}
	p.TotalPage = int(tpint)
	p.LastPage = int(tpint)
	p.FirstPage = 1
	p.execUrl()
}

func (p *Page) execUrl() {
	if strings.Contains(p.Url, "?") {
		p.Url = strings.Join([]string{p.Url, "&page="}, "")
	} else {
		p.Url = strings.Join([]string{p.Url, "?page="}, "")
	}
}

func (p *Page) getUrl(page int) string {
	return strings.Join([]string{p.Url, strconv.Itoa(page)}, "")
}

func (p *Page) Show() string {
	var buf bytes.Buffer
	p.getPageCount()
	buf.WriteString("<ul class=\"pagination\">")
	if p.PageNo > 1 {
		buf.WriteString("<li><a href=\"")
		buf.WriteString(p.getUrl(1))
		buf.WriteString("\">上一页</a></li>")
	}
	for i := 1; i <= p.TotalPage; i++ {
		if i == p.PageNo {
			buf.WriteString("<li class=\"active\"><a href=\"javascript:void(0);\">")
			buf.WriteString(strconv.Itoa(i))
		} else {
			buf.WriteString("<li><a href=\"")
			buf.WriteString(p.getUrl(i))
			buf.WriteString("\">")
			buf.WriteString(strconv.Itoa(i))
		}
		buf.WriteString("</a></li>")
	}

	if p.PageNo < p.TotalPage {
		buf.WriteString("<li><a href=\"")
		var nextPage = p.PageNo + 1
		buf.WriteString(p.getUrl(nextPage))
		buf.WriteString("\">下一页</a></li>")
	}
	buf.WriteString("<li><a href=\"javascript:void(0);\">")
	buf.WriteString(strconv.Itoa(p.PageNo))
	buf.WriteString("/")
	buf.WriteString(strconv.Itoa(p.TotalPage))
	buf.WriteString("</a></li></ul>")
	return buf.String()
}
