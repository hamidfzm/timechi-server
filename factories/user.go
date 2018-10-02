package factories

import (
	"github.com/hamidfzm/timechi-server/entities"
	"github.com/hamidfzm/timechi-server/models"
	"golang.org/x/crypto/bcrypt"
)

func GetUserPublicProfileV1(u *models.User) *entities.PublicProfileV1 {
	return &entities.PublicProfileV1{
		Name:  u.Name,
		Email: u.Email,
	}
}

func GetRegisterUser(e entities.RegisterV1) *models.User {
	password, _ := bcrypt.GenerateFromPassword([]byte(e.Password), bcrypt.DefaultCost)
	return &models.User{
		Name:     e.Name,
		Email:    e.Email,
		Password: password,
	}
}
