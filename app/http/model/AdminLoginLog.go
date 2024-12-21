package model

import (
	"go_admin/app/core"
)

type AdminLoginLog struct {
	core.ModelId
	Ip             string `gorm:"index;size:64" json:"ip,omitempty"`
	Input          string `gorm:"size:255" json:"input,omitempty"`
	UserAgent      string `gorm:"size:255" json:"user_agent,omitempty"`
	AcceptLanguage string `gorm:"size:255" json:"accept_language,omitempty"`
}
