package manager

import (
	"github.com/go-playground/validator/v10"
)

type Role struct {
	CheckedKeys   []uint `json:"checked_keys" binding:"required"`
	HalfcheckKeys []uint `json:"halfcheck_keys"`
}

func (this *Role) GetError(err validator.ValidationErrors) string {
	str := "格式错误"
	for _, v := range err {
		if v.Field() == "CheckedKeys" {
			switch v.Tag() {
			case "required":
				str = "请输入名称"
			}
		} else {
			str = v.Field() + " " + v.Tag() + " 格式错误"
		}
	}
	return str
}
