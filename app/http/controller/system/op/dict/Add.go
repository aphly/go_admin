package dict

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go_admin/app"
	"go_admin/app/core"
	"go_admin/app/http/form/system/perm/dict"
	"go_admin/app/http/model"
	"go_admin/app/res"
)

func Add(c *gin.Context) {
	uid, _ := c.Get("uid")
	uid_d := uid.(core.Int64)
	form := dict.Form{}
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
	err = app.Db().Create(&model.AdminDict{
		Uid:    uid_d,
		Title:  form.Title,
		Name:   form.Name,
		Sort:   *form.Sort,
		Status: *form.Status,
	}).Error
	if err != nil {
		res.Json(c, res.Code(11), res.Msg("保存失败"))
		return
	}
	res.Json(c, res.Msg("保存成功"))
	return
}