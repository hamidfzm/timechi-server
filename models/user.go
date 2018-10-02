package models

import (
	"github.com/jinzhu/gorm"
	"time"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name         string `gorm:"not null"`
	Email        string `gorm:"unique_index"`
	Password     []byte `gorm:"not null"`
	TimerTitle   *string
	TimerStartAt *time.Time
}

func (r *User) VerifyPassword(password string) error {
	return bcrypt.CompareHashAndPassword(r.Password, []byte(password))
}
