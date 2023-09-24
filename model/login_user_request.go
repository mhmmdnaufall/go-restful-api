package model

type LoginUserRequest struct {
	Username string `validation:"required,max=100"`
	Password string `validation:"required,max=100"`
}
