package model

type AdminLevelPath struct {
	PathId  uint `gorm:"primaryKey;autoIncrement:false" json:"path_id"`
	LevelId uint `gorm:"primaryKey;autoIncrement:false" json:"level_id"`
	Level   int  `gorm:"default:0" json:"level"`
}
