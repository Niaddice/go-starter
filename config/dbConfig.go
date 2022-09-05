package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var newLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		SlowThreshold:             time.Second * 30,                              // Slow SQL threshold
		LogLevel:                  logger.LogLevel(GetConfig("log.level").(int)), // Log level
		IgnoreRecordNotFoundError: true,                                          // Ignore ErrRecordNotFound error for logger
		Colorful:                  false,                                         // Disable color
	},
)

func GetDb(dbType string) *gorm.DB {
	str := GetConfig(dbType)
	str1 := str.(map[string]interface{})
	username := str1["username"].(string)
	password := str1["password"].(string)
	url := str1["url"].(string)
	database := str1["database"].(string)
	var db, _ = gorm.Open(
		mysql.Open(username+":"+password+"@tcp("+url+")/"+database+"?charset=utf8mb4&parseTime=True&loc=Local"),
		&gorm.Config{Logger: newLogger, SkipDefaultTransaction: true})
	return db
}
