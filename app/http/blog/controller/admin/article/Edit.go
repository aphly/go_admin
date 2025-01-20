package article

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go_admin/app"
	"go_admin/app/helper/editor"
	"go_admin/app/http/blog/form/article"
	"go_admin/app/http/blog/model"
	"go_admin/app/http/common/service/upload"
	"go_admin/app/res"
)

func Edit(c *gin.Context) {
	//uid, _ := c.Get("uid")
	form := &article.Form{}
	err := c.ShouldBind(&form)
	if err != nil {
		if jsonErr, ok := err.(*json.UnmarshalTypeError); ok {
			msg := fmt.Sprintf("%v 必须为 %v", jsonErr.Field, jsonErr.Type)
			res.Json(c, res.Code(1), res.Msg(msg))
			return
		} else if validatorErr, ok1 := err.(validator.ValidationErrors); ok1 {
			res.Json(c, res.Code(2), res.Msg(form.GetError(validatorErr)))
			return
		} else {
			res.Json(c, res.Code(3), res.Msg(err.Error()))
			return
		}
	}
	db := app.Db("blog").Model(&model.BlogArticle{}).Where("id=?", form.Id)
	var blogArticle model.BlogArticle
	err = db.Take(&blogArticle).Error
	if blogArticle.Id > 0 {
		needDelImg := editor.ExtractImageURLs(blogArticle.Content)
		for _, v := range needDelImg {
			upload.Del(v, 0)
		}
		content, err := editor.TempToImg(form.Content, "/public/upload/temp/article/", "/public/upload/article/")
		if err != nil {
			res.Json(c, res.Code(4), res.Msg(err.Error()))
			return
		}
		form.Content = content
		err = db.Updates(form).Error
		if err != nil {
			res.Json(c, res.Code(11), res.Msg("保存失败"))
			return
		}
	}

	res.Json(c, res.Msg("保存成功"))
	return
}
