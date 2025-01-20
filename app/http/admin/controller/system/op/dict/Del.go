package dict

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/http/admin/model"
	"go_admin/app/res"
	"gorm.io/gorm"
)

type del struct {
	Ids []int `json:"ids" binding:"required"`
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
	err = app.Db().Transaction(func(tx *gorm.DB) error {
		err := tx.Where("id in ?", form.Ids).Delete(&model.AdminDict{}).Error
		if err != nil {
			return err
		}
		err = tx.Where("dict_id in ?", form.Ids).Delete(&model.AdminDictValue{}).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		res.Json(c, res.Code(11), res.Msg("删除失败"))
		return
	}
	res.Json(c, res.Msg("删除成功"))
	return
}
