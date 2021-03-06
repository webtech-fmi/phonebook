package contacts

import (
	"github.com/webtech-fmi/phonebook/backend/go/contact-service/pkg/domain"
	"github.com/webtech-fmi/phonebook/backend/go/contact-service/pkg/service"
)

type FetchResponse struct {
	domain.Contact
}

// NewFetchResponse instantiate a new response post load
func NewFetchResponse(d domain.Contact, _ *service.ContactService) *FetchResponse {
	return &FetchResponse{d}
}

type FetchAllResponse struct {
	Contacts []domain.Contact `json:"contacts"`
}

// NewFetchResponse instantiate a new response post load
func NewFetchAllResponse(d []domain.Contact, _ *service.ContactService) *FetchAllResponse {
	return &FetchAllResponse{
		Contacts: d,
	}
}
