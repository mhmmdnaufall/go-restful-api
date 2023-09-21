package repository

import (
	"context"
	"database/sql"
	"mhmmdnaufall/go-restful-api/entity"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user *entity.User) *entity.User

	Update(ctx context.Context, tx *sql.Tx, user *entity.User) *entity.User

	IsUserExists(ctx context.Context, db *sql.DB, username string) bool

	FindByUsername(ctx context.Context, db *sql.DB, username string) (*entity.User, error)

	FindByToken(ctx context.Context, db *sql.DB, token string) (*entity.User, error)
}
