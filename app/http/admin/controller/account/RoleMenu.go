package account

import (
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/http/admin/model"
	"go_admin/app/res"
)

func RoleMenu(c *gin.Context) {
	uid, _ := c.Get("uid")
	role_id := c.Request.Header.Get("x-roleid")
	if role_id == "" {
		res.Json(c, res.Code(1), res.Msg("role_id 为空"))
		return
	}
	var adminManagerRole model.AdminManagerRole
	result := app.Db().Where("manager_uid = ? and role_id=?", uid, role_id).Take(&adminManagerRole)
	if result.RowsAffected == 0 {
		res.Json(c, res.Code(2), res.Msg("无权限"))
		return
	}
	var adminRoleMenu []model.AdminRoleMenu
	app.Db().InnerJoins("Menu", app.Db().Where(&model.AdminMenu{Status: 1}).Where(map[string]any{"type": []int{1, 2}})).
		Where("role_id=?", adminManagerRole.RoleId).Find(&adminRoleMenu)
	res.Json(c, res.Data(gin.H{
		"manager_role_menu": adminRoleMenu,
	}))
	return
}
