package article

import (
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/http/blog/model"
	"go_admin/app/res"
)

func Info(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	if id == "" {
		res.Json(c, res.Code(1), res.Msg("id为空"))
		return
	}

	var blogArticle model.BlogArticle
	err := app.Db("blog").Model(&model.BlogArticle{}).Preload("Category").Where("id=?", id).Take(&blogArticle).Error
	if err != nil {
		res.Json(c, res.Code(11), res.Msg("数据错误"))
		return
	}

	res.Json(c, res.Data(gin.H{
		"info": blogArticle,
	}))
	return
}
