package contacts

import (
	"time"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"github.com/webtech-fmi/phonebook/backend/go/contact-service/pkg/domain"
	"github.com/webtech-fmi/phonebook/backend/go/domain/vocabulary"
)

type CreateRequest struct {
	OwnerID  string             `json:"owner_id"`
	Email    domain.ContactInfo `json:"email"`
	Phone    domain.ContactInfo `json:"phone"`
	Personal domain.Personal    `json:"personal"`
	Metadata domain.Metadata    `json:"metadata"`
}

func (cr *CreateRequest) Validate() error {
	return ozzo.ValidateStruct(
		cr,
		ozzo.Field(&cr.OwnerID, ozzo.Required, is.UUID),
		ozzo.Field(&cr.Email, ozzo.Required, ozzo.Each(is.Email)),
		ozzo.Field(&cr.Phone, ozzo.Required),
	)
}

func (cr *CreateRequest) ToContact() (*domain.Contact, error) {
	now := time.Now().UTC()

	return &domain.Contact{
		ID:           uuid.New(),
		OwnerID:      uuid.MustParse(cr.OwnerID),
		CreatedTime:  &now,
		ModifiedTime: &now,
		Email:        cr.Email,
		Status:       vocabulary.Active,
		Phone:        cr.Phone,
		Personal:     cr.Personal,
		Metadata:     cr.Metadata,
	}, nil
}

type EditRequest struct {
	Email    domain.ContactInfo `json:"email"`
	Phone    domain.ContactInfo `json:"phone"`
	Personal domain.Personal    `json:"personal"`
	Metadata domain.Metadata    `json:"metadata"`
}

func (cr *EditRequest) ToContact() (*domain.Contact, error) {
	now := time.Now().UTC()

	return &domain.Contact{
		ModifiedTime: &now,
		Email:        cr.Email,
		Phone:        cr.Phone,
		Personal:     cr.Personal,
		Metadata:     cr.Metadata,
	}, nil
}
