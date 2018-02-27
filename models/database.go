package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func SetupDatabase() {
	if db, err := gorm.Open("sqlite3", "test.db"); err != nil {
		panic("failed to connect database")
	} else {
		DB = db
	}
	
	DB.AutoMigrate(&User{})
}
