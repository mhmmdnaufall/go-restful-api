package app

import (
	"github.com/julienschmidt/httprouter"
	"mhmmdnaufall/go-restful-api/controller"
	"mhmmdnaufall/go-restful-api/exception"
)

func NewRouter(userController controller.UserController, authController controller.AuthController) *httprouter.Router {
	router := httprouter.New()

	// user api
	router.POST("/api/users", userController.Register)
	router.GET("/api/users/current", userController.Get)
	router.PATCH("/api/users/current", userController.Update)

	// auth api
	router.POST("/api/auth/login", authController.Login)
	router.DELETE("/api/auth/logout", authController.Logout)

	router.PanicHandler = exception.ErrorHandler

	return router
}
