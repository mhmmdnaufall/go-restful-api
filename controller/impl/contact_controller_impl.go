package impl

import (
	"github.com/julienschmidt/httprouter"
	"mhmmdnaufall/go-restful-api/controller"
	"mhmmdnaufall/go-restful-api/helper"
	"mhmmdnaufall/go-restful-api/model"
	"mhmmdnaufall/go-restful-api/service"
	"net/http"
	"strconv"
)

type ContactControllerImpl struct {
	service.ContactService
}

func NewContactController(contactService service.ContactService) controller.ContactController {
	return &ContactControllerImpl{ContactService: contactService}
}

func (contactController *ContactControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userToken := request.Header.Get("X-API-TOKEN")

	contactRequest := &model.CreateContactRequest{}
	helper.ReadFromRequestBody(request, contactRequest)

	contactResponse := contactController.ContactService.Create(request.Context(), userToken, contactRequest)

	webResponse := &model.WebResponse[*model.ContactResponse]{
		Data: contactResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (contactController *ContactControllerImpl) Get(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userToken := request.Header.Get("X-API-TOKEN")
	contactId := params.ByName("contactId")
	contactResponse := contactController.ContactService.Get(request.Context(), userToken, contactId)

	webResponse := &model.WebResponse[*model.ContactResponse]{
		Data: contactResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (contactController *ContactControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userToken := request.Header.Get("X-API-TOKEN")

	updateContactRequest := &model.UpdateContactRequest{}
	helper.ReadFromRequestBody(request, updateContactRequest)

	contactId := params.ByName("contactId")
	updateContactRequest.Id = contactId

	contactResponse := contactController.ContactService.Update(request.Context(), userToken, updateContactRequest)

	webResponse := &model.WebResponse[*model.ContactResponse]{
		Data: contactResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (contactController *ContactControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userToken := request.Header.Get("X-API-TOKEN")
	contactId := params.ByName("contactId")

	contactController.ContactService.Delete(request.Context(), userToken, contactId)

	helper.WriteToResponseBody(writer, &model.WebResponse[string]{Data: "OK"})
}

func (contactController *ContactControllerImpl) Search(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userToken := request.Header.Get("X-API-TOKEN")

	name := request.URL.Query().Get("name")
	email := request.URL.Query().Get("email")
	phone := request.URL.Query().Get("phone")
	page, err := strconv.Atoi(request.URL.Query().Get("page"))

	if err != nil {
		// default value
		page = 0
	}

	size, err := strconv.Atoi(request.URL.Query().Get("size"))

	if err != nil {
		// default value
		size = 10
	}

	searchContactRequest := &model.SearchContactRequest{
		Name:  name,
		Phone: phone,
		Email: email,
		Page:  page,
		Size:  size,
	}

	contactResponses, totalPage := contactController.ContactService.Search(request.Context(), userToken, searchContactRequest)

	webResponse := &model.WebResponse[[]*model.ContactResponse]{
		Data: contactResponses,
		Paging: &model.PagingResponse{
			CurrentPage: page,
			TotalPage:   totalPage,
			Size:        size,
		},
	}

	helper.WriteToResponseBody(writer, webResponse)
}
