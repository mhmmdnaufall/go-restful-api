package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type AddressController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)

	Get(writer http.ResponseWriter, request *http.Request, params httprouter.Params)

	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)

	Remove(writer http.ResponseWriter, request *http.Request, params httprouter.Params)

	List(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
