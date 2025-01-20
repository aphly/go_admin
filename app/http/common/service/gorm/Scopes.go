package gorm

import (
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/http/admin/model"
	"gorm.io/gorm"
)

func HaveDataPerm(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		roleId, _ := c.Get("role_id")
		uid, _ := c.Get("uid")
		LevelIds := HaveLevelIds(roleId.(string))
		if len(LevelIds) > 0 {
			return db.Where("level_id in ?", LevelIds)
		} else {
			return db.Where("uid = ?", uid)
		}
	}
}

func BelongDataPerm(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		roleId, _ := c.Get("role_id")
		var LevelIds []uint
		if roleId != nil {
			LevelIds = BelongLevelIds(roleId.(string))
		} else {
			//uid, _ := c.Get("uid")
			//LevelIds = BelongLevelIdsByUid(uid)
			return db.Where("level_id=0")
		}
		if len(LevelIds) > 0 {
			return db.Where("level_id in ?", LevelIds)
		} else {
			return db.Where("level_id=0")
		}
	}
}

func PreloadManager(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload("Manager", func(db *gorm.DB) *gorm.DB {
			return db.Select("uid", "nickname", "username", "level_id", "avatar", "status")
		})
	}
}

func PreloadLevel(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload("Level", func(db *gorm.DB) *gorm.DB {
			return db.Select("uid", "title", "status")
		})
	}
}

func InnerLevel(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.InnerJoins("Level", app.Db().Where(&model.AdminLevel{Status: 1}).Select("uid", "title", "status"))
	}
}
