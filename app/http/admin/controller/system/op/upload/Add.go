package upload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/core"
	"go_admin/app/helper"
	"go_admin/app/http/admin/model"
	"go_admin/app/http/common/service/upload"
	"go_admin/app/res"
	"path"
	"time"
)

func Add(c *gin.Context) {
	uid, _ := c.Get("uid")
	uidD := uid.(core.Uint)
	levelId, _ := c.Get("level_id")
	levelIdd := levelId.(uint)
	file, _ := c.FormFile("file")
	uploadFile := upload.File{SizeConf: 1, TypeConf: []any{"image/jpeg", "image/png", "text/plain"}}
	var now = time.Now()
	var timeDir = "/private/upload/file/" + fmt.Sprintf("%d%d/%d/%d%d/", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute())
	var savePath = timeDir + helper.RandStr(16) + path.Ext(file.Filename)
	var uploadPath = "." + savePath
	err := uploadFile.LocalSave(c, file, uploadPath)
	if err != nil {
		res.Json(c, res.Code(1), res.Msg(err.Error()))
		return
	}
	adminUpload := model.AdminUpload{
		Uid:      uidD,
		LevelId:  levelIdd,
		Path:     savePath,
		FileType: file.Header["Content-Type"][0],
		FileSize: file.Size,
		Remote:   0,
	}
	err = app.Db().Create(&adminUpload).Error
	if err != nil {
		res.Json(c, res.Code(11), res.Msg("保存错误"))
		return
	}
	res.Json(c, res.Data(gin.H{
		"id":  adminUpload.Id,
		"url": upload.Url(savePath, 0),
	}))
	return
}
