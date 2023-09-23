package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"mhmmdnaufall/go-restful-api/app"
	controller "mhmmdnaufall/go-restful-api/controller/impl"
	"mhmmdnaufall/go-restful-api/helper"
	"mhmmdnaufall/go-restful-api/middleware"
	repository "mhmmdnaufall/go-restful-api/repository/impl"
	service "mhmmdnaufall/go-restful-api/service/impl"
	"net/http"
)

func main() {

	DB := app.NewDb()

	validate := validator.New()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, DB, validate)
	userController := controller.NewUserController(userService)
	router := app.NewRouter(userController)

	authMiddlerware := &middleware.AuthMiddleware{
		Handler:        router,
		UserRepository: userRepository,
		DB:             DB,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: authMiddlerware,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
