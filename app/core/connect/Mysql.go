package connect

import (
	"fmt"
	"go_admin/app/core/config"
	"go_admin/app/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"strconv"
	"sync"
	"time"
)

var dbs = make(map[string]*gorm.DB)
var dbLocker sync.RWMutex

func Mysql(config map[string][]config.DbSingle, key0 string, key1 int) *gorm.DB {
	str := key0 + strconv.Itoa(key1)
	dbLocker.RLock()
	db, ok := dbs[str]
	if ok {
		dbLocker.RUnlock()
		return db
	}
	dbLocker.RUnlock()

	dbLocker.Lock()
	defer dbLocker.Unlock()
	if _, ok1 := dbs[str]; ok1 {
		return dbs[str]
	}
	dbs[str] = getDb(config, key0, key1)
	return dbs[str]
}

func getDb(config map[string][]config.DbSingle, key0 string, key1 int) *gorm.DB {
	var LWriter helper.WriterLog
	newLogger := logger.New(
		LWriter,
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Warn, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local&timeout=%dms&writeTimeout=%dms&readTimeout=%dms",
		config[key0][key1].Username,
		config[key0][key1].Password,
		config[key0][key1].Host,
		config[key0][key1].Port,
		config[key0][key1].Database,
		config[key0][key1].Charset,
		config[key0][key1].TimeOut,
		config[key0][key1].WriteTimeOut,
		config[key0][key1].ReadTimeOut,
	)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 255,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",   // 表前缀
			SingularTable: true, // 禁用表名复数
		},
		Logger: newLogger,
	})
	if err != nil {
		panic("连接数据库失败, error:" + err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic("数据库池, error:" + err.Error())
	}
	sqlDB.SetMaxIdleConns(config[key0][key1].MaxIdleConnect)
	sqlDB.SetMaxOpenConns(config[key0][key1].MaxOpenConnect)
	sqlDB.SetConnMaxLifetime(time.Second * 120)
	return db
}
