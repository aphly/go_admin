package article

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_admin/app/helper"
	"go_admin/app/http/common/service/upload"
	"go_admin/app/res"
	"path"
	"time"
)

func File(c *gin.Context) {
	//uid, _ := c.Get("uid")
	//uidD := uid.(core.Uint)
	file, _ := c.FormFile("file")
	uploadFile := upload.File{SizeConf: 1, TypeConf: []any{"image/jpeg", "image/png"}}
	var now = time.Now()
	var timeDir = "/public/upload/temp/article/" + fmt.Sprintf("%d%d/%d/%d%d/", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute())
	var savePath = timeDir + helper.RandStr(16) + path.Ext(file.Filename)
	var uploadPath = "." + savePath

	err := uploadFile.LocalSave(c, file, uploadPath)
	if err != nil {
		res.Json(c, res.Code(11), res.Msg(err.Error()))
		return
	}
	res.Json(c, res.Data(gin.H{
		"url": upload.Url(savePath, 0),
	}))
	return
}
