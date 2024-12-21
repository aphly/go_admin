package model

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/core"
	"go_admin/app/core/crypt"
	"go_admin/app/helper"
	"strconv"
	"time"
)

const (
	AccessTokenExpire  = 300
	RefreshTokenExpire = 31536000
)

type AdminUser struct {
	Uid                core.Int64 `gorm:"primarykey" json:"uid,omitempty"`
	Nickname           string     `gorm:"size:16" json:"nickname,omitempty"`
	AccessToken        string     `gorm:"index;size:64" json:"access_token,omitempty"`
	RefreshToken       string     `gorm:"index;size:64" json:"refresh_token,omitempty"`
	AccessTokenExpire  int64      `gorm:"default:0" json:"-"`
	RefreshTokenExpire int64      `gorm:"default:0" json:"-"`
	Avatar             string     `gorm:"size:255" json:"avatar,omitempty"`
	Remote             int8       `gorm:"default:0" json:"remote,omitempty"`
	Status             int8       `gorm:"default:1" json:"status,omitempty"`
	core.Model
}

func (this AdminUser) GetToken(c *gin.Context) (error, string) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		return errors.New("No token"), ""
	}
	if len(token) < 7 || token[:6] != "Bearer " {
		return errors.New("Invalid token"), ""
	}
	return nil, token[7:]
}

func (this *AdminUser) Add(uid core.Int64) error {
	this.Uid = uid
	this.Nickname = helper.RandStr(10)
	now := time.Now().Unix()
	this.AccessToken = helper.RandStr(32)
	this.RefreshToken = helper.RandStr(32, 1)
	this.AccessTokenExpire = now + AccessTokenExpire
	this.RefreshTokenExpire = now + RefreshTokenExpire
	app.Db().Create(this)
	return nil
}

func (this *AdminUser) EnToken(token string) string {
	uidStr := strconv.FormatInt(int64(this.Uid), 10)
	en, _ := crypt.AesEn(uidStr + "_" + token)
	return en
}

func (this *AdminUser) DeToken(token string) (string, error) {
	de, err := crypt.AesDe(token)
	if err != nil {
		return "", err
	}
	return de, nil
}

func (this *AdminUser) GenToken() {
	now := time.Now().Unix()
	this.AccessToken = helper.RandStr(32)
	this.RefreshToken = helper.RandStr(32, 1)
	this.AccessTokenExpire = now + AccessTokenExpire
	this.RefreshTokenExpire = now + RefreshTokenExpire
}

func (this *AdminUser) GenAccessToken(now int64) {
	this.AccessToken = helper.RandStr(32)
	this.AccessTokenExpire = now + AccessTokenExpire
}
