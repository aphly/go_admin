package menu

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go_admin/app"
	"go_admin/app/core"
	"go_admin/app/http/admin/form/system/perm/menu"
	"go_admin/app/http/admin/model"
	"go_admin/app/res"
)

func Add(c *gin.Context) {
	uid, _ := c.Get("uid")
	uid_d := uid.(core.Uint)
	form := menu.Form{}
	err := c.ShouldBind(&form)
	if err != nil {
		if jsonErr, ok := err.(*json.UnmarshalTypeError); ok {
			msg := fmt.Sprintf("%v 必须为 %v", jsonErr.Field, jsonErr.Type)
			res.Json(c, res.Code(1), res.Msg(msg))
			return
		} else if validatorErr, ok1 := err.(validator.ValidationErrors); ok1 {
			res.Json(c, res.Code(2), res.Msg(form.GetError(validatorErr)))
			return
		} else {
			res.Json(c, res.Code(3), res.Msg(err.Error()))
			return
		}
	}
	addMenu := model.AdminMenu{
		Uid:       uid_d,
		Title:     form.Title,
		Name:      form.Name,
		Path:      form.Path,
		Pid:       *form.Pid,
		Sort:      *form.Sort,
		Status:    *form.Status,
		Type:      form.Type,
		Icon:      form.Icon,
		Component: form.Component,
	}
	err = app.Db().Create(&addMenu).Error
	if err != nil {
		res.Json(c, res.Code(11), res.Msg("保存失败"))
		return
	}
	var menus []model.AdminMenu
	app.Db().Find(&menus)
	halfIds := getAllPids(menus, addMenu.Id)
	if len(halfIds) > 0 {
		app.Db().Model(&model.AdminRoleMenu{}).Where("menu_id in ?", halfIds).Update("is_half", 1)
	}
	res.Json(c, res.Msg("保存成功"))
	return
}

func getAllPids(items []model.AdminMenu, id uint) []uint {
	var parentIDs []uint
	var findParent func(uint)
	findParent = func(id uint) {
		for _, item := range items {
			if item.Id == id && item.Pid != 0 {
				parentIDs = append(parentIDs, item.Pid)
				findParent(item.Pid)
				break
			}
		}
	}
	findParent(id)
	return parentIDs
}
