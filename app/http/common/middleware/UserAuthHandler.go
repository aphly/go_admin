package middleware

import (
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/helper"
	"go_admin/app/http/common/service"
	"go_admin/app/http/common/service/redis"
	"go_admin/app/res"
	"time"
)

func UserAuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := service.GetToken(c)
		if err != nil {
			res.Json(c, res.Code(401), res.Msg(err.Error()))
			c.Abort()
			return
		}

		payloadData, err := helper.ParseJwt(token, []byte(app.Config.AppKey))
		if err != nil {
			res.Json(c, res.Code(401), res.Msg(err.Error()))
			c.Abort()
			return
		}

		if payloadData.TokenType != "access" {
			res.Json(c, res.Code(3), res.Msg("Token 类型错误"))
			c.Abort()
			return
		}

		_, err = redis.IsTokenBlacklisted(c, "jwt:userBlacklist", int64(payloadData.Uid))
		if err != nil {
			res.Json(c, res.Code(402), res.Msg("Token 错误"))
			c.Abort()
			return
		}

		if payloadData.Expire < time.Now().Unix() {
			res.Json(c, res.Code(402), res.Msg("Token 过期"))
			c.Abort()
			return
		}
		c.Set("uid", payloadData.Uid)
		c.Next()
	}
}
