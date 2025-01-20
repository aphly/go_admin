package account

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go_admin/app"
	"go_admin/app/core/crypt"
	"go_admin/app/helper"
	"go_admin/app/http/admin/form/account"
	"go_admin/app/http/admin/model"
	"go_admin/app/res"
	"time"
)

func Login(c *gin.Context) {
	loginForm := account.LoginForm{}
	err := c.ShouldBind(&loginForm)
	if err != nil {
		if jsonErr, ok := err.(*json.UnmarshalTypeError); ok {
			msg := fmt.Sprintf("%v 必须为 %v", jsonErr.Field, jsonErr.Type)
			res.Json(c, res.Code(1), res.Msg(msg))
			return
		} else if validatorErr, ok1 := err.(validator.ValidationErrors); ok1 {
			res.Json(c, res.Code(2), res.Msg(loginForm.GetError(validatorErr)))
			return
		} else {
			res.Json(c, res.Code(3), res.Msg(err.Error()))
			return
		}
	}
	adminManager := model.AdminManager{}
	adminManager.Username = loginForm.Username
	app.Db().Where(&adminManager).Take(&adminManager)
	if adminManager.Uid == 0 {
		res.Json(c, res.Code(11), res.Msg("用户不存在"))
		loginLog(c, loginForm)
		return
	}
	if adminManager.Password == "" {
		res.Json(c, res.Code(12), res.Msg("请联系管理员"))
		loginLog(c, loginForm)
		return
	}
	if crypt.ShaEn(loginForm.Password) != adminManager.Password {
		res.Json(c, res.Code(13), res.Msg("用户名或密码错误"))
		loginLog(c, loginForm)
		return
	}
	if adminManager.Status == 0 {
		res.Json(c, res.Code(14), res.Msg("用户被冻结"))
		loginLog(c, loginForm)
		return
	}
	loginSave := model.AdminManager{
		LastIp:         c.ClientIP(),
		LastTime:       time.Now().Unix(),
		UserAgent:      c.Request.Header.Get("User-Agent"),
		AcceptLanguage: c.Request.Header.Get("Accept-Language"),
	}
	app.Db().Model(&model.AdminManager{}).Where("uid=?", adminManager.Uid).Updates(loginSave)

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
	}))
	return
}

func loginLog(c *gin.Context, form account.LoginForm) {
	marshal, err := json.Marshal(form)
	if err != nil {
		return
	}
	app.Db().Create(&model.AdminLoginLog{
		Ip:             c.ClientIP(),
		Input:          string(marshal),
		UserAgent:      c.Request.Header.Get("User-Agent"),
		AcceptLanguage: c.Request.Header.Get("Accept-Language"),
	})
}
