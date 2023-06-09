package mysql

import (
	"fmt"
	"math/rand"

	"github.com/jlu-cow-studio/common/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dbPool []*gorm.DB

func Init() {
	// 定义连接字符串
	dsn := getDSN(config.DBUser, config.DBPassword, config.DBAddress, config.DBPort, config.DBDB)

	dbPool = make([]*gorm.DB, config.DBPoolSize)

	for i := 0; i < config.DBPoolSize; i++ {
		var err error
		if dbPool[i], err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			// SkipDefaultTransaction: true, // 禁用默认事务
			Logger:      logger.Default.LogMode(logger.Info),
			PrepareStmt: true, // 创建并缓存预编译语句
		}); err != nil {
			panic("failed to connect database")
		}
	}
}

func GetDBConn() *gorm.DB {
	return dbPool[rand.Intn(config.DBPoolSize)]
}

func getDSN(username, passwrod, address, port, db string) string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", username, passwrod, address, port, db)
}
