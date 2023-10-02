//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"mhmmdnaufall/go-restful-api/app"
	controllerImpl "mhmmdnaufall/go-restful-api/controller/impl"
	"mhmmdnaufall/go-restful-api/middleware"
	repositoryImpl "mhmmdnaufall/go-restful-api/repository/impl"
	serviceImpl "mhmmdnaufall/go-restful-api/service/impl"
	"net/http"
)

var userSet = wire.NewSet(
	repositoryImpl.NewUserRepository,
	serviceImpl.NewUserService,
	controllerImpl.NewUserController,
)

var authSet = wire.NewSet(
	serviceImpl.NewAuthService,
	controllerImpl.NewAuthController,
)

var contactSet = wire.NewSet(
	serviceImpl.NewContactService,
	repositoryImpl.NewContactRepository,
	controllerImpl.NewContactController,
)

var addressSet = wire.NewSet(
	repositoryImpl.NewAddressRepository,
	serviceImpl.NewAddressService,
	controllerImpl.NewAddressController,
)

func InitializeServer() *http.Server {
	wire.Build(
		app.NewDb,
		ProvideValidator,
		userSet,
		authSet,
		contactSet,
		addressSet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}
