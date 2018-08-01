package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/hamidfzm/timechi-server/helpers"
)

type User struct {
	gorm.Model
	Name         string `gorm:"not null"`
	Email        string `gorm:"unique_index"`
	Password     []byte `gorm:"not null"`
	TimerTitle   *string
	TimerStartAt *time.Time
}

func (m *User) Create() error {
	return DB.Create(m).Error
}

func (m *User) VerifyPassword(password string) error {
	return bcrypt.CompareHashAndPassword(m.Password, []byte(password))
}

func (m *User) StartTimer(title string) error {
	now := time.Now()
	tx := DB.Model(m).
		Where("timer_title IS NULL AND timer_start_at IS NULL").
		Updates(map[string]interface{}{"timer_title": &title, "timer_start_at": &now})
	
	if tx.Error != nil {
		return tx.Error
	} else if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	} else {
		return nil
	}
}

func (m *User) StopTimer() (t Time, err error) {
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	
	if tx.Error != nil {
		err = tx.Error
		return
	}
	
	if m.TimerStartAt != nil && m.TimerTitle != nil {
		now := time.Now()
		t = Time{
			UserID:    m.ID,
			Title:     *m.TimerTitle,
			StartedAt: *m.TimerStartAt,
			StoppedAt: now,
			Duration:  now.Sub(*m.TimerStartAt),
		}
		
		if err = tx.Create(&t).Error; err != nil {
			tx.Rollback()
			return
		}
	} else {
		tx.Rollback()
		err = gorm.ErrRecordNotFound
		return
	}
	
	if subTx := tx.Model(m).
		Where("timer_title IS NOT NULL AND timer_start_at IS NOT NULL").
		Updates(map[string]interface{}{"timer_title": nil, "timer_start_at": nil}); subTx.Error != nil {
		tx.Rollback()
		err = subTx.Error
		return
	} else if subTx.RowsAffected == 0 {
		tx.Rollback()
		err = gorm.ErrRecordNotFound
		return
	}
	
	err = tx.Commit().Error
	return
}

func (m *User) Times() (times []Time) {
	DB.Model(m).Related(&times)
	return
}

type TokenClaims struct {
	jwt.StandardClaims
	ID uint `json:"id"`
}

func (m *User) GenerateToken() (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		TokenClaims{
			ID: m.ID,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: now.Add(time.Hour * 48).Unix(),
			},
		})
	
	return token.SignedString([]byte(helpers.Config.Secret))
}

func FindUserByEmail(email string) (*User, error) {
	var m User
	err := DB.Where("email = ?", email).First(&m).Error
	return &m, err
}

func FindUserByID(id uint) (*User, error) {
	var m User
	err := DB.Where("id = ?", id).First(&m).Error
	return &m, err
}
