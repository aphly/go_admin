package model

import (
	"go_admin/app/core"
)

type AdminDict struct {
	core.ModelId
	Uid   core.Int64 `gorm:"index" json:"uid,omitempty"`
	Title string     `gorm:"size:32" json:"title,omitempty"`
	Key   string     `gorm:"index;size:32" json:"key,omitempty"`
	Sort  int        `gorm:"default:0" json:"sort,omitempty"`
}
