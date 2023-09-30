package service

import (
	"context"
	"mhmmdnaufall/go-restful-api/model"
)

type ContactService interface {
	Create(ctx context.Context, userToken string, request *model.CreateContactRequest) *model.ContactResponse

	Get(ctx context.Context, userToken string, id string) *model.ContactResponse

	Update(ctx context.Context, userToken string, request *model.UpdateContactRequest) *model.ContactResponse

	Delete(ctx context.Context, userToken string, contactId string)

	Search(ctx context.Context, userToken string, request *model.SearchContactRequest) ([]*model.ContactResponse, int)
}
