package middleware

import (
	"database/sql"
	"encoding/json"
	"mhmmdnaufall/go-restful-api/helper"
	"mhmmdnaufall/go-restful-api/model"
	"mhmmdnaufall/go-restful-api/repository"
	"net/http"
	"strings"
)

var permitPath = map[string]string{
	"/api/users":      http.MethodPost, // register user
	"/api/auth/login": http.MethodPost, // login user
}

type AuthMiddleware struct {
	Handler        http.Handler
	UserRepository repository.UserRepository
	DB             *sql.DB
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	requestPath := request.URL.Path
	requestMethod := request.Method

	allowedMethod, pathAllowed := permitPath[requestPath]

	token := request.Header.Get("X-API-TOKEN")
	_, err := middleware.UserRepository.FindByToken(request.Context(), middleware.DB, token)

	if (pathAllowed) && (allowedMethod == requestMethod) {

		middleware.Handler.ServeHTTP(writer, request)

		return

	}

	if (len(strings.TrimSpace(token)) == 0) || (err != nil) {

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := &model.WebResponse[any]{
			Errors: "Unauthorized",
		}

		encoder := json.NewEncoder(writer)
		err := encoder.Encode(webResponse)
		helper.PanicIfError(err)

		return

	}

	middleware.Handler.ServeHTTP(writer, request)

}
