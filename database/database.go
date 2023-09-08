package database

import (
	"github.com/montexristos/haggle/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetDb() *gorm.DB {
	//CREATE SCHEMA `haggle` DEFAULT CHARACTER SET utf8 ;
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:123@(localhost:6602)/haggle?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err.Error())
	}
	err = db.AutoMigrate(&models.Event{}, &models.Market{}, &models.Selection{})
	if err != nil {
		panic(err.Error())
	}
	return db
}

func ClearDB() {
	db := GetDb()
	db.Where("1 = 1").Unscoped().Delete(&models.Selection{})
	db.Where("1 = 1").Unscoped().Delete(&models.Market{})
	db.Where("1 = 1").Unscoped().Delete(&models.Event{})
	if db.Error != nil {
		panic(db.Error.Error())
	}
}
