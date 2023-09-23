package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type UserController interface {
	Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params)

	Get(writer http.ResponseWriter, request *http.Request, params httprouter.Params)

	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
