package account

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go_admin/app"
	"go_admin/app/core/crypt"
	"go_admin/app/helper"
	"go_admin/app/http/common/form/front/account"
	"go_admin/app/http/common/model"
	"go_admin/app/res"
	"time"
)

func Forget(c *gin.Context) {
	form := account.LoginForm{}
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
	userAuth := model.UserAuth{}
	userAuth.Id = form.Id
	userAuth.IdType = form.IdType
	app.Db("common").Where(&userAuth).Take(&userAuth)
	if userAuth.Uid == 0 {
		res.Json(c, res.Code(11), res.Msg("用户不存在"))
		return
	}

	if crypt.ShaEn(form.Password) != userAuth.Password {
		res.Json(c, res.Code(13), res.Msg("用户名或密码错误"))
		return
	}

	userModel := model.User{}
	userModel.Uid = userAuth.Uid
	app.Db("common").Where(&userModel).Take(&userModel)

	if userModel.Status == 0 {
		res.Json(c, res.Code(14), res.Msg("用户被冻结"))
		return
	}
	loginSave := model.UserAuth{
		LastIp:         c.ClientIP(),
		LastTime:       time.Now().Unix(),
		UserAgent:      c.Request.Header.Get("User-Agent"),
		AcceptLanguage: c.Request.Header.Get("Accept-Language"),
	}
	app.Db("common").Model(&model.UserAuth{}).Where("uid=? and id_type=?", userModel.Uid, userAuth.IdType).Updates(loginSave)
	access_token, _ := helper.CreateToken([]byte(app.Config.AppKey), helper.PayloadData{time.Now().Add(time.Hour * 2).Unix(), userModel.Uid, "access"})
	refresh_token, _ := helper.CreateToken([]byte(app.Config.AppKey), helper.PayloadData{time.Now().Add(time.Hour * 24 * 365).Unix(), userModel.Uid, "refresh"})
	res.Json(c, res.Data(gin.H{
		"user": gin.H{
			"uid":      userModel.Uid,
			"nickname": userModel.Nickname,
			"avatar":   userModel.Avatar,
		},
		"token": gin.H{
			"access_token":  access_token,
			"refresh_token": refresh_token,
		},
	}))
	return
}
