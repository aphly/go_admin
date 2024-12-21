package model

import (
	"go_admin/app/core"
	"time"
)

type AdminRoleMenu struct {
	RoleId    uint       `gorm:"primarykey" json:"role_id"`
	MenuId    uint       `gorm:"primarykey" json:"menu_id"`
	IsHalf    int8       `gorm:"default:0" json:"is_half"`
	Uid       core.Int64 `json:"uid,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	Menu      AdminMenu  `gorm:"-:migration"`
}
