package model

import (
	"go_admin/app/core"
)

type AdminDictValue struct {
	core.ModelId
	DictId uint      `gorm:"index;default:0" json:"dict_id,omitempty"`
	Title  string    `gorm:"size:32" json:"title,omitempty"`
	Value  string    `gorm:"size:255" json:"value,omitempty"`
	Sort   int       `gorm:"default:0" json:"sort,omitempty"`
	Dict   AdminDict `gorm:"foreignKey:DictId;-:migration"`
}
