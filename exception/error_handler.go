package exception

import (
	"github.com/go-playground/validator/v10"
	"mhmmdnaufall/go-restful-api/helper"
	"mhmmdnaufall/go-restful-api/model"
	"net/http"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err any) {

	if badRequestError(writer, request, err) {
		return
	}

	if unauthorizedError(writer, request, err) {
		return
	}

	if validationErrors(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func validationErrors(writer http.ResponseWriter, request *http.Request, err any) bool {
	exception, ok := err.(validator.ValidationErrors)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := &model.WebResponse[any]{
			Errors: exception.Error(),
		}

		helper.WriteToResponseBody(writer, webResponse)

		return true

	} else {
		return false
	}
}

func badRequestError(writer http.ResponseWriter, request *http.Request, err any) bool {
	exception, ok := err.(*BadRequestError)

	if ok {

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := &model.WebResponse[any]{
			Errors: exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}

}

func unauthorizedError(writer http.ResponseWriter, request *http.Request, err any) bool {
	exception, ok := err.(*UnauthorizedError)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := &model.WebResponse[any]{
			Errors: exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)

		return true
	} else {
		return false
	}

}

func internalServerError(writer http.ResponseWriter, request *http.Request, err any) {

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := &model.WebResponse[any]{
		Errors: "INTERNAL SERVER ERROR",
	}

	helper.WriteToResponseBody(writer, webResponse)

}
