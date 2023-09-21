package impl

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"mhmmdnaufall/go-restful-api/entity"
	"mhmmdnaufall/go-restful-api/exception"
	"mhmmdnaufall/go-restful-api/helper"
	"mhmmdnaufall/go-restful-api/model"
	"mhmmdnaufall/go-restful-api/repository"
	"mhmmdnaufall/go-restful-api/service"
	"strings"
)

type UserServiceImpl struct {
	repository.UserRepository
	*sql.DB
	*validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) service.UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (userService *UserServiceImpl) Register(ctx context.Context, request *model.RegisterUserRequest) error {
	err := userService.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := userService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	if userService.UserRepository.IsUserExists(ctx, userService.DB, request.Username) {
		panic(exception.NewBadRequestError("username already registered"))
	}

	encodedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	helper.PanicIfError(err)

	user := &entity.User{
		Username: request.Username,
		Password: string(encodedPassword),
		Name:     request.Name,
	}

	userService.UserRepository.Save(ctx, tx, user)
	return nil
}

func (userService *UserServiceImpl) Get(ctx context.Context, userToken string) *model.UserResponse {
	tx, err := userService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := userService.UserRepository.FindByToken(ctx, userService.DB, userToken)
	helper.PanicIfError(err)

	return &model.UserResponse{
		Username: user.Username,
		Name:     user.Name,
	}
}

func (userService *UserServiceImpl) Update(ctx context.Context, userToken string, request *model.UpdateUserRequest) *model.UserResponse {
	err := userService.Validate.Struct(request)
	helper.PanicIfError(err)

	user, err := userService.UserRepository.FindByToken(ctx, userService.DB, userToken)
	helper.PanicIfError(err)

	if len(strings.TrimSpace(request.Name)) != 0 {
		user.Name = request.Name
	}

	if len(strings.TrimSpace(request.Password)) != 0 {
		password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		helper.PanicIfError(err)
		user.Password = string(password)
	}

	tx, err := userService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	userService.UserRepository.Update(ctx, tx, user)

	return &model.UserResponse{
		Name:     user.Name,
		Username: user.Username,
	}
}
