package wechat

import (
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/helper"
	sign2 "go_admin/app/helper/sign"
	"go_admin/app/http/common/model"
	"go_admin/app/http/common/service/upload"
	"go_admin/app/res"
	"time"
)

func Login(c *gin.Context) {
	rand_str := c.DefaultQuery("rand_str", "")
	if rand_str == "" {
		res.Json(c, res.Code(1), res.Msg("错误"))
		return
	}
	sign := c.DefaultQuery("sign", "")
	if sign == "" {
		res.Json(c, res.Code(2), res.Msg("sign错误"))
		return
	}
	signStr := sign2.String(rand_str, app.Config.AppKey)
	if sign != signStr {
		res.Json(c, res.Code(3), res.Msg("签名错误"))
		return
	}

	wechatScan := model.WechatScan{
		RandStr: rand_str,
	}

	err := app.Db("common").Where(&wechatScan).Take(&wechatScan).Error
	if err != nil {
		res.Json(c, res.Code(11), res.Msg("错误"))
		return
	}
	if wechatScan.Openid == "" {
		res.Json(c, res.Code(12), res.Msg("错误"))
		return
	}
	if wechatScan.Status == 0 {
		res.Json(c, res.Code(13), res.Msg("错误"))
		return
	}
	userAuth := model.UserAuth{}
	userAuth.Id = wechatScan.Openid
	userAuth.IdType = "wechat"
	app.Db("common").Where(&userAuth).Take(&userAuth)
	if userAuth.Uid == 0 {
		res.Json(c, res.Code(11), res.Msg("用户不存在"))
		return
	}

	userModel := model.User{}
	userModel.Uid = userAuth.Uid
	app.Db("common").Where(&userModel).Take(&userModel)

	if userModel.Status == 0 {
		res.Json(c, res.Code(14), res.Msg("用户被冻结"))
		return
	}

	access_token, _ := helper.CreateToken([]byte(app.Config.AppKey), helper.PayloadData{time.Now().Add(time.Hour * 2).Unix(), userModel.Uid, "access"})
	refresh_token, _ := helper.CreateToken([]byte(app.Config.AppKey), helper.PayloadData{time.Now().Add(time.Hour * 24 * 365).Unix(), userModel.Uid, "refresh"})
	avatar := upload.Url(userModel.Avatar, userModel.Remote)
	if avatar == "" {
		avatar = app.Config.Http.Host + "/public/img/avatar.png"
	}
	res.Json(c, res.Data(gin.H{
		"user": gin.H{
			"uid":      userModel.Uid,
			"nickname": userModel.Nickname,
			"avatar":   avatar,
		},
		"token": gin.H{
			"access_token":  access_token,
			"refresh_token": refresh_token,
		},
	}))
	return
}
