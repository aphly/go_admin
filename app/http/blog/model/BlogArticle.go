package model

import (
	"go_admin/app/core"
)

type BlogArticle struct {
	core.ModelId
	Uid        core.Uint    `gorm:"index" json:"uid"`
	Title      string       `gorm:"size:64" json:"title"`
	Content    string       `gorm:"type:text" json:"content"`
	Viewed     uint         `gorm:"default:0" json:"viewed"`
	Sort       int          `gorm:"default:0;index:idx_s" json:"sort"`
	Status     int8         `gorm:"default:0;index:idx_s" json:"status"`
	CategoryId uint         `gorm:"default:0;index:idx_s" json:"category_id"`
	Category   BlogCategory `gorm:"foreignKey:CategoryId;-:migration"`
}
