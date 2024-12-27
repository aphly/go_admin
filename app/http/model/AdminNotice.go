package model

import (
	"go_admin/app/core"
)

type AdminNotice struct {
	core.ModelId
	Uid     core.Int64 `gorm:"index" json:"uid,omitempty"`
	Title   string     `gorm:"size:32" json:"title,omitempty"`
	Content string     `gorm:"type:text" json:"content,omitempty"`
	Status  int8       `gorm:"default:0" json:"status"`
	LevelId uint       `gorm:"index" json:"level_id"`
	Level   AdminLevel `gorm:"foreignKey:LevelId;-:migration"`
}
