package wechat

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/helper"
	sign2 "go_admin/app/helper/sign"
	"go_admin/app/http/common/model"
	"go_admin/app/res"
	"gorm.io/gorm"
	"io"
	"net/http"
	"time"
)

type WechatData struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
}

func Openid(c *gin.Context) {
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

	code := c.DefaultQuery("code", "")
	if code == "" {
		res.Json(c, res.Code(4), res.Msg("code错误"))
		return
	}

	url := "https://api.weixin.qq.com/sns/oauth2/access_token?appid=" +
		app.Config.Wechat.Appid + "&secret=" + app.Config.Wechat.AppSecret + "&code=" + code + "&grant_type=authorization_code"
	resp, err := http.Get(url)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		res.Json(c, res.Code(11), res.Msg("错误"))
		return
	}
	var data WechatData
	if err = json.Unmarshal(body, &data); err != nil {
		res.Json(c, res.Code(12), res.Msg("错误"))
		return
	}

	wechatScan := model.WechatScan{
		RandStr: rand_str,
	}

	err = app.Db("common").Where(&wechatScan).Updates(model.WechatScan{
		Openid: data.Openid,
		Status: 1,
	}).Error

	if err != nil {
		res.Json(c, res.Code(13), res.Msg("错误"))
		return
	}
	userAuth := model.UserAuth{}
	userAuth.Id = data.Openid
	userAuth.IdType = "wechat"
	app.Db("common").Where(&userAuth).Take(&userAuth)
	if userAuth.Uid > 0 {
		res.Json(c, res.Code(11), res.Msg("用户已存在"))
		return
	}

	userAuth.Uid = helper.NewSnowflake.NextId()
	userModel := model.User{
		Uid:      userAuth.Uid,
		Nickname: helper.RandStr(8),
	}

	err = app.Db("common").Transaction(func(tx *gorm.DB) error {
		userAuth.Password = helper.RandStr(32)
		userAuth.LastIp = c.ClientIP()
		userAuth.LastTime = time.Now().Unix()
		userAuth.UserAgent = c.Request.Header.Get("User-Agent")
		userAuth.AcceptLanguage = c.Request.Header.Get("Accept-Language")
		err = tx.Create(&userAuth).Error
		if err != nil {
			return err
		}

		err = tx.Create(&userModel).Error
		if err != nil {
			return err
		}
		return nil
	})

	res.Json(c, res.Msg("已同意"))
	return
}
