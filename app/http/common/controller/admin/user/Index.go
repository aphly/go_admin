package user

import (
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/http/common/model"
	"go_admin/app/http/common/service/upload"
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
	//uid, _ := c.Get("uid")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize := app.Config.PageSize
	offset := (page - 1) * pageSize

	var count int64
	db := app.Db("common").Model(&model.User{})

	if status := c.DefaultQuery("status", ""); status != "" {
		db.Where("status=?", status)
	}

	if uid := c.DefaultQuery("uid", ""); uid != "" {
		db.Where("uid=?", uid)
	}

	err := db.Count(&count).Error
	if err != nil {
		res.Json(c, res.Code(11), res.Msg(err.Error()))
		return
	}
	var list []model.User
	if count > 0 {
		err = db.Preload("UserAuth").Order("uid desc").Offset(offset).Limit(pageSize).Find(&list).Error
		if err != nil {
			res.Json(c, res.Code(12), res.Msg(err.Error()))
			return
		}
	}

	for k := range list {
		list[k].Avatar = upload.Url(list[k].Avatar, list[k].Remote)
	}

	res.Json(c, res.Data(gin.H{
		"list":  list,
		"count": count,
	}))
	return
}
