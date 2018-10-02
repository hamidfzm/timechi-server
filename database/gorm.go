package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	
	"github.com/hamidfzm/timechi-server/config"
)

func NewGormDB(config *config.Config) (*gorm.DB, error) {
	return gorm.Open("sqlite3", config.DBName)
}
