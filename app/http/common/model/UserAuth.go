package model

import (
	"go_admin/app/core"
)

type UserAuth struct {
	IdType         string    `gorm:"primaryKey;type:char(16)" json:"id_type"`
	Id             string    `gorm:"primaryKey;size:64" json:"id"`
	Password       string    `gorm:"size:255" json:"-"`
	Uid            core.Uint `gorm:"index" json:"uid"`
	LastIp         string    `gorm:"size:64" json:"last_ip"`
	LastTime       int64     `json:"last_time"`
	Note           string    `gorm:"size:255" json:"-"`
	UserAgent      string    `gorm:"type:text" json:"-"`
	AcceptLanguage string    `gorm:"size:255" json:"-"`
	Verified       int8      `gorm:"default:0" json:"verified"`
	core.Model
}
