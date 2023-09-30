package impl

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"mhmmdnaufall/go-restful-api/entity"
	"mhmmdnaufall/go-restful-api/exception"
	"mhmmdnaufall/go-restful-api/helper"
	"mhmmdnaufall/go-restful-api/model"
	"mhmmdnaufall/go-restful-api/repository"
	"mhmmdnaufall/go-restful-api/service"
)

type ContactServiceImpl struct {
	repository.UserRepository
	repository.ContactRepository
	*sql.DB
	*validator.Validate
}

func NewContactService(userRepository repository.UserRepository, contactRepository repository.ContactRepository, DB *sql.DB, validate *validator.Validate) service.ContactService {
	return &ContactServiceImpl{UserRepository: userRepository, ContactRepository: contactRepository, DB: DB, Validate: validate}
}

func (contactService *ContactServiceImpl) Create(ctx context.Context, userToken string, request *model.CreateContactRequest) *model.ContactResponse {
	err := contactService.Validate.Struct(request)
	helper.PanicIfError(err)

	user, err := contactService.UserRepository.FindByToken(ctx, contactService.DB, userToken)
	helper.PanicIfError(err)

	contact := &entity.Contact{
		Id:        uuid.New().String(),
		FirstName: request.FirstName,
		LastName: sql.NullString{
			String: request.LastName,
			Valid:  true,
		},
		Phone: sql.NullString{
			String: request.Phone,
			Valid:  true,
		},
		Email: sql.NullString{
			String: request.Email,
			Valid:  true,
		},
		User: user,
	}

	tx, err := contactService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	contactService.ContactRepository.Save(ctx, tx, contact)

	return contactService.toContactResponse(contact)
}

func (contactService *ContactServiceImpl) Get(ctx context.Context, userToken string, id string) *model.ContactResponse {
	user, err := contactService.UserRepository.FindByToken(ctx, contactService.DB, userToken)
	helper.PanicIfError(err)

	contact, err := contactService.ContactRepository.FindByUserAndId(ctx, contactService.DB, user, id)
	helper.PanicIfError(err)

	return contactService.toContactResponse(contact)
}

func (contactService *ContactServiceImpl) Update(ctx context.Context, userToken string, request *model.UpdateContactRequest) *model.ContactResponse {
	err := contactService.Validate.Struct(request)
	helper.PanicIfError(err)

	user, err := contactService.UserRepository.FindByToken(ctx, contactService.DB, userToken)
	helper.PanicIfError(err)

	contact, err := contactService.ContactRepository.FindByUserAndId(ctx, contactService.DB, user, request.Id)
	helper.PanicIfError(err)

	contact.FirstName = request.FirstName
	contact.LastName = sql.NullString{
		String: request.LastName,
		Valid:  true,
	}
	contact.Phone = sql.NullString{
		String: request.Phone,
		Valid:  true,
	}
	contact.Email = sql.NullString{
		String: request.Email,
		Valid:  true,
	}

	tx, err := contactService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	contactService.ContactRepository.Update(ctx, tx, contact)

	return contactService.toContactResponse(contact)
}

func (contactService *ContactServiceImpl) Delete(ctx context.Context, userToken string, contactId string) {
	user, err := contactService.UserRepository.FindByToken(ctx, contactService.DB, userToken)
	helper.PanicIfError(err)

	contact, err := contactService.FindByUserAndId(ctx, contactService.DB, user, contactId)
	if err != nil {
		panic(exception.NewNotFoundError("Contact Not Found"))
	}

	tx, err := contactService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	contactService.ContactRepository.Delete(ctx, tx, contact)
}

func (contactService *ContactServiceImpl) Search(ctx context.Context, userToken string, request *model.SearchContactRequest) ([]*model.ContactResponse, int) {
	user, err := contactService.UserRepository.FindByToken(ctx, contactService.DB, userToken)
	helper.PanicIfError(err)

	contacts, totalPage := contactService.ContactRepository.Search(ctx, contactService.DB, user, request)

	var contactResponses []*model.ContactResponse

	for _, contact := range contacts {
		contactResponses = append(contactResponses, contactService.toContactResponse(contact))
	}

	return contactResponses, totalPage
}

func (contactService *ContactServiceImpl) toContactResponse(contact *entity.Contact) *model.ContactResponse {
	return &model.ContactResponse{
		Id:        contact.Id,
		FirstName: contact.FirstName,
		LastName:  contact.LastName.String,
		Email:     contact.Email.String,
		Phone:     contact.Phone.String,
	}
}
