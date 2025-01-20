package category

import (
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/http/blog/model"
	"go_admin/app/res"
)

func FRebuild(pid uint) {
	var categoryData []model.BlogCategory
	app.Db("blog").Where("pid=?", pid).Find(&categoryData)
	for _, v := range categoryData {
		app.Db("blog").Where("category_id=?", v.Id).Delete(&model.BlogCategoryPath{})
		level := 0
		var categoryPathData []model.BlogCategoryPath
		app.Db("blog").Where("category_id=?", v.Pid).Order("level asc").Find(&categoryPathData)
		var data []*model.BlogCategoryPath
		for _, v1 := range categoryPathData {
			data = append(data, &model.BlogCategoryPath{
				CategoryId: v.Id,
				PathId:     v1.PathId,
				Level:      level,
			})
			level++
		}
		data = append(data, &model.BlogCategoryPath{
			CategoryId: v.Id,
			PathId:     v.Id,
			Level:      level,
		})
		app.Db("blog").Create(data)
		FRebuild(v.Id)
	}
}

func Rebuild(c *gin.Context) {
	app.Db("blog").Where("1 = 1").Delete(&model.BlogCategoryPath{})
	FRebuild(0)
	res.Json(c, res.Msg("重建成功"))
	return
}
