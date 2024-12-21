package model

import (
	"go_admin/app/core"
	"gorm.io/gorm"
)

type AdminOperation struct {
	ID        uint           `gorm:"primaryKey"  json:"id"`
	Uid       core.Int64     `gorm:"index" json:"uid,omitempty"`
	Path      string         `gorm:"size:255" json:"path,omitempty"`
	Request   string         `gorm:"type:text" json:"request,omitempty"`
	Res       string         `gorm:"type:text" json:"res,omitempty"`
	Ip        string         `gorm:"size:64" json:"ip,omitempty"`
	CreatedAt int64          `gorm:"autoUpdateTime" json:"created_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
