package router

import (
	"github.com/gin-gonic/gin"
	admin_router "go_admin/app/http/admin/router"
	blog_router "go_admin/app/http/blog/router"
	"go_admin/app/http/common/middleware"
)

func Reg() *gin.Engine {
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	r := gin.New()
	r.MaxMultipartMemory = 32 << 20
	r.Use(middleware.CrosHandler())
	r.Use(middleware.PanicHandler())
	r.Use(middleware.LogHandler())
	admin_router.Admin(r)
	v1(r)
	blog_router.Blog(r)
	return r
}
