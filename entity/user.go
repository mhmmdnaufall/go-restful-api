package entity

import "database/sql"

type User struct {
	Username       string
	Password       string
	Name           string
	Token          sql.NullString
	TokenExpiredAt sql.NullInt64
	Contacts       []*Contact
}
