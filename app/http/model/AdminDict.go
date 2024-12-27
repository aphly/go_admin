package model

import (
	"go_admin/app/core"
)

type AdminDict struct {
	core.ModelId
	Uid    core.Int64 `gorm:"index" json:"uid"`
	Title  string     `gorm:"size:32" json:"title"`
	Name   string     `gorm:"index;size:32" json:"name"`
	Sort   int        `gorm:"default:0" json:"sort"`
	Status int8       `gorm:"default:0" json:"status"`
}
