package controllers

import (
	"fmt"
	"github.com/markbest/Gblog/models"
	"github.com/markbest/Gblog/utils"
	"os"
	"path"
	"time"
)

type AdminPictureController struct {
	AdminBaseController
}

// @router /admin/picture [get]
func (c *AdminPictureController) ListPictures() {
	c.Layout = "admin/layout/2columns-left.tpl"
	c.TplName = "admin/picture/list.tpl"
}

// @router /admin/picture/edit [get]
func (c *AdminPictureController) EditPicture() {
	c.Data["pictures"] = models.GetPicturesList()
	c.Layout = "admin/layout/2columns-left.tpl"
	c.TplName = "admin/picture/edit.tpl"
}

// @router /admin/picture/:id [post]
func (c *AdminPictureController) UpdatePicture() {
	params := make(map[string]string)
	id, _ := c.GetInt64(":id")

	if c.GetString("is_delete") == "1" {
		models.DeletePicture(id)
	} else {
		if c.GetString("note") != "" {
			params["note"] = c.GetString("note")
		}
		models.UpdatePicture(id, params)
	}
	c.Redirect("/admin/picture/edit", 302)
}

// @router /admin/picture/upload [post]
func (c *AdminPictureController) UploadPicture() {
	f, h, err := c.GetFile("file")
	if err == nil {
		f.Close()
	} else {
		c.Data["json"] = map[string]interface{}{"success": 0, "message": err.Error()}
		c.ServeJSON()
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
	err = c.SaveToFile("file", fmt.Sprintf("%s/%s", dirPath, saveToFile))
	if err != nil {
		// 出错则输出错误信息
		c.Data["json"] = map[string]interface{}{"success": 0, "message": err.Error()}
		c.ServeJSON()
		return
	}

	// 上传图片
	picture := &models.Picture{}
	picture.ImgUrl = dirDatePrefix + "/" + saveToFile
	picture.Note = ""
	models.InsertPicture(picture)

	c.Data["json"] = map[string]interface{}{"success": 1, "message": dirDatePrefix + "/" + saveToFile}
	c.ServeJSON()
}

// @router /admin/markdown/upload [post]
func (c *AdminPictureController) UploadMarkdownPicture() {
	f, h, err := c.GetFile("editormd-image-file")
	if err == nil {
		f.Close()
	} else {
		c.Data["json"] = map[string]interface{}{"success": 0, "message": err.Error()}
		c.ServeJSON()
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
	err = c.SaveToFile("editormd-image-file", fmt.Sprintf("%s/%s", dirPath, saveToFile))
	if err != nil {
		c.Data["json"] = map[string]interface{}{"success": 0, "message": err.Error()}
		c.ServeJSON()
		return
	}

	// 上传图片
	picture := &models.Picture{}
	picture.ImgUrl = dirDatePrefix + "/" + saveToFile
	picture.Note = ""
	models.InsertPicture(picture)

	c.Data["json"] = map[string]interface{}{"success": 1, "message": "success upoload", "url": "/static/uploads/" + dirDatePrefix + "/" + saveToFile}
	c.ServeJSON()
}
