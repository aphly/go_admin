package dict

import (
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/http/model"
	"go_admin/app/res"
)

func Status(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	if id == "" {
		res.Json(c, res.Code(1), res.Msg("id 为空"))
		return
	}
	status := c.DefaultQuery("status", "0")
	if status == "1" {
		status = "1"
	} else {
		status = "0"
	}
	err := app.Db().Model(&model.AdminManager{}).Where("id=?", id).Update("status", status).Error
	if err != nil {
		res.Json(c, res.Code(11), res.Msg("操作失败"))
		return
	}
	res.Json(c, res.Msg("操作成功"))
	return
}
