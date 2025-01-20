package dict

import "github.com/go-playground/validator/v10"

type ValueForm struct {
	Id     uint   `json:"id"`
	Title  string `json:"title" binding:"required,max=32"`
	Value  string `json:"value" binding:"required,max=255"`
	Sort   *int   `json:"sort" `
	DictId uint   `json:"dict_id" binding:"required"`
}

func (this *ValueForm) GetError(err validator.ValidationErrors) string {
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
