package gorm

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DataPerm(c *gin.Context, uid any) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		level_ids, _ := c.Get("level_ids")
		LevelIds := level_ids.([]uint)
		if len(LevelIds) > 0 {
			return db.Where("level_id in ?", LevelIds)
		} else {
			return db.Where("uid = ?", uid)
		}
	}
}
