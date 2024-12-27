package middleware

import (
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/http/model"
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
		resI := checkPermission(roleId, url, c)
		if resI != 0 {
			res.Json(c, res.Code(resI), res.Msg("无权限"))
			c.Abort()
			return
		}
		c.Set("level_ids", HaveLevelIds(roleId))
		c.Next()
	}
}

func checkPermission(roleId, url string, c *gin.Context) int {
	uid, _ := c.Get("uid")
	var adminManagerRole model.AdminManagerRole
	result := app.Db().Where("manager_uid = ? and role_id=?", uid, roleId).Take(&adminManagerRole)
	if result.RowsAffected == 0 {
		return 3
	}
	if roleId == "1" {
		return 0
	}
	var adminRoleMenu []model.AdminRoleMenu
	res1 := app.Db().InnerJoins("Menu", app.Db().Where(&model.AdminMenu{Status: 1, Path: url})).
		Where("role_id=?", adminManagerRole.RoleId).Take(&adminRoleMenu)
	if res1.RowsAffected > 0 {
		return 0
	}
	var adminRoleApi []model.AdminRoleApi
	res2 := app.Db().InnerJoins("Api", app.Db().Where(&model.AdminApi{Status: 1, Path: url})).
		Where("role_id=?", adminManagerRole.RoleId).Take(&adminRoleApi)
	if res2.RowsAffected > 0 {
		return 0
	}
	return 2
}

func HaveLevelIds(role_id string) []uint {
	var level_ids []uint
	var roleInfo model.AdminRole
	result := app.Db().Where("id=?", role_id).Take(&roleInfo)
	if result.RowsAffected == 0 {
		return level_ids
	}
	if roleInfo.DataPerm == 3 {
		app.Db().Model(&model.AdminLevelPath{}).Where("path_id=?", roleInfo.LevelId).Pluck("level_id", &level_ids)
	} else if roleInfo.DataPerm == 2 {
		level_ids = append(level_ids, roleInfo.LevelId)
	}
	return level_ids
}
