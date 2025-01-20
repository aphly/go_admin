package article

import (
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/http/blog/model"
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
	db := app.Db("blog").Model(&model.BlogArticle{})

	if status := c.DefaultQuery("status", ""); status != "" {
		db.Where("status=?", status)
	}

	if title := c.DefaultQuery("title", ""); title != "" {
		db.Where("title like ?", title+"%")
	}

	if uid := c.DefaultQuery("uid", ""); uid != "" {
		db.Where("uid=?", uid)
	}

	err := db.Count(&count).Error
	if err != nil {
		res.Json(c, res.Code(11), res.Msg(err.Error()))
		return
	}
	var list []model.BlogArticle
	if count > 0 {
		err = db.Preload("Category").Order("id desc").Offset(offset).Limit(pageSize).Find(&list).Error
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
