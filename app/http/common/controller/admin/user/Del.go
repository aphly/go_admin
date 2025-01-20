package user

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/http/common/model"
	"go_admin/app/http/common/service/upload"
	"go_admin/app/res"
)

type del struct {
	Ids []string `json:"ids" binding:"required"`
}

func Del(c *gin.Context) {
	form := del{}
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
	for _, v := range form.Ids {
		var userModel model.User
		app.Db("common").Model(&model.User{}).Where("uid = ?", v).Take(&userModel)
		if userModel.Uid > 0 {
			err = app.Db("common").Model(&model.User{}).Where("uid = ?", v).Delete(&model.User{}).Error
			if err != nil {
				res.Json(c, res.Code(11), res.Msg("删除失败"))
				return
			}
			upload.Del(userModel.Avatar, userModel.Remote)
		}
	}

	res.Json(c, res.Msg("删除成功"))
	return
}
