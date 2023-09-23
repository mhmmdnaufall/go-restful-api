package impl

import (
	"context"
	"database/sql"
	"errors"
	"mhmmdnaufall/go-restful-api/entity"
	"mhmmdnaufall/go-restful-api/helper"
	"mhmmdnaufall/go-restful-api/repository"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() repository.UserRepository {
	return &UserRepositoryImpl{}

}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user *entity.User) *entity.User {
	SQL := "INSERT INTO users(username, password, name, token, token_expired_at) VALUES (?, ?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx, SQL, user.Username, user.Password, user.Name, user.Token, user.TokenExpiredAt)
	helper.PanicIfError(err)

	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user *entity.User) *entity.User {
	SQL := "UPDATE users SET password = ?, name = ? WHERE username = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Password, user.Name, user.Username)
	helper.PanicIfError(err)

	return user
}

func (repository *UserRepositoryImpl) IsUserExists(ctx context.Context, db *sql.DB, username string) bool {
	SQL := "SELECT * FROM users WHERE username = ?"
	rows, err := db.QueryContext(ctx, SQL, username)
	helper.PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		return true
	} else {
		return false
	}

}

func (repository *UserRepositoryImpl) FindByUsername(ctx context.Context, db *sql.DB, username string) (*entity.User, error) {
	SQL := "SELECT username, password, name, token, token_expired_at FROM users WHERE username = ?"
	rows, err := db.QueryContext(ctx, SQL, username)
	helper.PanicIfError(err)
	defer rows.Close()

	user := &entity.User{}

	if rows.Next() {
		err := rows.Scan(&user.Username, &user.Password, &user.Name, &user.Token, &user.TokenExpiredAt)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return nil, errors.New("user not found")
	}
}

func (repository *UserRepositoryImpl) FindByToken(ctx context.Context, db *sql.DB, token string) (*entity.User, error) {
	SQL := "SELECT username, password, name, token, token_expired_at FROM users WHERE token = ?"
	rows, err := db.QueryContext(ctx, SQL, token)
	helper.PanicIfError(err)
	defer rows.Close()

	user := &entity.User{}
	if rows.Next() {
		err := rows.Scan(&user.Username, &user.Password, &user.Name, &user.Token, &user.TokenExpiredAt)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return nil, errors.New("user not found")
	}
}
