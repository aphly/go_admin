package model

import (
	"go_admin/app/core"
)

type AdminNotice struct {
	core.ModelId
	Uid       core.Uint    `gorm:"index" json:"uid"`
	Title     string       `gorm:"size:32" json:"title"`
	Content   string       `gorm:"type:text" json:"content"`
	Status    int8         `gorm:"default:0" json:"status"`
	DictValue string       `gorm:"index" json:"dict_value"`
	LevelId   uint         `gorm:"index" json:"level_id"`
	Level     AdminLevel   `gorm:"foreignKey:LevelId;-:migration"`
	Manager   AdminManager `gorm:"foreignKey:Uid;references:Uid;-:migration"`
}
