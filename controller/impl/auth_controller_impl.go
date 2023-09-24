package impl

import (
	"github.com/julienschmidt/httprouter"
	"mhmmdnaufall/go-restful-api/controller"
	"mhmmdnaufall/go-restful-api/helper"
	"mhmmdnaufall/go-restful-api/model"
	"mhmmdnaufall/go-restful-api/service"
	"net/http"
)

type AuthControllerImpl struct {
	service.AuthService
}

func NewAuthController(authService service.AuthService) controller.AuthController {
	return &AuthControllerImpl{AuthService: authService}
}

func (authController *AuthControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	loginUserRequest := &model.LoginUserRequest{}
	helper.ReadFromRequestBody(request, loginUserRequest)

	tokenResponse := authController.AuthService.Login(request.Context(), loginUserRequest)

	webResponse := &model.WebResponse[*model.TokenResponse]{
		Data: tokenResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (authController *AuthControllerImpl) Logout(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	token := request.Header.Get("X-API-TOKEN")
	authController.AuthService.Logout(request.Context(), token)

	webResponse := &model.WebResponse[string]{
		Data: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
