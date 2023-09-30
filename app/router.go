package app

import (
	"github.com/julienschmidt/httprouter"
	"mhmmdnaufall/go-restful-api/controller"
	"mhmmdnaufall/go-restful-api/exception"
)

func NewRouter(userController controller.UserController, authController controller.AuthController, contactController controller.ContactController) *httprouter.Router {
	router := httprouter.New()

	// user api
	router.POST("/api/users", userController.Register)
	router.GET("/api/users/current", userController.Get)
	router.PATCH("/api/users/current", userController.Update)

	// auth api
	router.POST("/api/auth/login", authController.Login)
	router.DELETE("/api/auth/logout", authController.Logout)

	// contact api
	router.POST("/api/contacts", contactController.Create)
	router.GET("/api/contacts/:contactId", contactController.Get)
	router.PUT("/api/contacts/:contactId", contactController.Update)
	router.DELETE("/api/contacts/:contactId", contactController.Delete)
	router.GET("/api/contacts", contactController.Search)

	router.PanicHandler = exception.ErrorHandler

	return router
}
