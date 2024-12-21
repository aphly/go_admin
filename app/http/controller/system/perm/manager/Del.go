package manager

import (
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/http/model"
	"go_admin/app/res"
)

func Del(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	if len(ids) == 0 {
		res.Json(c, res.Code(1), res.Msg("ids为空"))
		return
	}
	err := app.Db().Where("uid in ?", ids).Delete(&model.AdminManager{}).Error
	if err != nil {
		res.Json(c, res.Code(11), res.Msg("删除失败"))
		return
	}
	res.Json(c, res.Msg("删除成功"))
	return
}
