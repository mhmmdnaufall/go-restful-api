package impl

import (
	"github.com/julienschmidt/httprouter"
	"mhmmdnaufall/go-restful-api/controller"
	"mhmmdnaufall/go-restful-api/helper"
	"mhmmdnaufall/go-restful-api/model"
	"mhmmdnaufall/go-restful-api/service"
	"net/http"
)

type UserControllerImpl struct {
	service.UserService
}

func NewUserController(userService service.UserService) controller.UserController {
	return &UserControllerImpl{UserService: userService}
}

func (userController *UserControllerImpl) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	registerUserRequest := &model.RegisterUserRequest{}
	helper.ReadFromRequestBody(request, registerUserRequest)

	err := userController.UserService.Register(request.Context(), registerUserRequest)
	helper.PanicIfError(err)

	webResponse := &model.WebResponse[string]{
		Data: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (userController *UserControllerImpl) Get(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	token := request.Header.Get("X-API-TOKEN")
	userResponse := userController.UserService.Get(request.Context(), token)

	helper.WriteToResponseBody(writer, userResponse)
}

func (userController *UserControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	updateUserRequest := &model.UpdateUserRequest{}
	helper.ReadFromRequestBody(request, updateUserRequest)

	token := request.Header.Get("X-API-TOKEN")
	userResponse := userController.UserService.Update(request.Context(), token, updateUserRequest)

	webResponse := &model.WebResponse[*model.UserResponse]{
		Data: userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
