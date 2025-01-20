package model

import (
	"go_admin/app/core"
)

type User struct {
	Uid      core.Uint `gorm:"primarykey" json:"uid"`
	Nickname string    `gorm:"size:16" json:"nickname"`
	Avatar   string    `gorm:"size:255" json:"avatar,omitempty"`
	Remote   int8      `gorm:"default:0" json:"remote,omitempty"`
	Status   int8      `gorm:"default:1" json:"status,omitempty"`
	core.Model
	UserAuth []UserAuth `gorm:"foreignKey:Uid;-:migration"`
}
