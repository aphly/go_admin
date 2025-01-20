package wechat

import (
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/helper"
	sign2 "go_admin/app/helper/sign"
	"go_admin/app/http/common/model"
	"go_admin/app/res"
)

func Scan(c *gin.Context) {
	wechatScan := model.WechatScan{
		RandStr: helper.RandStr(32),
		Openid:  "",
		Status:  0,
	}
	err := app.Db("common").Create(&wechatScan).Error
	if err != nil {
		res.Json(c, res.Code(11), res.Msg("错误"))
		return
	}
	signStr := sign2.String(wechatScan.RandStr, app.Config.AppKey)
	res.Json(c, res.Data(gin.H{
		"RandStr": wechatScan.RandStr,
		"sign":    signStr,
	}))
	return
}
