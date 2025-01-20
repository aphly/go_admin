package model

import (
	"go_admin/app/core"
)

type AdminUpload struct {
	core.ModelId
	Uid      core.Uint    `gorm:"index" json:"uid"`
	LevelId  uint         `gorm:"index" json:"level_id"`
	Path     string       `gorm:"size:255" json:"path"`
	FileType string       `gorm:"size:255" json:"file_type"`
	FileSize int64        `gorm:"default:0" json:"file_size"`
	Remote   int8         `gorm:"default:0" json:"remote"`
	Level    AdminLevel   `gorm:"foreignKey:LevelId;-:migration"`
	Manager  AdminManager `gorm:"foreignKey:Uid;references:Uid;-:migration"`
}
