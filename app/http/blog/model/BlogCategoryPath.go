package model

type BlogCategoryPath struct {
	CategoryId uint `gorm:"primaryKey;autoIncrement:false" json:"category_id"`
	PathId     uint `gorm:"primaryKey;autoIncrement:false" json:"path_id"`
	Level      int  `gorm:"default:0" json:"level"`
}
