package manager

import (
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/http/model"
	"go_admin/app/res"
	"strconv"
)

type Search struct {
	status   int8
	username string
	uid      int64
	phone    string
}

func Index(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize := app.Config.PageSize
	offset := (page - 1) * pageSize

	var count int64
	db := app.Db().Model(&model.AdminManager{})

	if status := c.DefaultQuery("status", ""); status != "" {
		db.Where("status=?", status)
	}

	if username := c.DefaultQuery("username", ""); username != "" {
		db.Where("username like ?", username+"%")
	}

	if uid := c.DefaultQuery("uid", ""); uid != "" {
		db.Where("uid=?", uid)
	}

	if phone := c.DefaultQuery("phone", ""); phone != "" {
		db.Where("phone like ?", phone+"%")
	}

	err := db.Count(&count).Error
	if err != nil {
		res.Json(c, res.Code(1), res.Msg(err.Error()))
		return
	}
	var list []model.AdminManager
	if count > 0 {
		err = db.Offset(offset).Limit(pageSize).Find(&list).Error
		if err != nil {
			res.Json(c, res.Code(2), res.Msg(err.Error()))
			return
		}
	}

	res.Json(c, res.Data(gin.H{
		"list":  list,
		"count": count,
	}))
	return
}
