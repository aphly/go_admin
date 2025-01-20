package manager

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/http/admin/model"
	"go_admin/app/http/common/service/redis"
	"go_admin/app/res"
	"gorm.io/gorm"
	"strconv"
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
	uid, _ := strconv.ParseUint(form.Uid, 10, 64)
	uidI := int64(uid)
	err = app.Db().Transaction(func(tx *gorm.DB) error {
		if *(form.Status) == 1 {
			err := tx.Model(&model.AdminManager{}).Where("uid=?", uidI).Update("status", 1).Error
			if err != nil {
				return err
			}
			err1 := redis.RemoveTokenToBlacklist(c, "jwt:managerBlacklist", uidI)
			if err1 != nil {
				return err1
			}
		} else {
			err := tx.Model(&model.AdminManager{}).Where("uid=?", form.Uid).Update("status", 0).Error
			if err != nil {
				return err
			}
			err1 := redis.AddTokenToBlacklist(c, "jwt:managerBlacklist", uidI)
			if err1 != nil {
				return err1
			}
		}
		return nil
	})
	if err != nil {
		res.Json(c, res.Code(11), res.Msg("操作失败"))
		return
	}
	res.Json(c, res.Msg("操作成功"))
	return
}
