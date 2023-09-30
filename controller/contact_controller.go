package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type ContactController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)

	Get(writer http.ResponseWriter, request *http.Request, params httprouter.Params)

	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)

	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)

	Search(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
