package router

import (
	"github.com/gin-gonic/gin"
	"go_admin/app/http/controller/account"
	"go_admin/app/http/controller/home"
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
				rbac.GET("/home/test", home.Test)

				systemPerm := auth.Group("/system/perm")
				{
					systemPerm.GET("/level/all", level.All)
					systemPerm.POST("/level/add", level.Add)
					systemPerm.POST("/level/edit", level.Edit)
					systemPerm.GET("/level/del", level.Del)

					systemPerm.GET("/menu/all", menu.All)
					systemPerm.POST("/menu/add", menu.Add)
					systemPerm.POST("/menu/edit", menu.Edit)
					systemPerm.GET("/menu/del", menu.Del)

					systemPerm.GET("/role/index", role.Index)
					systemPerm.POST("/role/add", role.Add)
					systemPerm.POST("/role/edit", role.Edit)
					systemPerm.GET("/role/del", role.Del)
					systemPerm.GET("/role/info", role.Info)
					systemPerm.GET("/role/all", role.All)
					systemPerm.Match([]string{"GET", "POST"}, "/role/menu", role.Menu)
					systemPerm.Match([]string{"GET", "POST"}, "/role/api", role.Api)

					systemPerm.GET("/manager/index", manager.Index)
					systemPerm.POST("/manager/add", manager.Add)
					systemPerm.POST("/manager/edit", manager.Edit)
					systemPerm.GET("/manager/del", manager.Del)
					systemPerm.GET("/manager/info", manager.Info)

					systemPerm.GET("/manager/blacklist", manager.Blacklist)
					systemPerm.Match([]string{"GET", "POST"}, "/manager/role", manager.Role)

					systemPerm.GET("/api/all", api.All)
					systemPerm.POST("/api/add", api.Add)
					systemPerm.POST("/api/edit", api.Edit)
					systemPerm.GET("/api/del", api.Del)
				}

			}
		}

		ipLimit := admin.Group("")
		ipLimit.Use(middleware.IpLimitHandler())
		{

		}

	}
}
