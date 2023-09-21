package entity

import "database/sql"

type Contact struct {
	Id        string
	FirstName string
	LastName  sql.NullString
	Phone     sql.NullString
	Email     sql.NullString
	User      *User
	Addresses []*Address
}
