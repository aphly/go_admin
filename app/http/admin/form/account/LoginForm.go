package account

import "github.com/go-playground/validator/v10"

type LoginForm struct {
	Username string `json:"username" binding:"required,max=32"`
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
		} else if v.Field() == "Username" {
			switch v.Tag() {
			case "required":
				str = "请输入账号"
			case "max":
				str = "最大长度为32位"
			}
		} else {
			//str = v.Field() + " " + v.Tag() + "格式错误"
		}
	}
	return str
}
