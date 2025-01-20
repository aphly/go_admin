package middleware

import (
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/http/admin/model"
	"go_admin/app/res"
)

func RbacHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Request.URL.Path
		roleId := c.Request.Header.Get("X-RoleId")
		if roleId == "" {
			res.Json(c, res.Code(1), res.Msg("无权限"))
			c.Abort()
			return
		}
		uid, _ := c.Get("uid")
		var adminManagerRole model.AdminManagerRole
		result := app.Db().InnerJoins("Role", app.Db().Where(&model.AdminRole{Status: 1})).
			Where("manager_uid = ? and role_id=?", uid, roleId).Take(&adminManagerRole)
		if result.RowsAffected == 0 {
			res.Json(c, res.Code(3), res.Msg("无权限"))
			c.Abort()
			return
		}
		resI := checkPermission(roleId, url, c)
		if resI != 0 {
			res.Json(c, res.Code(resI), res.Msg("无权限"))
			c.Abort()
			return
		}
		c.Set("role_id", roleId)
		c.Set("level_id", adminManagerRole.Role.LevelId)
		c.Next()
	}
}

func checkPermission(roleId, url string, c *gin.Context) int {
	if roleId == "1" {
		return 0
	}
	var adminRoleMenu []model.AdminRoleMenu
	res1 := app.Db().InnerJoins("Menu", app.Db().Where(&model.AdminMenu{Status: 1, Path: url})).
		Where("role_id=?", roleId).Take(&adminRoleMenu)
	if res1.RowsAffected > 0 {
		return 0
	}
	var adminRoleApi []model.AdminRoleApi
	res2 := app.Db().InnerJoins("Api", app.Db().Where(&model.AdminApi{Status: 1, Path: url})).
		Where("role_id=?", roleId).Take(&adminRoleApi)
	if res2.RowsAffected > 0 {
		return 0
	}
	return 2
}
