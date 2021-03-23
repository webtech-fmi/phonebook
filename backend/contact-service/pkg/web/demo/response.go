package demo

import (
	"github.com/webtech-fmi/phonebook/backend/contact-service/pkg/domain"
	"github.com/webtech-fmi/phonebook/backend/contact-service/pkg/service"
)

// FetchResponse is the shape of data for a loaded demo record
type FetchResponse struct {
	domain.Demo
}

// CreateResponse contains the ID post demo creation
type CreateResponse struct {
	ID string
}

// NewFetchResponse instantiate a new response post load
func NewFetchResponse(d domain.Demo, _ *service.DemoService) *FetchResponse {
	return &FetchResponse{d}
}

// NewCreateResponse instantiates a new response when demo is created
func NewCreateResponse(d domain.Demo, _ *service.DemoService) *CreateResponse {
	return &CreateResponse{ID: "demo"}
}
