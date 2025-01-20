package account

import "github.com/go-playground/validator/v10"

type LoginForm struct {
	IdType   string `json:"id_type" binding:"required,max=16"`
	Id       string `json:"id" binding:"required,len=11,numeric"`
	Password string `json:"password" binding:"required,min=6,max=32"`
}

func (this *LoginForm) GetError(err validator.ValidationErrors) string {
	str := "校验格式错误"
	for _, v := range err {
		if v.Field() == "Password" {
			switch v.Tag() {
			case "required":
				str = "请输入密码"
			case "min":
				str = "密码最小长度为6位"
			case "max":
				str = "密码最大长度为32位"
			}
		} else if v.Field() == "Id" {
			switch v.Tag() {
			case "required":
				str = "请输入手机号码"
			case "len":
				str = "手机号码必须11位"
			case "numeric":
				str = "手机号码必须数字"
			}
		} else {
			return v.Field() + " " + v.Tag() + "格式错误"
		}
	}
	return str
}
