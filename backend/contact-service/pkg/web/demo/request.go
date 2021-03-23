package demo

import (
	ozzo "github.com/go-ozzo/ozzo-validation/v4"
)

// CreateRequest is the payload shape for demo creation
type CreateRequest struct {
	Label string
}

// Validate is a proxy method to confirm payload satisfies expectations
func (cr *CreateRequest) Validate() error {
	return ozzo.ValidateStruct(
		cr,
		ozzo.Field(&cr.Label, ozzo.Required),
	)
}
