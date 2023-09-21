package model

type UpdateUserRequest struct {
	Name     string `validate:"max=100,min=1" json:"name"`
	Password string `validate:"max=100,min=3" json:"password"`
}
