package role

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go_admin/app"
	"go_admin/app/core"
	"go_admin/app/http/form/system/perm/role"
	"go_admin/app/http/model"
	"go_admin/app/res"
	"gorm.io/gorm"
	"strconv"
)

func Api(c *gin.Context) {
	role_id := c.DefaultQuery("role_id", "")
	if role_id == "" {
		res.Json(c, res.Code(1), res.Msg("role_id 为空"))
		return
	}
	if c.Request.Method == "POST" {
		uid, _ := c.Get("uid")
		uidD := uid.(core.Int64)
		form := role.Menu{}
		err := c.ShouldBind(&form)
		if err != nil {
			if jsonErr, ok := err.(*json.UnmarshalTypeError); ok {
				msg := fmt.Sprintf("%v 必须为 %v", jsonErr.Field, jsonErr.Type)
				res.Json(c, res.Code(2), res.Msg(msg))
				return
			} else if validatorErr, ok1 := err.(validator.ValidationErrors); ok1 {
				res.Json(c, res.Code(3), res.Msg(form.GetError(validatorErr)))
				return
			} else {
				res.Json(c, res.Code(4), res.Msg(err.Error()))
				return
			}
		}
		err = app.Db().Transaction(func(tx *gorm.DB) error {
			role_id64, err := strconv.ParseUint(role_id, 10, 64)
			if err != nil {
				return err
			}
			if err = app.Db().Where("role_id=?", role_id).Delete(&model.AdminRoleApi{}).Error; err != nil {
				return err
			}
			var adminRoleMenu []model.AdminRoleApi
			for _, v := range form.CheckedKeys {
				adminRoleMenu = append(adminRoleMenu, model.AdminRoleApi{
					RoleId: uint(role_id64),
					ApiId:  v,
					IsHalf: 0,
					Uid:    uidD,
				})
			}
			for _, v := range form.HalfcheckKeys {
				adminRoleMenu = append(adminRoleMenu, model.AdminRoleApi{
					RoleId: uint(role_id64),
					ApiId:  v,
					IsHalf: 1,
					Uid:    uidD,
				})
			}
			if len(adminRoleMenu) > 0 {
				if err = tx.Create(&adminRoleMenu).Error; err != nil {
					return err
				}
			}
			return nil
		})
		if err != nil {
			res.Json(c, res.Code(11), res.Msg("保存失败"))
			return
		}
		res.Json(c, res.Msg("保存成功"))
		return
	} else {
		var adminRoleApi []model.AdminRoleApi
		app.Db().InnerJoins("Api", app.Db().Where(&model.AdminApi{Status: 1})).
			Where("role_id=? and is_half=0", role_id).Find(&adminRoleApi)
		res.Json(c, res.Data(gin.H{
			"role_api": adminRoleApi,
		}))
		return
	}

}