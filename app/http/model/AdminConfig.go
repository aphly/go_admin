package model

import (
	"go_admin/app/core"
)

type AdminConfig struct {
	core.ModelId
	Uid   core.Int64 `gorm:"index" json:"uid,omitempty"`
	Title string     `gorm:"size:32" json:"title,omitempty"`
	Key   string     `gorm:"size:32" json:"key,omitempty"`
	Value string     `gorm:"size:255" json:"value,omitempty"`
}
