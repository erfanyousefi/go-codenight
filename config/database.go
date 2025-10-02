package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GORMDB *gorm.DB

func InitDB(cfg *ConfigEnv) {
	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)
	var dialector gorm.Dialector = mysql.Open(dsn)
	var err error
	GORMDB, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatal("failed connected to DB!")
	}
	fmt.Println("connected to DB successfully ")
}
