package level

import (
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/http/admin/model"
	"go_admin/app/res"
)

func FRebuild(pid uint) {
	var levelData []model.AdminLevel
	app.Db().Where("pid=?", pid).Find(&levelData)
	for _, v := range levelData {
		app.Db().Where("level_id=?", v.Id).Delete(&model.AdminLevelPath{})
		level := 0
		var levelPathData []model.AdminLevelPath
		app.Db().Where("level_id=?", v.Pid).Order("level asc").Find(&levelPathData)
		var data []*model.AdminLevelPath
		for _, v1 := range levelPathData {
			data = append(data, &model.AdminLevelPath{
				LevelId: v.Id,
				PathId:  v1.PathId,
				Level:   level,
			})
			level++
		}
		data = append(data, &model.AdminLevelPath{
			LevelId: v.Id,
			PathId:  v.Id,
			Level:   level,
		})
		app.Db().Create(data)
		FRebuild(v.Id)
	}
}

func Rebuild(c *gin.Context) {
	app.Db().Where("1 = 1").Delete(&model.AdminLevelPath{})
	FRebuild(0)
	res.Json(c, res.Msg("重建成功"))
	return
}
