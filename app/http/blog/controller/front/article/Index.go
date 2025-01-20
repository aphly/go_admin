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
	db := app.Db("blog").Model(&model.BlogArticle{}).Where("status=?", 1)

	if title := c.DefaultQuery("title", ""); title != "" {
		db.Where("title like ?", title+"%")
	}
	var category model.BlogCategory
	if category_id := c.DefaultQuery("category_id", ""); category_id != "" {
		db.Where("category_id = ?", category_id)
		app.Db("blog").Where("status=? and id = ?", 1, category_id).Take(&category)
	}

	err := db.Count(&count).Error
	if err != nil {
		res.Json(c, res.Code(11), res.Msg(err.Error()))
		return
	}
	var list []model.BlogArticle
	if count > 0 {
		err = db.Order("sort desc,id desc").
			Select("title", "id", "updated_at", "status").
			Offset(offset).Limit(pageSize).Find(&list).Error
		if err != nil {
			res.Json(c, res.Code(12), res.Msg(err.Error()))
			return
		}
	}

	res.Json(c, res.Data(gin.H{
		"list":     list,
		"count":    count,
		"category": category,
	}))
	return
}
