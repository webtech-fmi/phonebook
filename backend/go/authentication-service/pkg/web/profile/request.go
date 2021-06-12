package profile

import (
	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	// "github.com/go-ozzo/ozzo-validation/v4/is"
	// "github.com/webtech-fmi/phonebook/backend/go/domain/vocabulary"
)

type CreateRequest struct {
	Name string
	Date string
}

func (cr *CreateRequest) Validate() error {
	return ozzo.ValidateStruct(
		cr,
		ozzo.Field(&cr.Name, ozzo.Required),
		ozzo.Field(&cr.Date, ozzo.Required),
	)
}
