package manager

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go_admin/app"
	"go_admin/app/core/crypt"
	"go_admin/app/helper"
	"go_admin/app/http/admin/form/system/perm/manager"
	"go_admin/app/http/admin/model"
	"go_admin/app/res"
)

func Add(c *gin.Context) {
	form := &manager.Form{}
	err := c.ShouldBind(form)
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
	adminManager := model.AdminManager{}
	adminManager.Uid = helper.NewSnowflake.NextId()
	adminManager.Username = form.Username
	adminManager.Nickname = form.Nickname
	if form.Password != "" {
		adminManager.Password = crypt.ShaEn(form.Password)
	}
	adminManager.Phone = form.Phone
	//adminManager.LevelId = form.LevelId
	err = app.Db().Create(&adminManager).Error
	if err != nil {
		res.Json(c, res.Code(11), res.Msg("保存错误"))
		return
	}
	res.Json(c, res.Data(gin.H{
		"manager": gin.H{
			"uid":      adminManager.Uid,
			"username": adminManager.Username,
			"nickname": adminManager.Nickname,
			"phone":    adminManager.Phone,
		},
	}))
	return
}
