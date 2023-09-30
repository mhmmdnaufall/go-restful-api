package impl

import (
	"context"
	"database/sql"
	"errors"
	"mhmmdnaufall/go-restful-api/entity"
	"mhmmdnaufall/go-restful-api/helper"
	"mhmmdnaufall/go-restful-api/model"
	"mhmmdnaufall/go-restful-api/repository"
	"strconv"
	"strings"
)

type ContactRepositoryImpl struct {
}

func NewContactRepository() repository.ContactRepository {
	return &ContactRepositoryImpl{}
}

func (contactRepository *ContactRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, contact *entity.Contact) *entity.Contact {
	SQL := "INSERT INTO contacts(id, username, first_name, last_name, phone, email) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx, SQL,
		contact.Id, contact.User.Username, contact.FirstName,
		contact.LastName, contact.Phone, contact.Email,
	)
	helper.PanicIfError(err)

	return contact
}

func (contactRepository *ContactRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, contact *entity.Contact) *entity.Contact {
	SQL := "UPDATE contacts SET first_name = ?, last_name = ?, email = ?, phone = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, contact.FirstName, contact.LastName, contact.Email, contact.Phone, contact.Id)
	helper.PanicIfError(err)

	return contact
}

func (contactRepository *ContactRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, contact *entity.Contact) {
	SQL := "DELETE FROM contacts WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, contact.Id)
	helper.PanicIfError(err)
}

func (contactRepository *ContactRepositoryImpl) FindByUserAndId(ctx context.Context, db *sql.DB, user *entity.User, id string) (*entity.Contact, error) {
	SQL := `
		SELECT id, first_name, last_name, phone, email
		FROM contacts
		WHERE username = ?
		  AND id = ?
	`
	rows, err := db.QueryContext(ctx, SQL, user.Username, id)
	helper.PanicIfError(err)
	defer rows.Close()

	contact := &entity.Contact{}
	if rows.Next() {
		err := rows.Scan(&contact.Id, &contact.FirstName, &contact.LastName, &contact.Phone, &contact.Email)
		helper.PanicIfError(err)
		contact.User = user
		return contact, nil
	} else {
		return nil, errors.New("contact not found")
	}
}

func (contactRepository *ContactRepositoryImpl) Search(ctx context.Context, db *sql.DB, user *entity.User, request *model.SearchContactRequest) ([]*entity.Contact, int) {
	SQL := "SELECT id, first_name, last_name, email, phone FROM contacts"
	query := " WHERE username = ?"
	params := []any{user.Username}

	if len(strings.TrimSpace(request.Name)) != 0 {
		query += " AND (first_name LIKE ? OR last_name LIKE ?)"
		params = append(params, "%"+request.Name+"%", "%"+request.Name+"%")
	}

	if len(strings.TrimSpace(request.Email)) != 0 {
		query += " AND email LIKE ?"
		params = append(params, "%"+request.Email+"%")
	}

	if len(strings.TrimSpace(request.Phone)) != 0 {
		query += " AND phone LIKE ?"
		params = append(params, "%"+request.Phone+"%")
	}

	SQL += query
	SQL += " LIMIT " + strconv.Itoa(request.Page*request.Size) + ", " + strconv.Itoa(request.Size)

	var totalPage int
	totalPageSQL := "SELECT CEIL(COUNT(*) / " + strconv.Itoa(request.Size) + ") FROM contacts"
	totalPageSQL += query
	rows, err := db.QueryContext(ctx, totalPageSQL, params...)
	helper.PanicIfError(err)

	if rows.Next() {
		err := rows.Scan(&totalPage)
		helper.PanicIfError(err)
	}

	rows, err = db.QueryContext(ctx, SQL, params...)
	helper.PanicIfError(err)

	var contacts []*entity.Contact
	for rows.Next() {
		contact := &entity.Contact{}
		err := rows.Scan(&contact.Id, &contact.FirstName, &contact.LastName, &contact.Email, &contact.Phone)
		helper.PanicIfError(err)
		contact.User = user
		contacts = append(contacts, contact)
	}

	return contacts, totalPage
}
