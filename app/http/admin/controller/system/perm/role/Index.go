package role

import (
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/http/admin/model"
	"go_admin/app/res"
	"strconv"
)

type Search struct {
	status int8
	name   string
}

func Index(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize := app.Config.PageSize
	offset := (page - 1) * pageSize

	var count int64
	db := app.Db().Model(&model.AdminRole{})

	if status := c.DefaultQuery("status", ""); status != "" {
		db.Where("status=?", status)
	}

	if title := c.DefaultQuery("title", ""); title != "" {
		db.Where("title like ?", title+"%")
	}

	if level_id := c.DefaultQuery("level_id", ""); level_id != "" {
		db.Where("level_id = ?", level_id)
	}

	err := db.Count(&count).Error
	if err != nil {
		res.Json(c, res.Code(11), res.Msg(err.Error()))
		return
	}
	var list []model.AdminRole
	if count > 0 {
		err = db.Preload("Level").Offset(offset).Limit(pageSize).Find(&list).Error
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
