package repository

import (
	"context"
	"database/sql"
	"mhmmdnaufall/go-restful-api/entity"
)

type AddressRepository interface {
	Save(ctx context.Context, tx *sql.Tx, address *entity.Address) *entity.Address

	Update(ctx context.Context, tx *sql.Tx, address *entity.Address) *entity.Address

	Delete(ctx context.Context, tx *sql.Tx, address *entity.Address)

	FindByContactAndId(ctx context.Context, db *sql.DB, contact *entity.Contact, id string) (*entity.Address, error)

	FindAllByContact(ctx context.Context, db *sql.DB, contact *entity.Contact) []*entity.Address
}
