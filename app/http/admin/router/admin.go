package router

import (
	"github.com/gin-gonic/gin"
	"go_admin/app/http/admin/controller/account"
	account_notice "go_admin/app/http/admin/controller/account/notice"
	"go_admin/app/http/admin/controller/home"
	"go_admin/app/http/admin/controller/system/op/dict"
	"go_admin/app/http/admin/controller/system/op/login_log"
	"go_admin/app/http/admin/controller/system/op/notice"
	"go_admin/app/http/admin/controller/system/op/operation"
	"go_admin/app/http/admin/controller/system/op/server"
	"go_admin/app/http/admin/controller/system/op/upload"
	"go_admin/app/http/admin/controller/system/perm/api"
	"go_admin/app/http/admin/controller/system/perm/level"
	"go_admin/app/http/admin/controller/system/perm/manager"
	"go_admin/app/http/admin/controller/system/perm/menu"
	"go_admin/app/http/admin/controller/system/perm/role"
	"go_admin/app/http/common/middleware"
)

func Admin(router *gin.Engine) {
	router.Static("/public", "./public")
	admin := router.Group("/admin")
	{
		admin.POST("/login", account.Login)
		admin.GET("/refresh_token", account.RefreshToken)

		auth := admin.Group("")
		auth.Use(middleware.ManagerAuthHandler())
		{
			accountGroup := auth.Group("/account")
			{
				accountGroup.GET("/role_menu", account.RoleMenu)
				accountGroup.GET("/manager_role", account.ManagerRole)
				accountGroup.POST("/avatar", account.Avatar)
				accountGroup.POST("/index", account.Index)
				accountGroup.GET("/notice/index", account_notice.Index)
			}

			auth.GET("/dict/info", dict.Info)

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
					systemOp.POST("/operation/del", operation.Del)

					systemOp.GET("/dict/index", dict.Index)
					systemOp.POST("/dict/del", dict.Del)
					systemOp.POST("/dict/add", dict.Add)
					systemOp.POST("/dict/edit", dict.Edit)
					systemOp.POST("/dict/status", dict.Status)
					systemOp.GET("/dict/value_index", dict.ValueIndex)
					systemOp.POST("/dict/value_del", dict.ValueDel)
					systemOp.POST("/dict/value_add", dict.ValueAdd)
					systemOp.POST("/dict/value_edit", dict.ValueEdit)

					systemOp.GET("/login_log/index", login_log.Index)
					systemOp.POST("/login_log/del", login_log.Del)

					systemOp.GET("/notice/index", notice.Index)
					systemOp.POST("/notice/add", notice.Add)
					systemOp.POST("/notice/edit", notice.Edit)
					systemOp.POST("/notice/del", notice.Del)
					systemOp.POST("/notice/status", notice.Status)

					systemOp.GET("/server/index", server.Index)

					systemOp.GET("/upload/index", upload.Index)
					systemOp.POST("/upload/add", upload.Add)
					systemOp.POST("/upload/del", upload.Del)
					systemOp.GET("/upload/info", upload.Info)
				}

			}
		}

	}
}
