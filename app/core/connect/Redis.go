package connect

//var Rediss = make(map[int]*redis.Client)
//var RedisLocker sync.RWMutex
//
//func Redis(config config.Redis, key int) *redis.Client {
//	RedisLocker.RLock()
//	rd, ok := Rediss[key]
//	if ok {
//		RedisLocker.RUnlock()
//		return rd
//	}
//	RedisLocker.RUnlock()
//
//	RedisLocker.Lock()
//	defer RedisLocker.Unlock()
//	if _, ok1 := Rediss[key]; ok1 {
//		return Rediss[key]
//	}
//	Rediss[key] = getRedis(config, key)
//	return Rediss[key]
//}

//	func getRedis(config config.Redis, key int) *redis.Client {
//		return redis.NewClient(&redis.Options{
//			Addr:       config.Single[key].Addr,
//			Password:   config.Single[key].Password,
//			DB:         config.Single[key].Db,
//			PoolSize:   config.Single[key].PoolSize,
//			MaxRetries: config.Single[key].Retries,
//		})
//	}

//	func RedisSingle(keys ...int) *redis.Client {
//		key := 0
//		if len(keys) > 0 {
//			key = keys[0]
//		}
//		if len(Config.Redis.Single) <= 0 {
//			panic("Redis配置错误")
//		}
//		return connect.Redis(Config.Redis, key)
//	}
