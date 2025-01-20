package model

import (
	"go_admin/app/core"
	"time"
)

type AdminManagerRole struct {
	ManagerUid core.Uint `gorm:"primarykey" json:"manager_uid"`
	RoleId     uint      `gorm:"primarykey" json:"role_id"`
	IsHalf     int8      `gorm:"default:0" json:"is_half"`
	Uid        core.Uint `json:"uid"`
	CreatedAt  time.Time `json:"created_at"`
	Role       AdminRole `gorm:"-:migration"`
}
