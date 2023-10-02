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

type AddressServiceImpl struct {
	repository.UserRepository
	repository.ContactRepository
	repository.AddressRepository
	*sql.DB
	*validator.Validate
}

func NewAddressService(userRepository repository.UserRepository, contactRepository repository.ContactRepository, addressRepository repository.AddressRepository, DB *sql.DB, validate *validator.Validate) service.AddressService {
	return &AddressServiceImpl{UserRepository: userRepository, ContactRepository: contactRepository, AddressRepository: addressRepository, DB: DB, Validate: validate}
}

func (addressService *AddressServiceImpl) Create(ctx context.Context, userToken string, request *model.CreateAddressRequest) *model.AddressResponse {
	err := addressService.Validate.Struct(request)
	helper.PanicIfError(err)

	user, err := addressService.UserRepository.FindByToken(ctx, addressService.DB, userToken)
	helper.PanicIfError(err)

	contact, err := addressService.ContactRepository.FindByUserAndId(ctx, addressService.DB, user, request.ContactId)
	exception.PanicNotFoundIfError(err, "contact not found")

	address := &entity.Address{
		Id: uuid.New().String(),
		Street: sql.NullString{
			String: request.Street,
			Valid:  true,
		},
		City: sql.NullString{
			String: request.City,
			Valid:  true,
		},
		Province: sql.NullString{
			String: request.Province,
			Valid:  true,
		},
		Country: request.Country,
		PostalCode: sql.NullString{
			String: request.PostalCode,
			Valid:  true,
		},
		Contact: contact,
	}

	helper.NullStringCheck(address)

	tx, err := addressService.DB.Begin()
	defer helper.CommitOrRollback(tx)

	addressService.AddressRepository.Save(ctx, tx, address)

	return addressService.toAddressResponse(address)
}

func (addressService *AddressServiceImpl) Get(ctx context.Context, userToken string, contactId string, addressId string) *model.AddressResponse {
	user, err := addressService.UserRepository.FindByToken(ctx, addressService.DB, userToken)
	helper.PanicIfError(err)

	contact, err := addressService.ContactRepository.FindByUserAndId(ctx, addressService.DB, user, contactId)
	exception.PanicNotFoundIfError(err, "contact not found")

	address, err := addressService.AddressRepository.FindByContactAndId(ctx, addressService.DB, contact, addressId)
	exception.PanicNotFoundIfError(err, "address not found")

	return addressService.toAddressResponse(address)
}

func (addressService *AddressServiceImpl) Update(ctx context.Context, userToken string, request *model.UpdateAddressRequest) *model.AddressResponse {
	err := addressService.Validate.Struct(request)
	helper.PanicIfError(err)

	user, err := addressService.UserRepository.FindByToken(ctx, addressService.DB, userToken)
	helper.PanicIfError(err)

	contact, err := addressService.ContactRepository.FindByUserAndId(ctx, addressService.DB, user, request.ContactId)
	exception.PanicNotFoundIfError(err, "contact not found")

	address, err := addressService.AddressRepository.FindByContactAndId(ctx, addressService.DB, contact, request.AddressId)
	exception.PanicNotFoundIfError(err, "address not found")

	address.Street = sql.NullString{
		String: request.Street,
		Valid:  true,
	}

	address.City = sql.NullString{
		String: request.City,
		Valid:  true,
	}

	address.Province = sql.NullString{
		String: request.Province,
		Valid:  true,
	}

	address.Country = request.Country

	address.PostalCode = sql.NullString{
		String: request.PostalCode,
		Valid:  true,
	}

	helper.NullStringCheck(address)

	tx, err := addressService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	addressService.AddressRepository.Update(ctx, tx, address)

	return addressService.toAddressResponse(address)

}

func (addressService *AddressServiceImpl) Remove(ctx context.Context, userToken string, contactId string, addressId string) {
	user, err := addressService.UserRepository.FindByToken(ctx, addressService.DB, userToken)
	helper.PanicIfError(err)

	contact, err := addressService.ContactRepository.FindByUserAndId(ctx, addressService.DB, user, contactId)
	exception.PanicNotFoundIfError(err, "contact not found")

	address, err := addressService.AddressRepository.FindByContactAndId(ctx, addressService.DB, contact, addressId)
	exception.PanicNotFoundIfError(err, "address not found")

	tx, err := addressService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	addressService.AddressRepository.Delete(ctx, tx, address)
}

func (addressService *AddressServiceImpl) List(ctx context.Context, userToken string, contactId string) []*model.AddressResponse {
	user, err := addressService.UserRepository.FindByToken(ctx, addressService.DB, userToken)
	helper.PanicIfError(err)

	contact, err := addressService.ContactRepository.FindByUserAndId(ctx, addressService.DB, user, contactId)
	exception.PanicNotFoundIfError(err, "contact not found")

	var addresses []*entity.Address = addressService.AddressRepository.FindAllByContact(ctx, addressService.DB, contact)

	var addressResponses []*model.AddressResponse = make([]*model.AddressResponse, len(addresses))
	for i, address := range addresses {
		addressResponses[i] = addressService.toAddressResponse(address)
	}

	return addressResponses
}

func (addressService *AddressServiceImpl) toAddressResponse(address *entity.Address) *model.AddressResponse {
	return &model.AddressResponse{
		Id:         address.Id,
		Street:     address.Street.String,
		City:       address.City.String,
		Province:   address.Province.String,
		Country:    address.Country,
		PostalCode: address.PostalCode.String,
	}
}
