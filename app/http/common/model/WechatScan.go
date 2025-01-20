package model

type WechatScan struct {
	RandStr string `gorm:"type:char(32);primarykey" json:"rand_str"`
	Openid  string `gorm:"size:32" json:"openid"`
	Status  int8   `gorm:"default:0" json:"status"`
}
