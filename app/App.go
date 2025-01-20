package app

import (
	"github.com/redis/go-redis/v9"
	"go_admin/app/core/config"
	"go_admin/app/core/connect"
	"go_admin/app/core/log"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
	"os"
	"strconv"
)

var Config config.Config

func Init() {
	data, err := os.ReadFile("config/env.yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		panic(err)
	}
	//Config = config.NewConfig()
}

func Db(keys ...string) *gorm.DB {
	key0 := "default"
	key1 := 0
	if len(keys) == 1 {
		key0 = keys[0]
	} else if len(keys) == 2 {
		key0 = keys[0]
		key1, _ = strconv.Atoi(keys[1])
	}
	if len((Config.Db)[key0]) <= 0 {
		panic("数据库配置错误")
	}
	return connect.Mysql(Config.Db, key0, key1)
}

// app.Log("sss").debug("wwwwww")
func Log(names ...string) *log.Logger {
	name := "default"
	if len(names) > 0 {
		name = names[0]
	}
	return log.NewLogger(Config.Log, name)
}

//func RedisSingle(keys ...int) *redis.Client {
//	key := 0
//	if len(keys) > 0 {
//		key = keys[0]
//	}
//	if len(Config.Redis.Single) <= 0 {
//		panic("Redis配置错误")
//	}
//	return connect.Redis(Config.Redis, key)
//}

func RedisSingle() *redis.Client {
	if Config.Redis.Single.Addr == "" {
		panic("Redis配置错误")
	}
	return redis.NewClient(&redis.Options{
		Addr:       Config.Redis.Single.Addr,
		Password:   Config.Redis.Single.Password,
		DB:         Config.Redis.Single.Db,
		PoolSize:   Config.Redis.Single.PoolSize,
		MaxRetries: Config.Redis.Single.Retries,
	})
}

func RedisCluster() *redis.ClusterClient {
	return redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:      Config.Redis.Cluster.Addrs,
		Password:   Config.Redis.Cluster.Password,
		PoolSize:   Config.Redis.Cluster.PoolSize,
		MaxRetries: Config.Redis.Cluster.Retries,
	})
}
