package service

import (
	"context"
	"mhmmdnaufall/go-restful-api/model"
)

type AuthService interface {
	Login(ctx context.Context, request *model.LoginUserRequest) *model.TokenResponse

	Logout(ctx context.Context, userToken string)
}
