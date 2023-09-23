package app

import (
	"github.com/julienschmidt/httprouter"
	"mhmmdnaufall/go-restful-api/controller"
	"mhmmdnaufall/go-restful-api/exception"
)

func NewRouter(userController controller.UserController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/users", userController.Register)
	router.GET("/api/users/current", userController.Get)
	router.PATCH("/api/users/current", userController.Update)

	router.PanicHandler = exception.ErrorHandler

	return router
}
