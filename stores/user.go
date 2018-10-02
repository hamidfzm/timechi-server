package stores

import (
	"github.com/jinzhu/gorm"
	"github.com/hamidfzm/timechi-server/config"
	"github.com/hamidfzm/timechi-server/models"
)

type User interface {
	Create(*models.User) error
	FindByEmail(string) (*models.User, error)
	FindByID(uint) (*models.User, error)
}

func NewUser(database *gorm.DB, config *config.Config) User {
	return &user{database, config}
}

type user struct {
	database *gorm.DB
	config   *config.Config
}

func (s *user) Create(m *models.User) error {
	return s.database.Create(m).Error
}

func (s *user) FindByEmail(email string) (m *models.User, err error) {
	err = s.database.Where("email = ?", email).First(&m).Error
	return
}

func (s *user) FindByID(id uint) (m *models.User, err error) {
	err = s.database.Where("id = ?", id).First(&m).Error
	return
}
