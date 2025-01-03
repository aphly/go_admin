package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go_admin/app"
	"go_admin/app/core/crypt"
	"go_admin/app/http/model"
	"go_admin/app/http/service"
	"go_admin/app/res"
)

func Login(c *gin.Context) {
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
	if userAuth.Uid == 0 {
		res.Json(c, res.Code(1), res.Msg("手机号码不存在"))
		return
	}
	if crypt.ShaEn(userAuthForm.Password) != userAuth.Password {
		res.Json(c, res.Code(1), res.Msg("密码错误"))
		return
	}
	user := &model.AdminUser{
		Uid: userAuth.Uid,
	}
	app.Db().Where(&user).Take(&user)
	user.GenToken()
	app.Db().Save(&user)

	res.Json(c, res.Data(gin.H{
		"user": gin.H{
			"id_type":       userAuth.IdType,
			"id":            userAuth.Id,
			"uid":           userAuth.Uid,
			"access_token":  user.EnToken(user.AccessToken),
			"refresh_token": user.EnToken(user.RefreshToken),
			"nickname":      user.Nickname,
			"avatar_path":   service.UploadPath(user.Avatar, user.Remote),
		},
	}))
	return
}
