package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/core"
	"go_admin/app/helper"
	"go_admin/app/http/admin/model"
	"go_admin/app/http/common/service"
	"go_admin/app/http/common/service/redis"
	"go_admin/app/res"
	"io"
	"time"
)

func ManagerAuthHandler() gin.HandlerFunc {
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

		_, err = redis.IsTokenBlacklisted(c, "jwt:managerBlacklist", int64(payloadData.Uid))
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
		operationLog(payloadData.Uid, c)
		c.Next()

	}
}

func operationLog(uid core.Uint, c *gin.Context) {
	if c.Request.Method == "POST" && c.Request.URL.Path != "/admin/system/op/operation/index" && c.Request.URL.Path != "/admin/system/op/operation/del" {
		contentType := c.Request.Header.Get("Content-Type")
		RequestData := ""
		if contentType == "application/json" || contentType == "application/x-www-form-urlencoded" {
			body, err := c.GetRawData()
			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			if err == nil {
				RequestData = string(body)
			}
		}
		app.Db().Create(&model.AdminOperation{
			Uid:         uid,
			Url:         c.Request.URL.String(),
			Method:      c.Request.Method,
			RequestData: RequestData,
			Ip:          c.ClientIP(),
		})
	}

}
