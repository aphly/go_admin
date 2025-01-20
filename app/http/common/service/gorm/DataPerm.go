package gorm

import (
	"go_admin/app"
	"go_admin/app/http/admin/model"
)

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

func BelongLevelIds(role_id string) []uint {
	var path_ids []uint
	var roleInfo model.AdminRole
	result := app.Db().Where("id=?", role_id).Take(&roleInfo)
	if result.RowsAffected == 0 {
		return path_ids
	}
	app.Db().Model(&model.AdminLevelPath{}).Where("level_id=?", roleInfo.LevelId).Pluck("path_id", &path_ids)
	return path_ids
}

//func BelongLevelIdsByUid(uid any) []uint {
//	var path_ids []uint
//	var manager model.AdminManager
//	result := app.Db().Where("uid=?", uid).Take(&manager)
//	if result.RowsAffected == 0 {
//		return path_ids
//	}
//	app.Db().Model(&model.AdminLevelPath{}).Where("level_id=?", manager.LevelId).Pluck("path_id", &path_ids)
//	return path_ids
//}
