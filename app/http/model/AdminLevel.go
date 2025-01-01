package model

import (
	"go_admin/app/core"
)

type AdminLevel struct {
	core.ModelId
	Uid    core.Uint `gorm:"index" json:"uid"`
	Title  string    `gorm:"size:32" json:"title"`
	Pid    uint      `gorm:"default:0" json:"pid"`
	Sort   int       `gorm:"default:0" json:"sort"`
	Status int8      `gorm:"default:0" json:"status"`
}
