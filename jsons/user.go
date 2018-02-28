package jsons

import (
	"github.com/hamidfzm/timechi-server/models"
	"github.com/hamidfzm/timechi-server/helpers"
)

type RegisterV1 struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (j *RegisterV1) To(m *models.User) {
	m.Name = j.Name
	m.Email = j.Email
	m.Password = helpers.HashPassword(j.Password)
}

type PublicProfileV1 struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (j *PublicProfileV1) From(m *models.User) {
	j.Name = m.Name
	j.Email = m.Email
}
