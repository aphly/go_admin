package manager

import (
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/core"
	"go_admin/app/http/admin/model"
	"go_admin/app/res"
	"strconv"
)

type Search struct {
	status   int8
	username string
	uid      int64
	phone    string
}

func Index(c *gin.Context) {
	//uid, _ := c.Get("uid")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize := app.Config.PageSize
	offset := (page - 1) * pageSize

	var count int64
	db := app.Db().Model(&model.AdminManager{})

	if status := c.DefaultQuery("status", ""); status != "" {
		db.Where("status=?", status)
	}

	if username := c.DefaultQuery("username", ""); username != "" {
		db.Where("username like ?", username+"%")
	}

	if uid := c.DefaultQuery("uid", ""); uid != "" {
		db.Where("uid=?", uid)
	}

	if phone := c.DefaultQuery("phone", ""); phone != "" {
		db.Where("phone like ?", phone+"%")
	}

	err := db.Count(&count).Error
	if err != nil {
		res.Json(c, res.Code(11), res.Msg(err.Error()))
		return
	}
	var list []model.AdminManager
	managerRoleMap := make(map[core.Uint][]model.AdminManagerRole)
	if count > 0 {
		err = db.Order("uid desc").Offset(offset).Limit(pageSize).Find(&list).Error
		if err != nil {
			res.Json(c, res.Code(12), res.Msg(err.Error()))
			return
		}
		var managerUids []core.Uint
		for _, v := range list {
			managerUids = append(managerUids, v.Uid)
		}
		var managerRole []model.AdminManagerRole
		app.Db().InnerJoins("Role", app.Db().Where(&model.AdminRole{Status: 1})).
			Where("manager_uid in ? and is_half=0", managerUids).Find(&managerRole)
		for _, v := range managerRole {
			managerRoleMap[v.ManagerUid] = append(managerRoleMap[v.ManagerUid], v)
		}
	}

	res.Json(c, res.Data(gin.H{
		"list":         list,
		"count":        count,
		"manager_role": managerRoleMap,
	}))
	return
}
