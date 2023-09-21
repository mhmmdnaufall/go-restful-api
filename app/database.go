package app

import (
	"database/sql"
	"mhmmdnaufall/go-restful-api/helper"
	"time"
)

func NewDb() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/go_restful_api")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(50 * time.Minute)
	db.SetConnMaxIdleTime(20 * time.Minute)

	return db
}
