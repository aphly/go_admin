package manager

import (
	"github.com/go-playground/validator/v10"
)

type Form struct {
	Uid      string `json:"uid" `
	Username string `json:"username" binding:"required,max=32"`
	Password string `json:"password"`
	Nickname string `json:"nickname" binding:"required,max=32"`
	Phone    string `json:"phone" binding:"max=32"`
	Note     string `json:"note" binding:"max=255"`
	//LevelId  uint   `json:"level_id" binding:"required"`
}

func (this *Form) GetError(err validator.ValidationErrors) string {
	str := ""
	for _, v := range err {
		if v.Field() == "Nickname" {
			switch v.Tag() {
			case "required":
				str = "请输入昵称"
			case "max":
				str = "密码最大长度为32位"
			}
		} else if v.Field() == "Phone" {
			switch v.Tag() {
			case "max":
				str = "最大长度为32位"
			}
		} else if v.Field() == "Note" {
			switch v.Tag() {
			case "max":
				str = "最大长度为255位"
			}
		} else {
			str = v.Field() + " " + v.Tag() + "格式错误"
		}
	}
	return str
}
