package models

import "time"

type Time struct {
	ID        uint          `gorm:"primary_key"`
	Title     string        `gorm:"not null"`
	StartedAt time.Time     `gorm:"not null"`
	StoppedAt time.Time     `gorm:"not null"`
	Duration  time.Duration `gorm:"not null"`
	UserID    uint          `gorm:"index;not null"`
}
