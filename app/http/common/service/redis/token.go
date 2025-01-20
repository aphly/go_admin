package redis

import (
	"github.com/gin-gonic/gin"
	"go_admin/app"
	"time"
)

func AddTokenToBlacklist(c *gin.Context, blacklistKey string, uid int64) error {
	//blacklistKey := "jwt:managerBlacklist"
	_, err := app.RedisSingle().SAdd(c, blacklistKey, uid).Result()
	if err != nil {
		return err
	}
	_, err = app.RedisSingle().Expire(c, blacklistKey, time.Hour*2).Result()
	if err != nil {
		return err
	}
	return nil
}

func RemoveTokenToBlacklist(c *gin.Context, blacklistKey string, uid int64) error {
	//blacklistKey := "jwt:managerBlacklist"
	err := app.RedisSingle().SRem(c, blacklistKey, uid).Err()
	if err != nil {
		return err
	}
	return nil
}

func IsTokenBlacklisted(c *gin.Context, blacklistKey string, uid int64) (bool, error) {
	//blacklistKey := "jwt:managerBlacklist"
	_, err := app.RedisSingle().Exists(c, blacklistKey).Result()
	if err != nil {
		return false, err
	}
	isBlacklisted, err := app.RedisSingle().SIsMember(c, blacklistKey, uid).Result()
	if err != nil {
		return false, err
	}
	return isBlacklisted, nil
}
