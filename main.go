package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"mhmmdnaufall/go-restful-api/helper"
	"mhmmdnaufall/go-restful-api/middleware"
	"net/http"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:8080",
		Handler: authMiddleware,
	}
}

func ProvideValidator() *validator.Validate {
	return validator.New()
}

func main() {

	server := InitializeServer()

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
