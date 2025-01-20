package upload

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/core"
	"go_admin/app/helper/sliceFunc"
	"mime/multipart"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type File struct {
	SizeConf int64 //单位兆
	TypeConf []any
}

func NewFile() *File {
	return &File{
		SizeConf: 1,
		TypeConf: []any{
			"image/jpeg",
			"image/png",
		},
	}
}

func (this *File) TypeLimit(fileTypes []string) bool {
	for _, i := range fileTypes {
		if !sliceFunc.InArray(i, this.TypeConf) {
			return false
		}
	}
	return true
}

func (this *File) SizeLimit(fileSize int64) bool {
	if fileSize <= this.SizeConf*1024*1024 {
		return true
	}
	return false
}

func (this *File) UidPath(uid core.Uint) string {
	str := fmt.Sprintf("%0*d", 20, uid)
	splitted := make([]string, len(str)/2)
	for i := 0; i < len(str); i += 2 {
		splitted[i/2] = str[i : i+2]
	}
	return strings.Join(splitted, "/")
}

func (this *File) LocalSave(c *gin.Context, file *multipart.FileHeader, uploadPath string) error {
	dir := filepath.Dir(uploadPath)
	err := os.MkdirAll("./"+dir, 0755)
	if err != nil {
		return errors.New("文件夹错误")
	}
	b := this.SizeLimit(file.Size)
	if b == false {
		return errors.New("大小超过" + strconv.FormatInt(this.SizeConf, 10) + "M")
	}
	stringArr := file.Header["Content-Type"]
	b = this.TypeLimit(stringArr)
	if b == false {
		return errors.New("格式只支持" + sliceFunc.Join(this.TypeConf, "、"))
	}
	err = c.SaveUploadedFile(file, uploadPath)
	if err != nil {
		return errors.New("上传错误")
	}
	return nil
}

func Url(path string, remote int8) string {
	if path == "" {
		return ""
	}
	if remote == 1 {
		if app.Config.Oss.IsCname == "true" {
			return app.Config.Oss.Url + path
		} else {
			return app.Config.Oss.Endpoint + path
		}
	} else if remote == 0 {
		if path == "" {
			return ""
		} else {
			return app.Config.Http.Host + path
		}
	} else {
		return ""
	}
}

func Del(path string, remote int8) error {
	if remote == 1 {
		if app.Config.Oss.IsCname == "true" {
			return nil
		} else {
			return nil
		}
	} else if remote == 0 {
		u, _ := url.Parse(path)
		if err := os.Remove("." + u.Path); err != nil {
			return err
		}
		return nil
	} else {
		return nil
	}
}
