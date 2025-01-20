package category

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go_admin/app"
	"go_admin/app/core"
	"go_admin/app/http/blog/form/category"
	"go_admin/app/http/blog/model"
	"go_admin/app/res"
)

func Add(c *gin.Context) {
	uid, _ := c.Get("uid")
	uidD := uid.(core.Uint)
	form := category.Form{}
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
	err = app.Db("blog").Create(&model.BlogCategory{
		Uid:             uidD,
		Title:           form.Title,
		Name:            form.Name,
		Pid:             *form.Pid,
		Sort:            *form.Sort,
		Status:          *form.Status,
		Type:            form.Type,
		MetaTitle:       form.MetaTitle,
		MetaDescription: form.MetaDescription,
	}).Error
	if err != nil {
		res.Json(c, res.Code(11), res.Msg("保存失败"))
		return
	}
	res.Json(c, res.Msg("保存成功"))
	return
}
