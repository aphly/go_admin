package manager

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go_admin/app"
	"go_admin/app/core"
	"go_admin/app/http/admin/form/system/perm/manager"
	"go_admin/app/http/admin/model"
	"go_admin/app/res"
	"gorm.io/gorm"
	"strconv"
)

func Role(c *gin.Context) {
	manager_uid := c.DefaultQuery("manager_uid", "")
	if manager_uid == "" {
		res.Json(c, res.Code(1), res.Msg("manager_uid 为空"))
		return
	}
	if c.Request.Method == "POST" {
		uid, _ := c.Get("uid")
		uidD := uid.(core.Uint)
		form := manager.Role{}
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
			manager_uid64, err := strconv.ParseUint(manager_uid, 10, 64)
			if err != nil {
				return err
			}
			if err = tx.Where("manager_uid=?", manager_uid64).Delete(&model.AdminManagerRole{}).Error; err != nil {
				return err
			}
			var adminRoleMenu []model.AdminManagerRole
			for _, v := range form.CheckedKeys {
				adminRoleMenu = append(adminRoleMenu, model.AdminManagerRole{
					ManagerUid: core.Uint(manager_uid64),
					RoleId:     v,
					IsHalf:     0,
					Uid:        uidD,
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
			res.Json(c, res.Code(11), res.Msg("保存失败:"+err.Error()))
			return
		}
		res.Json(c, res.Msg("保存成功"))
		return
	} else {
		var adminManagerRole []model.AdminManagerRole
		app.Db().InnerJoins("Role", app.Db().Where(&model.AdminRole{Status: 1})).
			Where("manager_uid=? and is_half=0", manager_uid).Find(&adminManagerRole)
		res.Json(c, res.Data(gin.H{
			"manager_role": adminManagerRole,
		}))
		return
	}

}
