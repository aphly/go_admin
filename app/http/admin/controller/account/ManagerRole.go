package account

import (
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/core"
	"go_admin/app/http/admin/model"
	"go_admin/app/res"
)

func ManagerRole(c *gin.Context) {
	uid, _ := c.Get("uid")
	uidD := uid.(core.Uint)
	var adminManagerRole []model.AdminManagerRole
	app.Db().InnerJoins("Role", app.Db().Where(&model.AdminRole{Status: 1})).
		Where("admin_manager_role.manager_uid = ?", uidD).Find(&adminManagerRole)
	res.Json(c, res.Data(gin.H{
		"manager_role": adminManagerRole,
	}))
	return
}
