package service

import (
	"context"
	"mhmmdnaufall/go-restful-api/model"
)

type AddressService interface {
	Create(ctx context.Context, userToken string, request *model.CreateAddressRequest) *model.AddressResponse

	Get(ctx context.Context, userToken string, contactId string, addressId string) *model.AddressResponse

	Update(ctx context.Context, userToken string, request *model.UpdateAddressRequest) *model.AddressResponse

	Remove(ctx context.Context, userToken string, contactId string, addressId string)

	List(ctx context.Context, userToken string, contactId string) []*model.AddressResponse
}
