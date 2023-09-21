package service

import (
	"context"
	"mhmmdnaufall/go-restful-api/model"
)

type UserService interface {
	Register(ctx context.Context, request *model.RegisterUserRequest) error

	Get(ctx context.Context, userToken string) *model.UserResponse

	Update(ctx context.Context, userToken string, request *model.UpdateUserRequest) *model.UserResponse
}
