package router

import (
	"github.com/gin-gonic/gin"
	"go_admin/app/http/controller/account"
	"go_admin/app/http/controller/home"
	"go_admin/app/http/controller/system/op/dict"
	"go_admin/app/http/controller/system/op/login_log"
	"go_admin/app/http/controller/system/op/operation"
	"go_admin/app/http/controller/system/perm/api"
	"go_admin/app/http/controller/system/perm/level"
	"go_admin/app/http/controller/system/perm/manager"
	"go_admin/app/http/controller/system/perm/menu"
	"go_admin/app/http/controller/system/perm/role"
	"go_admin/app/http/middleware"
)

func admin(router *gin.Engine) {
	router.Static("/public", "./public")
	admin := router.Group("/admin")
	{
		admin.POST("/login", account.Login)
		admin.GET("/refresh_token", account.RefreshToken)

		auth := admin.Group("")
		auth.Use(middleware.ManagerAuthHandler())
		{
			auth.GET("/account/role_menu", account.RoleMenu)
			auth.POST("/account/avatar", account.Avatar)
			auth.POST("/account/index", account.Index)
			rbac := auth.Group("")
			rbac.Use(middleware.RbacHandler())
			{
				rbac.GET("/home/index", home.Index)

				systemPerm := rbac.Group("/system/perm")
				{
					systemPerm.GET("/level/all", level.All)
					systemPerm.POST("/level/add", level.Add)
					systemPerm.POST("/level/edit", level.Edit)
					systemPerm.POST("/level/del", level.Del)
					systemPerm.POST("/level/rebuild", level.Rebuild)

					systemPerm.GET("/manager/index", manager.Index)
					systemPerm.POST("/manager/add", manager.Add)
					systemPerm.POST("/manager/edit", manager.Edit)
					systemPerm.POST("/manager/del", manager.Del)
					systemPerm.GET("/manager/info", manager.Info)
					systemPerm.POST("/manager/status", manager.Status)
					systemPerm.Match([]string{"GET", "POST"}, "/manager/role", manager.Role)

					systemPerm.GET("/role/index", role.Index)
					systemPerm.POST("/role/add", role.Add)
					systemPerm.POST("/role/edit", role.Edit)
					systemPerm.POST("/role/del", role.Del)
					systemPerm.GET("/role/info", role.Info)
					systemPerm.GET("/role/all", role.All)
					systemPerm.Match([]string{"GET", "POST"}, "/role/menu", role.Menu)
					systemPerm.Match([]string{"GET", "POST"}, "/role/api", role.Api)

					systemPerm.GET("/menu/all", menu.All)
					systemPerm.POST("/menu/add", menu.Add)
					systemPerm.POST("/menu/edit", menu.Edit)
					systemPerm.POST("/menu/del", menu.Del)

					systemPerm.GET("/api/all", api.All)
					systemPerm.POST("/api/add", api.Add)
					systemPerm.POST("/api/edit", api.Edit)
					systemPerm.POST("/api/del", api.Del)
				}

				systemOp := rbac.Group("/system/op")
				{
					systemOp.GET("/operation/index", operation.Index)
					//systemPerm.GET("/operation/info", operation.Info)
					systemOp.POST("/operation/del", operation.Del)

					systemOp.GET("/dict/index", dict.Index)
					systemOp.POST("/dict/del", dict.Del)
					systemOp.POST("/dict/add", dict.Add)
					systemOp.POST("/dict/edit", dict.Edit)
					systemOp.GET("/dict/value_index", dict.ValueIndex)
					systemOp.POST("/dict/value_del", dict.ValueDel)
					systemOp.POST("/dict/value_add", dict.ValueAdd)
					systemOp.POST("/dict/value_edit", dict.ValueEdit)

					systemOp.GET("/login_log/index", login_log.Index)
					systemOp.POST("/login_log/del", login_log.Del)
				}

			}
		}

		ipLimit := admin.Group("")
		ipLimit.Use(middleware.IpLimitHandler())
		{

		}

	}
}
