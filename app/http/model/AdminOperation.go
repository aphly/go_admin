package model

import (
	"go_admin/app/core"
	"gorm.io/gorm"
)

type AdminOperation struct {
	ID          uint           `gorm:"primaryKey"  json:"id"`
	Uid         core.Int64     `gorm:"index" json:"uid,omitempty"`
	Url         string         `gorm:"size:255" json:"url,omitempty"`
	Method      string         `gorm:"size:16" json:"method,omitempty"`
	RequestData string         `gorm:"type:text" json:"request_data,omitempty"`
	Ip          string         `gorm:"size:64" json:"ip,omitempty"`
	CreatedAt   int64          `gorm:"autoUpdateTime" json:"created_at,omitempty"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Manager     AdminManager   `gorm:"foreignKey:Uid;references:Uid;-:migration"`
}
