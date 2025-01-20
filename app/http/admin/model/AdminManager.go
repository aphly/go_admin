package model

import (
	"go_admin/app"
	"go_admin/app/core"
)

type AdminManager struct {
	Uid core.Uint `gorm:"primarykey" json:"uid"`
	//LevelId        uint      `gorm:"index" json:"level_id"`
	Username       string `gorm:"size:32;unique" json:"username"`
	Nickname       string `gorm:"size:32" json:"nickname"`
	Phone          string `gorm:"size:32" json:"phone"`
	Password       string `gorm:"size:255" json:"-"`
	LastIp         string `gorm:"size:64" json:"last_ip"`
	LastTime       int64  `gorm:"default:0" json:"last_time"`
	Status         int8   `gorm:"default:1" json:"status"`
	Note           string `gorm:"size:255" json:"note"`
	Avatar         string `gorm:"size:255" json:"avatar"`
	Remote         int8   `gorm:"default:0" json:"remote"`
	UserAgent      string `gorm:"size:255" json:"user_agent"`
	AcceptLanguage string `gorm:"size:255" json:"accept_language"`
	core.Model
	//ManagerRole []AdminManagerRole `gorm:"foreignKey:Uid;-:migration"`
	//Level AdminLevel `gorm:"foreignKey:LevelId;-:migration"`
}

func (this *AdminManager) Add(uid core.Uint, username, nickname, password, phone string) (error, *AdminManager) {
	this.Uid = uid
	this.Username = username
	this.Nickname = nickname
	this.Password = password
	this.Phone = phone
	err := app.Db().Create(this).Error
	if err != nil {
		return err, nil
	}
	return nil, this
}
