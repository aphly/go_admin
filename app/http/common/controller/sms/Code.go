package sms

import (
	"github.com/gin-gonic/gin"
	"go_admin/app/res"
)

func Code(c *gin.Context) {

	res.Json(c, res.Msg("短信已发送"))
	return
}
