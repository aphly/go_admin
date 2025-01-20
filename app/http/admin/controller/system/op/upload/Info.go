package upload

import (
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/http/admin/model"
	"go_admin/app/http/common/service/gorm"
	"go_admin/app/res"
	"os"
)

func Info(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	var adminUpload model.AdminUpload
	app.Db().Scopes(gorm.HaveDataPerm(c)).Where("id = ?", id).Take(&adminUpload)
	if adminUpload.Id == 0 {
		res.Json(c, res.Code(11), res.Msg("无权限"))
		return
	}
	content, err := os.ReadFile("." + adminUpload.Path)
	if err != nil {
		res.Json(c, res.Code(12), res.Msg(err.Error()))
		return
	}
	res.Json(c, res.Data(gin.H{
		"info": map[string]any{
			"id":        adminUpload.Id,
			"file_type": adminUpload.FileType,
			//"content":   string(content),
			"content": content,
		},
	}))
	return
}
