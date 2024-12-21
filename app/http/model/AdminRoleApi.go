package model

import (
	"go_admin/app/core"
	"time"
)

type AdminRoleApi struct {
	RoleId    uint       `gorm:"primarykey" json:"role_id"`
	ApiId     uint       `gorm:"primarykey" json:"api_id"`
	IsHalf    int8       `gorm:"default:0" json:"is_half"`
	Uid       core.Int64 `json:"uid,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	Api       AdminApi   `gorm:"-:migration"`
}
