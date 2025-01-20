package router

import (
	"github.com/gin-gonic/gin"
	"go_admin/app/http/blog/controller/admin/article"
	"go_admin/app/http/blog/controller/admin/category"
	front_article "go_admin/app/http/blog/controller/front/article"
	"go_admin/app/http/common/middleware"
)

func Blog(router *gin.Engine) {
	admin := router.Group("/blog")
	{
		admin.Use(middleware.ManagerAuthHandler())
		{
			rbac := admin.Group("")
			rbac.Use(middleware.RbacHandler())
			{
				rbac.GET("/article/index", article.Index)
				rbac.POST("/article/add", article.Add)
				rbac.POST("/article/edit", article.Edit)
				rbac.POST("/article/del", article.Del)
				rbac.POST("/article/status", article.Status)
				rbac.POST("/article/file", article.File)

				rbac.GET("/category/all", category.All)
				rbac.POST("/category/add", category.Add)
				rbac.POST("/category/edit", category.Edit)
				rbac.POST("/category/del", category.Del)
				rbac.POST("/category/rebuild", category.Rebuild)
			}
		}

		ipLimit := admin.Group("")
		ipLimit.Use(middleware.IpLimitHandler())
		{
		}

	}

	fornt := router.Group("")
	{
		fornt.GET("/article/index", front_article.Index)
		fornt.GET("/article/info", front_article.Info)
		fornt.GET("/article/category", front_article.Category)
	}

}
