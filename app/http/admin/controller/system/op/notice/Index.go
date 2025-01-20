package notice

import (
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/http/admin/model"
	"go_admin/app/res"
	"strconv"
)

func Index(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize := app.Config.PageSize
	offset := (page - 1) * pageSize

	var count int64
	db := app.Db().Model(&model.AdminNotice{})

	if title := c.DefaultQuery("title", ""); title != "" {
		db.Where("title like ?", title+"%")
	}

	err := db.Count(&count).Error
	if err != nil {
		res.Json(c, res.Code(11), res.Msg(err.Error()))
		return
	}
	var list []model.AdminNotice
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
