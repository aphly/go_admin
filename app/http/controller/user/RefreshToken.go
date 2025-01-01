package user

import (
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/core"
	"go_admin/app/core/crypt"
	"go_admin/app/http/model"
	"go_admin/app/http/service"
	"go_admin/app/res"
	"strconv"
	"strings"
	"time"
)

func RefreshToken(c *gin.Context) {
	token, err := service.GetToken(c)
	if err != nil {
		res.Json(c, res.Code(1), res.Msg(err.Error()))
		return
	}
	de, err := crypt.AesDe(token)
	if err != nil {
		res.Json(c, res.Code(2), res.Msg("Token 错误"))
		return
	}

	uid_token := strings.Split(de, "_")
	parseInt, err := strconv.ParseInt(uid_token[0], 10, 64)
	if err != nil {
		res.Json(c, res.Code(3), res.Msg("Token 错误"))
		return
	}
	user := model.AdminUser{
		Uid: core.Uint(parseInt),
	}
	app.Db().Where(&user).Take(&user)
	if uid_token[1] != user.RefreshToken {
		res.Json(c, res.Code(4), res.Msg("Token 错误"))
		return
	}
	now := time.Now().Unix()
	if user.RefreshTokenExpire < now {
		res.Json(c, res.Code(5), res.Msg("Refresh Token 过期"))
		return
	}
	user.GenAccessToken(now)
	app.Db().Save(&user)
	res.Json(c, res.Data(gin.H{
		"user": gin.H{
			"uid":          user.Uid,
			"access_token": user.EnToken(user.AccessToken),
		},
	}))
	return
}
