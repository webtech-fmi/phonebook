package profile

import (
	"github.com/webtech-fmi/phonebook/backend/go/profile-service/pkg/domain"
	"github.com/webtech-fmi/phonebook/backend/go/profile-service/pkg/service"
)

// FetchResponse is the shape of data for a loaded demo record
type FetchResponse struct {
	domain.Profile
}

// CreateResponse contains the ID post demo creation
type CreateResponse struct {
	ID string
}

// NewFetchResponse instantiate a new response post load
func NewFetchResponse(d domain.Profile, _ *service.ProfileService) *FetchResponse {
	return &FetchResponse{d}
}

// NewCreateResponse instantiates a new response when demo is created
func NewCreateResponse(d domain.Profile, _ *service.ProfileService) *CreateResponse {
	return &CreateResponse{ID: "demo"}
}
