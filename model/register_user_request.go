package model

type RegisterUserRequest struct {
	Username string `validate:"required,max=100,min=1" json:"username"`
	Password string `validate:"required,max=100,min=3" json:"password"`
	Name     string `validate:"required,max=100,min=1" json:"name"`
}
