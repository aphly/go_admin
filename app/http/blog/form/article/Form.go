package article

import "github.com/go-playground/validator/v10"

type Form struct {
	Id         uint   `json:"id" `
	Title      string `json:"title" binding:"required,max=64"`
	Content    string `json:"content"`
	Status     *int8  `json:"status" binding:"required"`
	Sort       *int   `json:"sort"`
	Viewed     *int   `json:"viewed"`
	CategoryId uint   `json:"category_id" binding:"required"`
}

func (this *Form) GetError(err validator.ValidationErrors) string {
	str := "格式错误"
	for _, v := range err {
		if v.Field() == "Title" {
			switch v.Tag() {
			case "required":
				str = "请输入名称"
			case "max":
				str = "最大长度为64位"
			}
		} else {
			str = v.Field() + " " + v.Tag() + "格式错误"
		}
	}
	return str
}
