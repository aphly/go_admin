package router

import (
	"github.com/gin-gonic/gin"
	"go_admin/app/http/middleware"
)

func Reg() *gin.Engine {
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	r := gin.New()
	r.Use(middleware.CrosHandler())
	r.Use(middleware.PanicHandler())
	r.Use(middleware.LogHandler())
	admin(r)
	v1(r)
	return r
}
