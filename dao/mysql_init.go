package dao

import (
	"fmt"
	"goim/config"
	"goim/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func MysqlInit() {
	// 获取配置
	username := config.GetConfig().MySQL.Username
	password := config.GetConfig().MySQL.Password
	host := config.GetConfig().MySQL.Host
	port := config.GetConfig().MySQL.Port
	database := config.GetConfig().MySQL.Database
	timeout := "10s"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, database, timeout)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(fmt.Errorf("Fatal error connect database: %s \n", err))
	}

	DB = db
	sqlDB, _ := db.DB()

	// 设置连接池
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)

	// 迁移
	if err := db.AutoMigrate(&model.User{}); err != nil {
		panic(fmt.Errorf("Fatal error migrate database: %s \n", err))
	}
}
