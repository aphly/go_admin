package oss

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_admin/app/http/common/service/upload"
	"go_admin/app/res"
)

type Form struct {
	Filename string
	Size     string
	MimeType string
}

func Upload(c *gin.Context) {
	form := Form{}
	err := c.ShouldBind(&form)
	if err != nil {
		if jsonErr, ok := err.(*json.UnmarshalTypeError); ok {
			msg := fmt.Sprintf("%v 必须为 %v", jsonErr.Field, jsonErr.Type)
			res.Json(c, res.Code(1), res.Msg(msg))
			return
		} else {
			res.Json(c, res.Code(3), res.Msg(err.Error()))
			return
		}
	}

	bytePublicKey, err := upload.GetPublicKey(c)
	if err != nil {
		res.Json(c, res.Code(4), res.Msg(err.Error()))
		return
	}

	byteAuthorization, err := upload.GetAuthorization(c)
	if err != nil {
		res.Json(c, res.Code(5), res.Msg(err.Error()))
		return
	}

	byteMD5, err := upload.GetMD5FromNewAuthString(c)
	if err != nil {
		res.Json(c, res.Code(6), res.Msg(err.Error()))
		return
	}

	if upload.OssVerifySignature(bytePublicKey, byteMD5, byteAuthorization) {
		// do something you want accoding to callback_body ...

		res.Json(c, res.Data(gin.H{
			"Status": "ok",
			"asd":    "自行车自行车",
		}))
	} else {
		res.Json(c, res.Code(7), res.Msg("校验失败"))
	}
	return

}
