package entity

import "database/sql"

type Address struct {
	Id         string
	Street     sql.NullString
	City       sql.NullString
	Province   sql.NullString
	Country    string
	PostalCode sql.NullString
	Contact    *Contact
}
