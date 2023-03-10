package models

import (
	"fmt"
	"log"
	"reflect"
	"sync"
	"time"

	"github.com/cm9510/cm-goweb-api/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	err  error
	once sync.Once
)

func GetDb() *gorm.DB {
	if reflect.ValueOf(db).IsZero() {
		initDB()
	}
	return db
}

func initDB() {
	config := config.GetConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local&timeout=%s",
		config.Mysql.Username,
		config.Mysql.Password,
		config.Mysql.Host,
		config.Mysql.Port,
		config.Mysql.Dbname,
		config.Mysql.Charset,
		config.Mysql.Timeout,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		QueryFields:            true,
	})
	if err != nil {
		log.Fatal("grom open connection err:", err)
	}
	// gorm.ErrRecordNotFound
	sqlDb, _ := db.DB()
	sqlDb.SetConnMaxIdleTime(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Minute)
}
