package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go_admin/app"
	"go_admin/app/core/crypt"
	"go_admin/app/helper"
	"go_admin/app/http/model"
	"go_admin/app/res"
	"time"
)

func Register(c *gin.Context) {
	userAuthForm := model.UserAuthForm{}
	err := c.ShouldBind(&userAuthForm)
	if err != nil {
		res.Json(c, res.Code(1), res.Msg(userAuthForm.GetError(err.(validator.ValidationErrors))))
		return
	}
	userAuth := model.AdminUserAuth{}
	userAuth.Id = userAuthForm.Id
	userAuth.IdType = "mobile"
	app.Db().Where(&userAuth).Take(&userAuth)
	if userAuth.Uid != 0 {
		res.Json(c, res.Code(1), res.Msg("手机号码已存在"))
		return
	}
	userAuth.Password = crypt.ShaEn(userAuthForm.Password)
	userAuth.LastIp = c.ClientIP()
	userAuth.LastTime = time.Now().Unix()
	userAuth.UserAgent = c.Request.Header.Get("User-Agent")
	userAuth.AcceptLanguage = c.Request.Header.Get("Accept-Language")
	userAuth.Uid = helper.NewSnowflake.NextID()
	result := app.Db().Create(&userAuth)
	if result.Error != nil {
		res.Json(c, res.Code(1), res.Msg("错误"))
		return
	}
	user := &model.AdminUser{}
	err = user.Add(userAuth.Uid)
	if err != nil {
		res.Json(c, res.Code(1), res.Msg("错误1"))
		return
	}
	res.Json(c, res.Data(gin.H{
		"user": gin.H{
			"id_type":       userAuth.IdType,
			"id":            userAuth.Id,
			"uid":           userAuth.Uid,
			"access_token":  user.EnToken(user.AccessToken),
			"refresh_token": user.EnToken(user.RefreshToken),
			"nickname":      user.Nickname,
			"avatar_path":   "",
		},
	}))
	return
}
