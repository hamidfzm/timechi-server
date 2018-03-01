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
	Name     string
	Email    string `gorm:"unique_index"`
	Password []byte
}

func (m *User) Create() error {
	return DB.Create(m).Error
}

func (m *User) VerifyPassword(password string) error {
	return bcrypt.CompareHashAndPassword(m.Password, []byte(password))
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
