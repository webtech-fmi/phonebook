package users

import (
	"time"

	"github.com/google/uuid"
	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/webtech-fmi/phonebook/backend/go/authentication-service/pkg/domain"
	"strings"
)

type CreateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
}

func (cr *CreateRequest) Validate() error {
	return ozzo.ValidateStruct(
		cr,
		ozzo.Field(&cr.FullName, ozzo.Required, ozzo.Length(5, 25)),
		ozzo.Field(&cr.Email, ozzo.Required, is.Email),
		ozzo.Field(&cr.Password, ozzo.Required, ozzo.Length(8, 32)),
	)
}

func (cr *CreateRequest) ToUser() (*domain.User, error) {
	return &domain.User{
		Email:    strings.ToLower(cr.Email),
		Password: cr.Password,
		FullName: cr.FullName,
		Metadata: domain.Metadata{},
	}, nil
}

type LockRequest struct {
	Reason string `json:"reason"`
}

func (r *LockRequest) ToLock() (*domain.Lock, error) {
	now := time.Now().UTC()
	return &domain.Lock{
		CreatedTime: &now,
		Code:        uuid.New().String(),
		Reason:      r.Reason,
	}, nil
}
