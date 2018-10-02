package entities

type RegisterV1 struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginV1 struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type PublicProfileV1 struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
