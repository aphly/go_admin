package manager

import (
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/http/admin/model"
	"go_admin/app/res"
)

func Info(c *gin.Context) {
	uid := c.DefaultQuery("uid", "")
	if uid == "" {
		res.Json(c, res.Code(1), res.Msg("uid 为空"))
		return
	}
	var info model.AdminManager
	err := app.Db().Where("uid=?", uid).Take(&info).Error
	if err != nil {
		res.Json(c, res.Code(11), res.Msg(err.Error()))
		return
	}
	res.Json(c, res.Data(gin.H{
		"info": info,
	}))
	return
}
