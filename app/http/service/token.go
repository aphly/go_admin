package service

import (
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"go_admin/app/core"
	"time"
)

func AddTokenToBlacklist(c *gin.Context, blacklistKey string, Uid core.Int64) error {
	//blacklistKey := "jwt:managerBlacklist"
	_, err := app.RedisSingle().SAdd(c, blacklistKey, int64(Uid)).Result()
	if err != nil {
		return err
	}
	_, err = app.RedisSingle().Expire(c, blacklistKey, time.Hour*2).Result()
	if err != nil {
		return err
	}
	return nil
}

func IsTokenBlacklisted(c *gin.Context, blacklistKey string, Uid core.Int64) (bool, error) {
	//blacklistKey := "jwt:managerBlacklist"
	_, err := app.RedisSingle().Exists(c, blacklistKey).Result()
	if err != nil {
		return false, err
	}
	isBlacklisted, err := app.RedisSingle().SIsMember(c, blacklistKey, int64(Uid)).Result()
	if err != nil {
		return false, err
	}
	return isBlacklisted, nil
}
