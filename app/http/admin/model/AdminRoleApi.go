package model

import (
	"go_admin/app/core"
	"time"
)

type AdminRoleApi struct {
	RoleId    uint      `gorm:"primarykey" json:"role_id"`
	ApiId     uint      `gorm:"primarykey" json:"api_id"`
	IsHalf    int8      `gorm:"default:0" json:"is_half"`
	Uid       core.Uint `json:"uid"`
	CreatedAt time.Time `json:"created_at"`
	Api       AdminApi  `gorm:"-:migration"`
}
