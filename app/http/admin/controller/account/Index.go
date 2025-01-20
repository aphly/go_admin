package account

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/core/crypt"
	"go_admin/app/http/admin/model"
	"go_admin/app/res"
)

type indexForm struct {
	Nickname string `json:"nickname" binding:"required,max=32"`
	Phone    string `json:"phone" binding:"max=32"`
	Password string `json:"password"`
}

func Index(c *gin.Context) {
	form := indexForm{}
	err := c.ShouldBind(&form)
	if err != nil {
		if jsonErr, ok := err.(*json.UnmarshalTypeError); ok {
			msg := fmt.Sprintf("%v 必须为 %v", jsonErr.Field, jsonErr.Type)
			res.Json(c, res.Code(1), res.Msg(msg))
			return
		} else {
			res.Json(c, res.Code(3), res.Msg(err.Error()))
			return
		}
	}
	uid, _ := c.Get("uid")
	if form.Password != "" {
		form.Password = crypt.ShaEn(form.Password)
	}
	err = app.Db().Model(&model.AdminManager{}).Where("uid=?", uid).Updates(form).Error
	if err != nil {
		res.Json(c, res.Code(11), res.Msg("更新错误"))
		return
	}
	var manager_info model.AdminManager
	err = app.Db().Where("uid=?", uid).Take(&manager_info).Error
	if err != nil {
		res.Json(c, res.Code(12), res.Msg("错误"))
		return
	}
	res.Json(c, res.Msg("更新成功"), res.Data(gin.H{
		"manager_info": manager_info,
	}))
	return
}
