package router

import (
	"github.com/gin-gonic/gin"
	"go_admin/app/http/common/controller/admin/user"
	"go_admin/app/http/common/controller/front/account"
	"go_admin/app/http/common/controller/front/account/wechat"
	"go_admin/app/http/common/controller/oss"
	"go_admin/app/http/common/middleware"
)

func v1(router *gin.Engine) {
	ipLimit := router.Group("")
	ipLimit.Use(middleware.IpLimitHandler())
	{
		ipLimit.GET("/oss/token", oss.Token)
		ipLimit.POST("/oss/upload", oss.Upload)

		ipLimit.POST("/account/login", account.Login)
		ipLimit.POST("/account/register", account.Register)
		ipLimit.POST("/account/refresh_token", account.RefreshToken)
		ipLimit.POST("/account/forget", account.Forget)

	}

	front := router.Group("")
	{
		front.GET("/wechat/scan", wechat.Scan)
		front.GET("/wechat/openid", wechat.Openid)
		front.GET("/wechat/login", wechat.Login)
	}

	admin := router.Group("/common")
	{
		admin.Use(middleware.ManagerAuthHandler())
		{
			rbac := admin.Group("")
			rbac.Use(middleware.RbacHandler())
			{
				rbac.GET("/user/index", user.Index)
				rbac.POST("/user/add", user.Add)
				rbac.POST("/user/edit", user.Edit)
				rbac.POST("/user/del", user.Del)
				rbac.POST("/user/status", user.Status)
				rbac.POST("/user/avatar", user.Avatar)

			}
		}

	}
}
