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
	"gorm.io/gorm"
	"time"
)

func Register(c *gin.Context) {
	form := account.RegisterForm{}
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
	if userAuth.Uid > 0 {
		res.Json(c, res.Code(11), res.Msg("用户已存在"))
		return
	}

	userAuth.Uid = helper.NewSnowflake.NextId()
	userModel := model.User{
		Uid:      userAuth.Uid,
		Nickname: helper.RandStr(8),
	}
	err = app.Db("common").Transaction(func(tx *gorm.DB) error {
		userAuth.Password = crypt.ShaEn(form.Password)
		userAuth.LastIp = c.ClientIP()
		userAuth.LastTime = time.Now().Unix()
		userAuth.UserAgent = c.Request.Header.Get("User-Agent")
		userAuth.AcceptLanguage = c.Request.Header.Get("Accept-Language")
		err = tx.Create(&userAuth).Error
		if err != nil {
			return err
		}

		err = tx.Create(&userModel).Error
		if err != nil {
			return err
		}
		return nil
	})

	access_token, _ := helper.CreateToken([]byte(app.Config.AppKey), helper.PayloadData{time.Now().Add(time.Hour * 2).Unix(), userModel.Uid, "access"})
	refresh_token, _ := helper.CreateToken([]byte(app.Config.AppKey), helper.PayloadData{time.Now().Add(time.Hour * 24 * 365).Unix(), userModel.Uid, "refresh"})
	res.Json(c, res.Data(gin.H{
		"user": gin.H{
			"uid":      userModel.Uid,
			"nickname": userModel.Nickname,
			"avatar":   app.Config.Http.Host + "/public/img/avatar.png",
		},
		"token": gin.H{
			"access_token":  access_token,
			"refresh_token": refresh_token,
		},
	}))
	return
}
