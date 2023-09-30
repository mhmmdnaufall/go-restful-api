// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"mhmmdnaufall/go-restful-api/app"
	impl3 "mhmmdnaufall/go-restful-api/controller/impl"
	"mhmmdnaufall/go-restful-api/middleware"
	"mhmmdnaufall/go-restful-api/repository/impl"
	impl2 "mhmmdnaufall/go-restful-api/service/impl"
	"net/http"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from injector.go:

func InitializeServer() *http.Server {
	userRepository := impl.NewUserRepository()
	db := app.NewDb()
	validate := ProvideValidator()
	userService := impl2.NewUserService(userRepository, db, validate)
	userController := impl3.NewUserController(userService)
	authService := impl2.NewAuthService(userRepository, db, validate)
	authController := impl3.NewAuthController(authService)
	contactRepository := impl.NewContactRepository()
	contactService := impl2.NewContactService(userRepository, contactRepository, db, validate)
	contactController := impl3.NewContactController(contactService)
	router := app.NewRouter(userController, authController, contactController)
	authMiddleware := middleware.NewAuthMiddleware(router, userRepository, db)
	server := NewServer(authMiddleware)
	return server
}

// injector.go:

var userSet = wire.NewSet(impl.NewUserRepository, impl2.NewUserService, impl3.NewUserController)

var authSet = wire.NewSet(impl2.NewAuthService, impl3.NewAuthController)

var contactSet = wire.NewSet(impl2.NewContactService, impl.NewContactRepository, impl3.NewContactController)
