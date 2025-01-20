package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/helper"
	"go_admin/app/http/common/model"
	"go_admin/app/http/common/service/upload"
	"go_admin/app/res"
	"path"
	"time"
)

func Avatar(c *gin.Context) {
	//uid, _ := c.Get("uid")
	//uidD := uid.(core.Uint)
	uid := c.DefaultQuery("uid", "")
	if uid == "" {
		res.Json(c, res.Code(1), res.Msg("uid为空"))
		return
	}
	var userModel model.User
	app.Db("common").Model(&model.User{}).Where("uid=?", uid).Take(&userModel)
	if userModel.Uid == 0 {
		res.Json(c, res.Code(2), res.Msg("数据错误"))
		return
	}

	file, _ := c.FormFile("file")
	uploadFile := upload.File{SizeConf: 1, TypeConf: []any{"image/jpeg", "image/png"}}
	var now = time.Now()
	var timeDir = "/public/upload/avatar/" + fmt.Sprintf("%d%d/%d/%d%d/", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute())
	var savePath = timeDir + helper.RandStr(16) + path.Ext(file.Filename)
	var uploadPath = "." + savePath

	err := uploadFile.LocalSave(c, file, uploadPath)
	if err != nil {
		res.Json(c, res.Code(11), res.Msg(err.Error()))
		return
	}
	app.Db("common").Model(&model.User{}).Where("uid=?", uid).Updates(&model.User{
		Avatar: savePath,
		Remote: 0,
	})
	upload.Del(userModel.Avatar, userModel.Remote)
	res.Json(c, res.Data(gin.H{
		"url": upload.Url(savePath, 0),
	}))
	return
}
