package article

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/helper/editor"
	"go_admin/app/http/blog/model"
	"go_admin/app/http/common/service/upload"
	"go_admin/app/res"
)

type del struct {
	Ids []uint `json:"ids" binding:"required"`
}

func Del(c *gin.Context) {
	form := del{}
	err := c.ShouldBind(&form)
	if err != nil {
		if jsonErr, ok := err.(*json.UnmarshalTypeError); ok {
			msg := fmt.Sprintf("%v 必须为 %v", jsonErr.Field, jsonErr.Type)
			res.Json(c, res.Code(1), res.Msg(msg))
			return
		} else {
			res.Json(c, res.Code(3), res.Msg(err.Error()))
			return
		}
	}
	for _, v := range form.Ids {
		var blogArticle model.BlogArticle
		app.Db("blog").Model(&model.BlogArticle{}).Where("id = ?", v).Take(&blogArticle)
		if blogArticle.Id > 0 {
			needDelImg := editor.ExtractImageURLs(blogArticle.Content)
			for _, v := range needDelImg {
				upload.Del(v, 0)
			}
			err = app.Db("blog").Model(&model.BlogArticle{}).Where("id = ?", v).Delete(&model.BlogArticle{}).Error
			if err != nil {
				res.Json(c, res.Code(11), res.Msg("删除失败"))
				return
			}
		}
	}

	res.Json(c, res.Msg("删除成功"))
	return
}
