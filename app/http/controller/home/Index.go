package home

import (
	"github.com/gin-gonic/gin"
	"go_admin/app/res"
)

func Index(c *gin.Context) {
	res.Json(c, res.Code(1), res.Msg("ok"))
	return
}

func Test(c *gin.Context) {
	res.Json(c, res.Code(1), res.Msg("ok"))
	return
}
