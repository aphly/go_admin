package role

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go_admin/app"
	"go_admin/app/core"
	"go_admin/app/http/admin/form/system/perm/role"
	"go_admin/app/http/admin/model"
	"go_admin/app/res"
)

func Add(c *gin.Context) {
	uid, _ := c.Get("uid")
	uid_d := uid.(core.Uint)
	form := role.Form{}
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
	err = app.Db().Create(&model.AdminRole{
		Uid:      uid_d,
		Title:    form.Title,
		Desc:     form.Desc,
		Sort:     *form.Sort,
		Status:   *form.Status,
		DataPerm: form.DataPerm,
		LevelId:  form.LevelId,
	}).Error
	if err != nil {
		res.Json(c, res.Code(11), res.Msg("保存失败"))
		return
	}
	res.Json(c, res.Msg("保存成功"))
	return
}
