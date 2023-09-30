package impl

import (
	"context"
	"database/sql"
	"errors"
	"mhmmdnaufall/go-restful-api/entity"
	"mhmmdnaufall/go-restful-api/helper"
	"mhmmdnaufall/go-restful-api/repository"
)

type AddressRepository struct {
}

func NewAddressRepository() repository.AddressRepository {
	return &AddressRepository{}
}

func (addressRepository *AddressRepository) Save(ctx context.Context, tx *sql.Tx, address *entity.Address) *entity.Address {
	SQL := "INSERT INTO addresses(id, contact_id, street, city, province, country, postal_code) VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx, SQL,
		address.Id, address.Contact.Id, address.Street, address.City,
		address.Province, address.Country, address.PostalCode,
	)
	helper.PanicIfError(err)

	return address
}

func (addressRepository *AddressRepository) Update(ctx context.Context, tx *sql.Tx, address *entity.Address) *entity.Address {
	SQL := "UPDATE addresses SET street = ?, city = ?, province = ?, country = ?, postal_code = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL,
		address.Street, address.City, address.Province,
		address.Country, address.PostalCode, address.Id,
	)
	helper.PanicIfError(err)

	return address
}

func (addressRepository *AddressRepository) Delete(ctx context.Context, tx *sql.Tx, address *entity.Address) {
	SQL := "DELETE FROM addresses WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, address.Id)
	helper.PanicIfError(err)
}

func (addressRepository *AddressRepository) FindByContactAndId(ctx context.Context, db *sql.DB, contact *entity.Contact, id string) (*entity.Address, error) {
	SQL := `
		SELECT id, street, city, province, country, postal_code
		FROM addresses
		WHERE contact_id = ?
		  AND id = ?
	`
	rows, err := db.QueryContext(ctx, SQL, contact.Id, id)
	helper.PanicIfError(err)

	address := &entity.Address{}
	if rows.Next() {
		err := rows.Scan(&address.Id, &address.Street, &address.City, &address.Province, &address.Country, &address.PostalCode)
		helper.PanicIfError(err)
		address.Contact = contact
		return address, nil
	} else {
		return nil, errors.New("address not found")
	}

}

func (addressRepository *AddressRepository) FindAllByContact(ctx context.Context, db *sql.DB, contact *entity.Contact) []*entity.Address {
	SQL := `
		SELECT id, street, city, province, country, postal_code
		FROM addresses
		WHERE contact_id = ?
	`
	rows, err := db.QueryContext(ctx, SQL, contact.Id)
	helper.PanicIfError(err)

	var addresses []*entity.Address
	for rows.Next() {
		address := &entity.Address{}
		err := rows.Scan(&address.Id, &address.Street, &address.City, &address.Province, &address.Country, &address.PostalCode)
		helper.PanicIfError(err)
		address.Contact = contact
		addresses = append(addresses, address)
	}

	return addresses
}
