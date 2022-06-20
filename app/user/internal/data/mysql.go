package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"helloworld/app/user/internal/biz"
	"helloworld/app/user/internal/conf"
	"log"
	"os"
	"time"
)

func NewMysqlRepo(c *conf.Data) *gorm.DB {

	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  logger.Warn,
			Colorful:                  true,
			IgnoreRecordNotFoundError: false,
		}),
	})
	if err != nil {
		log.Println("failed to connect mysql:", err)
		panic(err)
	}
	dbPool, err := db.DB()
	if err != nil {
		log.Println("failed to connect mysql:", err)
		panic(err)
	}

	dbPool.SetMaxIdleConns(50)
	dbPool.SetMaxOpenConns(100)
	// 超时
	dbPool.SetConnMaxLifetime(time.Second * 30)

	//GORM将完成自动建表
	err = db.AutoMigrate(
		&biz.User{},
	)
	if err != nil {
		log.Println("auto migrate failed:", err)
		panic(err)
	}

	return db
}
