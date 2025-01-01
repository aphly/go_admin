package model

import (
	"go_admin/app/core"
)

type AdminLoginLog struct {
	core.ModelId
	Ip             string `gorm:"index;size:64" json:"ip"`
	Input          string `gorm:"size:255" json:"input"`
	UserAgent      string `gorm:"size:255" json:"user_agent"`
	AcceptLanguage string `gorm:"size:255" json:"accept_language"`
}
