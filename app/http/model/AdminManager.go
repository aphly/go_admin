package model

import (
	"go_admin/app"
	"go_admin/app/core"
)

type AdminManager struct {
	Uid            core.Int64 `gorm:"primarykey" json:"uid,omitempty"`
	LevelId        uint       `gorm:"index" json:"level_id,omitempty"`
	Username       string     `gorm:"size:32;unique" json:"username,omitempty"`
	Nickname       string     `gorm:"size:32" json:"nickname,omitempty"`
	Phone          string     `gorm:"size:32" json:"phone"`
	Password       string     `gorm:"size:255" json:"-"`
	LastIp         string     `gorm:"size:64" json:"last_ip,omitempty"`
	LastTime       int64      `gorm:"default:0" json:"last_time,omitempty"`
	Status         int8       `gorm:"default:1" json:"status,omitempty"`
	Note           string     `gorm:"size:255" json:"note,omitempty"`
	Avatar         string     `gorm:"size:255" json:"avatar"`
	Remote         int8       `gorm:"default:0" json:"remote"`
	UserAgent      string     `gorm:"size:255" json:"user_agent,omitempty"`
	AcceptLanguage string     `gorm:"size:255" json:"accept_language,omitempty"`
	core.Model
	AdminManagerRole []AdminManagerRole `gorm:"foreignKey:Uid;-:migration"`
}

func (this *AdminManager) Add(uid core.Int64, username, nickname, password, phone string) (error, *AdminManager) {
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