package repository

import (
	"context"
	"database/sql"
	"mhmmdnaufall/go-restful-api/entity"
	"mhmmdnaufall/go-restful-api/model"
)

type ContactRepository interface {
	Save(ctx context.Context, tx *sql.Tx, contact *entity.Contact) *entity.Contact

	Update(ctx context.Context, tx *sql.Tx, contact *entity.Contact) *entity.Contact

	Delete(ctx context.Context, tx *sql.Tx, contact *entity.Contact)

	FindByUserAndId(ctx context.Context, db *sql.DB, user *entity.User, id string) (*entity.Contact, error)

	Search(ctx context.Context, db *sql.DB, user *entity.User, request *model.SearchContactRequest) ([]*entity.Contact, int)
}
