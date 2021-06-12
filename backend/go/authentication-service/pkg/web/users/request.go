package users

import (
	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	// "github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/webtech-fmi/phonebook/backend/go/authentication-service/pkg/domain"
)

type CreateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
}

func (cr *CreateRequest) Validate() error {
	return ozzo.ValidateStruct(
		cr,
		ozzo.Field(&cr.FullName, ozzo.Required),
		ozzo.Field(&cr.Email, ozzo.Required),
		ozzo.Field(&cr.Password, ozzo.Required),
	)
}

func (cr *CreateRequest) ToUser() (*domain.User, error) {
	return &domain.User{
		Email:    cr.Email,
		Password: cr.Password,
		FullName: cr.FullName,
		Metadata: domain.Metadata{},
	}, nil
}
