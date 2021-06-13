package profile

import (
	"time"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"github.com/webtech-fmi/phonebook/backend/go/profile-service/pkg/domain"
)

type CreateRequest struct {
	UserID   string `json:"user_id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func (cr *CreateRequest) Validate() error {
	return ozzo.ValidateStruct(
		cr,
		ozzo.Field(&cr.UserID, ozzo.Required, is.UUID),
		ozzo.Field(&cr.FullName, ozzo.Required),
		ozzo.Field(&cr.Email, ozzo.Required, is.Email),
		ozzo.Field(&cr.Phone, ozzo.Required),
	)
}

func (cr *CreateRequest) ToProfile() (*domain.Profile, error) {
	now := time.Now().UTC()

	return &domain.Profile{
		ID:           uuid.New(),
		UserID:       uuid.MustParse(cr.UserID),
		CreatedTime:  &now,
		ModifiedTime: &now,
		Email: domain.Contact{
			Personal: []string{
				cr.Email,
			},
		},
		Phone: domain.Contact{
			Personal: []string{
				cr.Phone,
			},
		},
		Personal: domain.Personal{
			FullName: cr.FullName,
		},
	}, nil
}

type EditRequest struct {
	Email    domain.Contact  `json:"email"`
	Phone    domain.Contact  `json:"phone"`
	Personal domain.Personal `json:"personal"`
	Metadata domain.Metadata `json:"metadata"`
}

func (cr *EditRequest) ToProfile() (*domain.Profile, error) {
	now := time.Now().UTC()

	return &domain.Profile{
		ModifiedTime: &now,
		Email:        cr.Email,
		Phone:        cr.Phone,
		Personal:     cr.Personal,
		Metadata:     cr.Metadata,
	}, nil
}
