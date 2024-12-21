package manager

import (
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/core"
	"go_admin/app/http/model"
	"go_admin/app/http/service"
	"go_admin/app/res"
	"strconv"
)

func Blacklist(c *gin.Context) {
	uid := c.DefaultQuery("uid", "")
	if uid == "" {
		res.Json(c, res.Code(1), res.Msg("uid 为空"))
		return
	}
	uidP, _ := strconv.ParseInt(uid, 10, 64)
	err := app.Db().Model(&model.AdminManager{}).Where("uid=?", uidP).Update("status", 0).Error
	if err != nil {
		res.Json(c, res.Code(11), res.Msg("保存错误"))
		return
	}
	err1 := service.AddTokenToBlacklist(c, "jwt:managerBlacklist", core.Int64(uidP))
	if err1 != nil {
		res.Json(c, res.Code(12), res.Msg(err1.Error()))
		return
	}
	res.Json(c, res.Msg("冻结成功"))
	return
}
