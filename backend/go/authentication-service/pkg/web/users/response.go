package users

import (
	"github.com/webtech-fmi/phonebook/backend/go/authentication-service/pkg/domain"
	"github.com/webtech-fmi/phonebook/backend/go/authentication-service/pkg/service"
)

type FetchResponse struct {
	domain.User
}

type CreateResponse struct {
	ID string `json:"id"`
}

func NewFetchResponse(d domain.User, _ *service.UserService) *FetchResponse {
	return &FetchResponse{d}
}

func NewCreateResponse(id string, _ *service.UserService) *CreateResponse {
	return &CreateResponse{ID: id}
}

type StatusResponse struct {
	Status string `json:"status"`
}

func NewStatusResponse(status string) *StatusResponse {
	return &StatusResponse{
		Status: status,
	}
}
