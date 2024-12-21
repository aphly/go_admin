package model

import (
	"go_admin/app/core"
)

type AdminRole struct {
	core.ModelId
	Uid      core.Int64 `gorm:"index" json:"uid,omitempty"`
	Title    string     `gorm:"size:32" json:"title,omitempty"`
	Desc     string     `gorm:"size:255" json:"desc,omitempty"`
	Sort     int        `gorm:"default:0" json:"sort"`
	Status   int8       `gorm:"default:0" json:"status"`
	DataPerm int8       `gorm:"default:0" json:"data_perm,omitempty"`
	LevelId  uint       `gorm:"default:0" json:"level_id,omitempty"`
	Level    AdminLevel `gorm:"-:migration"`
}
