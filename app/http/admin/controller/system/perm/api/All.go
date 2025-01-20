package api

import (
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/http/admin/model"
	"go_admin/app/res"
)

func All(c *gin.Context) {
	var list []model.AdminApi
	err := app.Db().Find(&list).Error
	if err != nil {
		res.Json(c, res.Code(2), res.Msg(err.Error()))
		return
	}
	res.Json(c, res.Data(gin.H{
		"list": list,
	}))
	return
}
