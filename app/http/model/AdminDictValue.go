package model

import (
	"go_admin/app/core"
)

type AdminDictValue struct {
	core.ModelId
	DictId uint      `gorm:"index;default:0" json:"dict_id"`
	Title  string    `gorm:"size:32" json:"title"`
	Value  string    `gorm:"size:32" json:"value"`
	Sort   int       `gorm:"default:0" json:"sort"`
	Dict   AdminDict `gorm:"foreignKey:DictId;-:migration"`
}
