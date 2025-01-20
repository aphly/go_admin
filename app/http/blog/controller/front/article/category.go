package article

import (
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/http/blog/model"
	"go_admin/app/res"
)

func Category(c *gin.Context) {
	var list []model.BlogCategory
	err := app.Db("blog").Model(&model.BlogCategory{}).Where("status", 1).Find(&list).Error
	if err != nil {
		res.Json(c, res.Code(12), res.Msg("错误"))
		return
	}
	res.Json(c, res.Data(gin.H{
		"list": list,
	}))
	return
}
