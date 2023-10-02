package impl

import (
	"github.com/julienschmidt/httprouter"
	"mhmmdnaufall/go-restful-api/controller"
	"mhmmdnaufall/go-restful-api/helper"
	"mhmmdnaufall/go-restful-api/model"
	"mhmmdnaufall/go-restful-api/service"
	"net/http"
)

type AddressControllerImpl struct {
	service.AddressService
}

func NewAddressController(addressService service.AddressService) controller.AddressController {
	return &AddressControllerImpl{AddressService: addressService}
}

func (addressController *AddressControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userToken := request.Header.Get("X-API-TOKEN")

	addressRequest := &model.CreateAddressRequest{}
	helper.ReadFromRequestBody(request, addressRequest)
	addressRequest.ContactId = params.ByName("contactId")

	addressResponse := addressController.AddressService.Create(request.Context(), userToken, addressRequest)

	webResponse := &model.WebResponse[*model.AddressResponse]{
		Data: addressResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (addressController *AddressControllerImpl) Get(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userToken := request.Header.Get("X-API-TOKEN")
	contactId := params.ByName("contactId")
	addressId := params.ByName("addressId")

	addressResponse := addressController.AddressService.Get(request.Context(), userToken, contactId, addressId)

	webResponse := &model.WebResponse[*model.AddressResponse]{
		Data: addressResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (addressController *AddressControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userToken := request.Header.Get("X-API-TOKEN")

	addressRequest := &model.UpdateAddressRequest{}
	helper.ReadFromRequestBody(request, addressRequest)
	addressRequest.ContactId = params.ByName("contactId")
	addressRequest.AddressId = params.ByName("addressId")

	addressResponse := addressController.AddressService.Update(request.Context(), userToken, addressRequest)

	webResponse := &model.WebResponse[*model.AddressResponse]{
		Data: addressResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (addressController *AddressControllerImpl) Remove(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userToken := request.Header.Get("X-API-TOKEN")

	contactId := params.ByName("contactId")
	addressId := params.ByName("addressId")

	addressController.AddressService.Remove(request.Context(), userToken, contactId, addressId)

	webResponse := &model.WebResponse[string]{
		Data: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (addressController *AddressControllerImpl) List(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userToken := request.Header.Get("X-API-TOKEN")

	contactId := params.ByName("contactId")

	addressResponses := addressController.AddressService.List(request.Context(), userToken, contactId)

	webResponse := &model.WebResponse[[]*model.AddressResponse]{
		Data: addressResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
