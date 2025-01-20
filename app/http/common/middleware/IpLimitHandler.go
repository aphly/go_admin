package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go_admin/app"
	"go_admin/app/res"
	"time"
)

func rateLimiter(rdb *redis.Client, c *gin.Context, key string, limit int64, period time.Duration) bool {
	val, err := rdb.Incr(c, key).Result()
	if err != nil {
		panic(err)
	}
	if val > limit {
		return false
	}
	if val == 1 {
		err = rdb.Expire(c, key, period).Err()
		if err != nil {
			panic(err)
		}
	}
	return true
}

func IpLimitHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		rdb := app.RedisSingle()
		key := "rate_limit:" + c.ClientIP()
		if rateLimiter(rdb, c, key, 10, time.Minute) {
			c.Next()
		} else {
			res.Json(c, res.Code(1), res.Msg("超出限制"))
			c.Abort()
			return
		}
	}
}
