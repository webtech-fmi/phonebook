package login

import (
	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/webtech-fmi/phonebook/backend/go/authentication-service/pkg/domain"
	"github.com/webtech-fmi/phonebook/backend/go/domain/vocabulary"
)

type Request struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *Request) Validate() error {
	return ozzo.ValidateStruct(
		r,
		ozzo.Field(&r.Email, ozzo.Required, is.Email),
		ozzo.Field(&r.Password, ozzo.Required),
	)
}

func (r Request) ToCredentials() (*domain.Credentials, error) {
	return &domain.Credentials{
		Email:  r.Email,
		Secret: r.Password,
		Type:   vocabulary.CredentialsPassword,
	}, nil
}
