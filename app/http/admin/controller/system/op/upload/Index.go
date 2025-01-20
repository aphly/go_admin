package upload

import (
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/http/admin/model"
	"go_admin/app/http/common/service/gorm"
	"go_admin/app/res"
	"strconv"
)

func Index(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize := app.Config.PageSize
	offset := (page - 1) * pageSize

	var count int64
	db := app.Db().Model(&model.AdminUpload{}).Scopes(gorm.HaveDataPerm(c))

	if uid := c.DefaultQuery("uid", ""); uid != "" {
		db.Where("uid=?", uid)
	}

	if level_id := c.DefaultQuery("level_id", ""); level_id != "" {
		db.Where("level_id=?", level_id)
	}

	err := db.Count(&count).Error
	if err != nil {
		res.Json(c, res.Code(11), res.Msg(err.Error()))
		return
	}
	var list []model.AdminUpload
	if count > 0 {
		err = db.Preload("Manager").Preload("Level").Order("id desc").Offset(offset).Limit(pageSize).Find(&list).Error
		if err != nil {
			res.Json(c, res.Code(12), res.Msg(err.Error()))
			return
		}
	}
	res.Json(c, res.Data(gin.H{
		"list":  list,
		"count": count,
	}))
	return
}
