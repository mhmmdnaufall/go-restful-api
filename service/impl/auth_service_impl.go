package impl

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"mhmmdnaufall/go-restful-api/exception"
	"mhmmdnaufall/go-restful-api/helper"
	"mhmmdnaufall/go-restful-api/model"
	"mhmmdnaufall/go-restful-api/repository"
	"mhmmdnaufall/go-restful-api/service"
	"time"
)

type AuthServiceImpl struct {
	repository.UserRepository
	*sql.DB
	*validator.Validate
}

func NewAuthService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) service.AuthService {
	return &AuthServiceImpl{UserRepository: userRepository, DB: DB, Validate: validate}
}

func (authService *AuthServiceImpl) Login(ctx context.Context, request *model.LoginUserRequest) *model.TokenResponse {
	err := authService.Validate.Struct(request)
	helper.PanicIfError(err)

	user, err := authService.UserRepository.FindByUsername(ctx, authService.DB, request.Username)
	if err != nil {
		panic(exception.NewUnauthorizedError("Username or password wrong"))
		// jangan ngasih tau yang mana yang salah, takutnya nanti dicoba-coba
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err == nil {

		user.Token = sql.NullString{
			String: uuid.New().String(),
			Valid:  true,
		}

		user.TokenExpiredAt = sql.NullInt64{
			Int64: authService.next30Days(),
			Valid: true,
		}

		tx, err := authService.DB.Begin()
		helper.PanicIfError(err)
		defer helper.CommitOrRollback(tx)

		authService.UserRepository.Update(ctx, tx, user)

		return &model.TokenResponse{
			Token:     user.Token.String,
			ExpiredAt: user.TokenExpiredAt.Int64,
		}
	} else {
		// gagal
		panic(exception.NewUnauthorizedError("Username or password wrong"))
	}
}

func (authService *AuthServiceImpl) Logout(ctx context.Context, userToken string) {
	user, err := authService.UserRepository.FindByToken(ctx, authService.DB, userToken)
	helper.PanicIfError(err)

	user.Token.String = ""
	user.Token.Valid = false

	user.TokenExpiredAt.Int64 = 0
	user.TokenExpiredAt.Valid = false

	tx, err := authService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	authService.UserRepository.Update(ctx, tx, user)
}

func (authService *AuthServiceImpl) next30Days() int64 {
	return time.Now().UnixMilli() + (1000 * 60 * 60 * 24 * 30)
}
