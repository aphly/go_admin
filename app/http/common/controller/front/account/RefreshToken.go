package account

import (
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/helper"
	"go_admin/app/http/common/model"
	"go_admin/app/http/common/service"
	"go_admin/app/res"
	"time"
)

func RefreshToken(c *gin.Context) {
	token, err := service.GetToken(c)
	if err != nil {
		res.Json(c, res.Code(1), res.Msg(err.Error()))
		return
	}

	payloadData, err := helper.ParseJwt(token, []byte(app.Config.AppKey))
	if err != nil {
		res.Json(c, res.Code(2), res.Msg(err.Error()))
		return
	}

	if payloadData.TokenType != "refresh" {
		res.Json(c, res.Code(3), res.Msg("Token 类型错误"))
		return
	}

	if payloadData.Expire < time.Now().Unix() {
		res.Json(c, res.Code(4), res.Msg("Refresh Token 过期"))
		return
	}

	userModel := model.User{
		Uid: payloadData.Uid,
	}
	app.Db("common").Where(&userModel).Take(&userModel)

	if userModel.Status != 1 {
		res.Json(c, res.Code(5), res.Msg("用户被冻结"))
		return
	}
	access_token, _ := helper.CreateToken([]byte(app.Config.AppKey), helper.PayloadData{time.Now().Add(time.Hour * 2).Unix(), userModel.Uid, "access"})
	res.Json(c, res.Data(gin.H{
		"user": gin.H{
			"uid":          userModel.Uid,
			"access_token": access_token,
		},
	}))
	return
}
