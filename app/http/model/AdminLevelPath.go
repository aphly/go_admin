package model

type AdminLevelPath struct {
	PathId  uint `gorm:"primaryKey;autoIncrement:false" json:"path_id,omitempty"`
	LevelId uint `gorm:"primaryKey;autoIncrement:false" json:"level_id,omitempty"`
	Level   int  `gorm:"default:0" json:"level,omitempty"`
}
