package model

import (
	"go_admin/app/core"
)

type AdminApi struct {
	core.ModelId
	Uid    core.Int64 `gorm:"index" json:"uid,omitempty"`
	Title  string     `gorm:"size:32" json:"title,omitempty"`
	Path   string     `gorm:"size:255" json:"path,omitempty"`
	Pid    uint       `gorm:"default:0" json:"pid"`
	Sort   int        `gorm:"default:0" json:"sort"`
	Status int8       `gorm:"default:0" json:"status"`
	Type   int8       `gorm:"default:0" json:"type"`
}
