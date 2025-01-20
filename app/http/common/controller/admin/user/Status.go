package user

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/http/common/model"
	"go_admin/app/res"
)

type Form struct {
	Status *int   `json:"status" binding:"required"`
	Uid    string `json:"uid" binding:"required"`
}

func Status(c *gin.Context) {
	form := Form{}
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

	status := 1
	if *(form.Status) == 1 {
		status = 1
	} else {
		status = 0
	}
	err = app.Db("common").Model(&model.User{}).Where("uid=?", form.Uid).Update("status", status).Error

	if err != nil {
		res.Json(c, res.Code(11), res.Msg("操作失败"))
		return
	}
	res.Json(c, res.Msg("操作成功"))
	return
}
