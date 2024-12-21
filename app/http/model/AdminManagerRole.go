package model

import (
	"go_admin/app/core"
	"time"
)

type AdminManagerRole struct {
	ManagerUid core.Int64 `gorm:"primarykey" json:"manager_uid,omitempty"`
	RoleId     uint       `gorm:"primarykey" json:"role_id,omitempty"`
	IsHalf     int8       `gorm:"default:0" json:"is_half,omitempty"`
	Uid        core.Int64 `json:"uid,omitempty"`
	CreatedAt  time.Time  `json:"created_at,omitempty"`
	Role       AdminRole  `gorm:"-:migration"`
}
