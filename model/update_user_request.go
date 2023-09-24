package model

type UpdateUserRequest struct {
	Name     string `validate:"max=100" json:"name"`
	Password string `validate:"max=100" json:"password"`
}
