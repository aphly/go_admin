package model

type AdminLevelPath struct {
	LevelId uint `gorm:"primaryKey;autoIncrement:false" json:"level_id"`
	PathId  uint `gorm:"primaryKey;autoIncrement:false" json:"path_id"`
	Level   int  `gorm:"default:0" json:"level"`
}
