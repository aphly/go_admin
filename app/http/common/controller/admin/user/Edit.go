package user

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go_admin/app"
	"go_admin/app/core/crypt"
	"go_admin/app/http/common/form/admin/user"
	"go_admin/app/http/common/model"
	"go_admin/app/res"
)

func Edit(c *gin.Context) {
	//uid, _ := c.Get("uid")
	form := &user.Form{}
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
	db := app.Db("common").Model(&model.User{}).Where("uid=?", form.Uid)
	var userModel model.User
	err = db.Take(&userModel).Error
	if userModel.Uid > 0 {
		if form.Password != "" {
			err = app.Db("common").Model(&model.UserAuth{}).Where("uid=?", form.Uid).Updates(model.UserAuth{
				Password: crypt.ShaEn(form.Password),
			}).Error
			if err != nil {
				res.Json(c, res.Code(11), res.Msg("保存失败"))
				return
			}
		}

		err = db.Updates(model.User{
			Nickname: form.Nickname,
		}).Error
		if err != nil {
			res.Json(c, res.Code(11), res.Msg("保存失败"))
			return
		}
	}

	res.Json(c, res.Msg("保存成功"))
	return
}
