package model

import (
	"go_admin/app/core"
	"gorm.io/gorm"
)

type AdminOperation struct {
	Id          uint           `gorm:"primaryKey"  json:"id"`
	Uid         core.Uint      `gorm:"index" json:"uid"`
	Url         string         `gorm:"size:255" json:"url"`
	Method      string         `gorm:"size:16" json:"method"`
	RequestData string         `gorm:"type:text" json:"request_data"`
	Ip          string         `gorm:"size:64" json:"ip"`
	CreatedAt   int64          `gorm:"autoUpdateTime" json:"created_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Manager     AdminManager   `gorm:"foreignKey:Uid;references:Uid;-:migration"`
}
