package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique_index"`
	Password []byte
}

func (m *User) Create() error {
	return DB.Create(m).Error
}
