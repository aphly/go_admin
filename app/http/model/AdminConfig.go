package model

import (
	"go_admin/app/core"
)

type AdminConfig struct {
	core.ModelId
	Uid   core.Uint `gorm:"index" json:"uid"`
	Title string    `gorm:"size:32" json:"title"`
	Key   string    `gorm:"size:32" json:"key"`
	Value string    `gorm:"size:255" json:"value"`
}
