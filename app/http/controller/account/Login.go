package account

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go_admin/app"
	"go_admin/app/core/crypt"
	"go_admin/app/helper"
	"go_admin/app/http/form/account"
	"go_admin/app/http/model"
	"go_admin/app/res"
	"time"
)

func Login(c *gin.Context) {
	loginForm := account.LoginForm{}
	err := c.ShouldBind(&loginForm)
	if err != nil {
		res.Json(c, res.Code(1), res.Msg(loginForm.GetError(err.(validator.ValidationErrors))))
		return
	}
	adminManager := model.AdminManager{}
	adminManager.Username = loginForm.Username
	app.Db().Where(&adminManager).Take(&adminManager)
	if adminManager.Uid == 0 {
		res.Json(c, res.Code(1), res.Msg("用户不存在"))
		return
	}
	if crypt.ShaEn(loginForm.Password) != adminManager.Password {
		res.Json(c, res.Code(1), res.Msg("用户名或密码错误"))
		return
	}
	app.Db().Save(&adminManager)
	var adminManagerRole []model.AdminManagerRole
	app.Db().InnerJoins("Role", app.Db().Where(&model.AdminRole{Status: 1})).
		Where("admin_manager_role.manager_uid = ?", adminManager.Uid).Find(&adminManagerRole)
	access_token, _ := helper.CreateToken([]byte(app.Config.AppKey), helper.PayloadData{time.Now().Add(time.Hour * 2).Unix(), adminManager.Uid, "access"})
	refresh_token, _ := helper.CreateToken([]byte(app.Config.AppKey), helper.PayloadData{time.Now().Add(time.Hour * 24 * 365).Unix(), adminManager.Uid, "refresh"})
	res.Json(c, res.Data(gin.H{
		"manager": gin.H{
			"uid":      adminManager.Uid,
			"nickname": adminManager.Nickname,
			"username": adminManager.Username,
			"avatar":   adminManager.Avatar,
			"phone":    adminManager.Phone,
		},
		"token": gin.H{
			"access_token":  access_token,
			"refresh_token": refresh_token,
		},
		"manager_role": adminManagerRole,
	}))
	return
}
