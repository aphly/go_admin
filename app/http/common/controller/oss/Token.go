package oss

import (
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/http/common/service/upload"
	"go_admin/app/res"
)

func Token(c *gin.Context) {
	token, err := upload.OssGetPolicyToken(upload.CallbackParam{
		CallbackUrl:      app.Config.Http.Host + "/v1/oss/upload",
		CallbackBody:     "filename=${object}&size=${size}&mimeType=${mimeType}",
		CallbackBodyType: "application/json",
	})
	if err != nil {
		res.Json(c, res.Code(1), res.Msg(err.Error()))
		return
	}
	res.Json(c, res.Data(gin.H{
		"oss_token": token,
	}))
	return
}
