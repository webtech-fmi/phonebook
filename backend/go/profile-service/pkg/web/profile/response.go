package profile

import (
	"github.com/webtech-fmi/phonebook/backend/go/profile-service/pkg/domain"
	"github.com/webtech-fmi/phonebook/backend/go/profile-service/pkg/service"
)

type FetchResponse struct {
	domain.Profile
}

// NewFetchResponse instantiate a new response post load
func NewFetchResponse(d domain.Profile, _ *service.ProfileService) *FetchResponse {
	return &FetchResponse{d}
}
