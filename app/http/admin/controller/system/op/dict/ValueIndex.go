package dict

import (
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/http/admin/model"
	"go_admin/app/res"
	"strconv"
)

func ValueIndex(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize := app.Config.PageSize
	offset := (page - 1) * pageSize

	dict_id := c.DefaultQuery("dict_id", "")
	if dict_id == "" {
		res.Json(c, res.Code(1), res.Msg("dict_id 为空"))
		return
	}

	var count int64
	db := app.Db().Model(&model.AdminDictValue{}).Where("dict_id=?", dict_id)

	err := db.Count(&count).Error
	if err != nil {
		res.Json(c, res.Code(11), res.Msg(err.Error()))
		return
	}
	var list []model.AdminDictValue
	if count > 0 {
		err = db.Order("id desc").Offset(offset).Limit(pageSize).Find(&list).Error
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
