package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go_admin/app"
	"go_admin/app/core/crypt"
	"go_admin/app/helper"
	"go_admin/app/http/common/form/admin/user"
	"go_admin/app/http/common/model"
	"go_admin/app/res"
	"gorm.io/gorm"
)

func Add(c *gin.Context) {
	form := &user.Form{}
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
	userAuth := model.UserAuth{
		IdType:   "mobile",
		Id:       form.Id,
		Password: crypt.ShaEn(form.Password),
		Uid:      helper.NewSnowflake.NextId(),
	}

	err = app.Db("common").Transaction(func(tx *gorm.DB) error {
		err = tx.Create(&userAuth).Error
		if err != nil {
			return errors.New(err.Error())
		}
		userModel := model.User{
			Uid:      userAuth.Uid,
			Nickname: form.Nickname,
		}
		err = tx.Create(&userModel).Error
		if err != nil {
			return errors.New("保存错误2")
		}
		return nil
	})

	if err != nil {
		res.Json(c, res.Code(11), res.Msg(err.Error()))
		return
	}
	res.Json(c, res.Msg("保存成功"))
	return
}
