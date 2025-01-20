package model

import (
	"go_admin/app/core"
)

type BlogCategory struct {
	core.ModelId
	Uid             core.Uint `gorm:"index" json:"uid"`
	Title           string    `gorm:"size:64" json:"title"`
	Name            string    `gorm:"size:32" json:"name"`
	Path            string    `gorm:"size:255" json:"path"`
	Pid             uint      `gorm:"default:0" json:"pid"`
	Sort            int       `gorm:"default:0" json:"sort"`
	Status          int8      `gorm:"default:0" json:"status"`
	Type            int8      `gorm:"default:0" json:"type"`
	MetaTitle       string    `gorm:"size:32" json:"meta_title"`
	MetaDescription string    `gorm:"size:255" json:"meta_description"`
}
