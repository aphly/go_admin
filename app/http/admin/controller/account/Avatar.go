package account

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/http/admin/model"
	"go_admin/app/res"
)

type avatarForm struct {
	AvatarId string `json:"avatar_id" binding:"required"`
}

func Avatar(c *gin.Context) {
	form := avatarForm{}
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
	err = app.Db().Model(&model.AdminManager{}).Where("uid=?", uid).Update("avatar", form.AvatarId).Error
	if err != nil {
		res.Json(c, res.Code(11), res.Msg("更新错误"))
		return
	}
	res.Json(c, res.Msg("更新成功"))
	return
}
