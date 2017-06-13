package controllers

import (
	"blog/models"
	"fmt"
	"os"
	"path"
	"github.com/astaxie/beego"
)

type FileController struct {
	BaseController
}

// @router /category/files [get]
func (this *FileController) FileList() {
	//文件列表
	var pageSize int = 6
	page, err := this.GetInt("page")//获取页数
	if err != nil && page < 1 {
		page = 1
	}
	articles, num := models.GetFilesList(pageSize, (page - 1) * pageSize)

	//分页
	var pages models.Page = models.NewPage(page, pageSize, int(num), "/files")

	//侧边栏
	latest, _ := models.GetLatestArticles(8, 0)
	hot := models.GetTopViewArticles()
	tags := models.GetArticleTags()

	//模板变量
	this.Data["files"] = articles
	this.Data["latest"] = latest
	this.Data["hot"] = hot
	this.Data["tags"] = tags
	this.Data["page"] = pages.Show()
	this.Layout = "layout/frontend/2columns_right.tpl"
	this.TplName = "files.tpl"
}

// @router /files/download/:id [get]
func (this *FileController) FileDownload() {
	//文件详情
	id, _ := this.GetInt64(":id")
	file := models.GetFileInfo(id)
	this.Ctx.Output.Download("static/uploads/" + file.Link)
}

type AdminFileController struct {
	AdminBaseController
}

// @router /admin/file [get]
func (this *AdminFileController) AdminFileList() {
	//文件列表
	var pageSize int = 30
	page, err := this.GetInt("page")//获取页数
	if err != nil && page < 1 {
		page = 1
	}
	files, num := models.GetFilesList(pageSize, (page - 1) * pageSize)

	//分类列表
	allCategory := models.GetCategoryList()

	//分页
	var pages models.Page = models.NewPage(page, pageSize, int(num), "/admin/file")

	//模板变量
	this.Data["xsrf_token"] = this.XSRFToken()
	this.Data["files"] = files
	this.Data["page"] = pages.Show()
	this.Data["category"] = allCategory
	this.Layout = "layout/admin/2columns_left.tpl"
	this.TplName = "admin/files.tpl"
}

// 获取文件信息的接口
type Stat interface {
	Stat() (os.FileInfo, error)
}

// @router /admin/file/upload [post]
func (this *AdminFileController) AdminFileUpload() {
	// 获取上传文件
	f, h, err := this.GetFile("file")
	if err == nil {
		// 关闭文件
		f.Close()
	} else {
		// 获取错误则输出错误信息
		this.Data["json"] = map[string]interface{}{"success": 0, "message": err.Error()}
		this.ServeJSON()
		return
	}
	// 设置保存目录
	dirPath := "./static/uploads/file/"
	// 设置保存文件名
	FileName := h.Filename
	// 将文件保存到服务器中
	err = this.SaveToFile("file", fmt.Sprintf("%s/%s", dirPath, FileName))
	if err != nil {
		// 出错则输出错误信息
		this.Data["json"] = map[string]interface{}{"success": 0, "message": err.Error()}
		this.ServeJSON()
		return
	}

	// 上传资料文件
	default_file_category := models.GetCategoryInfo(11)
	file := models.File{}
	file.Title = FileName
	file.Name = FileName
	file.Cat = &default_file_category
	if statInterface, ok := f.(Stat); ok {
		fileInfo, _ := statInterface.Stat()
		file.Size = fileInfo.Size()
		file.Type = path.Ext(FileName)
	}
	file.Link = FileName
	id, err := models.InsertFile(&file)
	if err != nil {
		beego.Info(id)
		beego.Info(err)
	}

	this.Data["json"] = map[string]interface{}{"success": 1, "message": "avatar/" + FileName}
	this.ServeJSON()
}

// @router /admin/file/:id [post,put,delete]
func (this *AdminFileController) UpdateFile() {
	if this.GetString("_method") == "DELETE" {
		id, _ := this.GetInt64(":id")
		err := models.DeleteFile(id)
		if err == nil {
			this.Redirect("/admin/file", 302)
		}
	} else {
		id, _ := this.GetInt64(":id")
		params := make(map[string]string)
		params["title"] = this.GetString("title")
		params["cat_id"] = this.GetString("cat_id")
		err := models.UpdateFile(id, params)
		if err == nil {
			this.Redirect("/admin/file", 302)
		}
	}
}
