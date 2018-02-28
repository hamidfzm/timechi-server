package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB
var tables = []interface{}{&User{}}

func SetupDatabase() {
	if db, err := gorm.Open("sqlite3", "develop.db"); err != nil {
		panic("failed to connect database")
	} else {
		DB = db
	}
	
	DB.LogMode(false)
	DB.AutoMigrate(tables...)
}

func SetupTestDatabase() {
	if db, err := gorm.Open("sqlite3", "test.db"); err != nil {
		panic("failed to connect database")
	} else {
		DB = db
	}
	
	DB.LogMode(false)
	DB.DropTableIfExists(tables...)
	DB.AutoMigrate(tables...)
}
