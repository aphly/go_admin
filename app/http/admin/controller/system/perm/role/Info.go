package role

import (
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/http/admin/model"
	"go_admin/app/res"
)

func Info(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	if id == "" {
		res.Json(c, res.Code(1), res.Msg("id 为空"))
		return
	}
	var info model.AdminRole
	err := app.Db().Where("id=?", id).Take(&info).Error
	if err != nil {
		res.Json(c, res.Code(11), res.Msg(err.Error()))
		return
	}
	res.Json(c, res.Data(gin.H{
		"info": info,
	}))
	return
}
