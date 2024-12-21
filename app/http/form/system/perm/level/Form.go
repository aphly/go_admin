package level

import "github.com/go-playground/validator/v10"

type Form struct {
	Id     uint   `json:"id" `
	Title  string `json:"title" binding:"required,max=32"`
	Pid    *uint  `json:"pid" binding:"required"`
	Sort   *int   `json:"sort" `
	Status *int8  `json:"status" binding:"required"`
}

func (this *Form) GetError(err validator.ValidationErrors) string {
	str := "格式错误"
	for _, v := range err {
		if v.Field() == "Title" {
			switch v.Tag() {
			case "required":
				str = "请输入名称"
			case "max":
				str = "最大长度为32位"
			}
		} else {
			str = v.Field() + " " + v.Tag() + "格式错误"
		}
	}
	return str
}
