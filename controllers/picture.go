package controllers

import (
	"blog/models"
	"blog/utils"
	"fmt"
	"time"
	"os"
	"path"
)

type AdminPictureController struct {
	AdminBaseController
}

// @router /admin/picture [get]
func (this *AdminPictureController) ListPictures() {
	this.Layout = "layout/admin/2columns_left.tpl"
	this.TplName = "admin/picture/list.tpl"
}

// @router /admin/picture/edit [get]
func (this *AdminPictureController) EditPicture() {
	this.Data["pictures"] = models.GetPicturesList()
	this.Layout = "layout/admin/2columns_left.tpl"
	this.TplName = "admin/picture/edit.tpl"
}

// @router /admin/picture/:id [post]
func (this *AdminPictureController) UpdatePicture() {
	params := make(map[string]string)
	id, _ := this.GetInt64(":id")

	if this.GetString("is_delete") == "1" {
		models.DeletePicture(id)
	} else {
		if this.GetString("note") != ""{
			params["note"] = this.GetString("note")
		}
		models.UpdatePicture(id, params)
	}
	this.Redirect("/admin/picture/edit", 302)
}

// @router /admin/picture/upload [post]
func (this *AdminPictureController) UploadPicture() {
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
	dirDatePrefix := "picture/" + time.Unix(time.Now().Unix(), 0).Format("2006/01/02")
	dirPath := "./static/uploads/" + dirDatePrefix
	os.MkdirAll(dirPath, 0777)

	// 设置保存文件名
	FileName := h.Filename
	saveToFile := string(utils.Krand(8, utils.KC_RAND_KIND_ALL)) + path.Ext(FileName)

	// 将文件保存到服务器中
	err = this.SaveToFile("file", fmt.Sprintf("%s/%s", dirPath, saveToFile))
	if err != nil {
		// 出错则输出错误信息
		this.Data["json"] = map[string]interface{}{"success": 0, "message": err.Error()}
		this.ServeJSON()
		return
	}

	// 上传图片
	picture := models.Picture{}
	picture.Img_url = dirDatePrefix + "/" + saveToFile
	picture.Note = ""
	models.InsertPicture(&picture)

	this.Data["json"] = map[string]interface{}{"success": 1, "message": dirDatePrefix + "/" + saveToFile}
	this.ServeJSON()
}

// @router /admin/markdown/upload [post]
func (this *AdminPictureController) UploadMarkdownPicture() {
	// 获取上传文件
	f, h, err := this.GetFile("editormd-image-file")
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
	dirDatePrefix := "picture/" + time.Unix(time.Now().Unix(), 0).Format("2006/01/02")
	dirPath := "./static/uploads/" + dirDatePrefix
	os.MkdirAll(dirPath, 0777)

	// 设置保存文件名
	FileName := h.Filename
	saveToFile := string(utils.Krand(8, utils.KC_RAND_KIND_ALL)) + path.Ext(FileName)

	// 将文件保存到服务器中
	err = this.SaveToFile("editormd-image-file", fmt.Sprintf("%s/%s", dirPath, saveToFile))
	if err != nil {
		// 出错则输出错误信息
		this.Data["json"] = map[string]interface{}{"success": 0, "message": err.Error()}
		this.ServeJSON()
		return
	}

	// 上传图片
	picture := models.Picture{}
	picture.Img_url = dirDatePrefix + "/" + saveToFile
	picture.Note = ""
	models.InsertPicture(&picture)

	this.Data["json"] = map[string]interface{}{"success": 1, "message": "success upoload", "url": "/static/uploads/" + dirDatePrefix + "/" + saveToFile}
	this.ServeJSON()
}