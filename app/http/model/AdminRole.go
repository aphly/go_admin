package model

import (
	"go_admin/app/core"
)

type AdminRole struct {
	core.ModelId
	Uid      core.Uint  `gorm:"index" json:"uid"`
	Title    string     `gorm:"size:32" json:"title"`
	Desc     string     `gorm:"size:255" json:"desc"`
	Sort     int        `gorm:"default:0" json:"sort"`
	Status   int8       `gorm:"default:0" json:"status"`
	DataPerm int8       `gorm:"default:0" json:"data_perm"`
	LevelId  uint       `gorm:"default:0" json:"level_id"`
	Level    AdminLevel `gorm:"foreignKey:LevelId;-:migration"`
}
