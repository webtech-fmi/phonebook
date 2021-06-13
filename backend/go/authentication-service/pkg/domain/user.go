package domain

import (
	"database/sql/driver"
	"time"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/webtech-fmi/phonebook/backend/go/domain/vocabulary"
	"github.com/webtech-fmi/phonebook/backend/go/infrastructure/storage"
)

type Repository interface {
	Add(User) error
	GetUserByID(string) (*User, error)
	GetUserByCredentials(Credentials) (*User, error)
	SetPassword(string, string) error
	SetLock(string, *Lock) error
}

type Lock struct {
	Reason      string     `json:"reason"`
	CreatedTime *time.Time `json:"created_time"`
	Code        string     `json:"code"`
}

// SQL Valuer/Scan interface implementations
func (l Lock) Value() (driver.Value, error) {
	return storage.JSONBValue(l)
}

// SQL Valuer/Scan interface implementations
func (l *Lock) Scan(src interface{}) error {
	return storage.ScanJSONB(l, src)
}

type Metadata struct{}

// SQL Valuer/Scan interface implementations
func (m Metadata) Value() (driver.Value, error) {
	return storage.JSONBValue(m)
}

// SQL Valuer/Scan interface implementations
func (m *Metadata) Scan(src interface{}) error {
	return storage.ScanJSONB(m, src)
}

type User struct {
	ID           uuid.UUID  `json:"id" db:"id"`
	CreatedTime  *time.Time `json:"created_time" db:"created_time"`
	ModifiedTime *time.Time `json:"modified_time" db:"modified_time"`
	Email        string     `json:"email" db:"email"`
	FullName     string     `json:"full_name" db:"full_name"`
	Password     string     `json:"password" db:"password"`
	Lock         *Lock      `json:"lock" db:"lock"`
	Metadata     Metadata   `json:"metadata" db:"metadata"`
}

func (u *User) ToSession() (*SessionInfo, error) {
	return &SessionInfo{
		ID:    u.ID.String(),
		Email: u.Email,
	}, nil
}

// Credentials struct
type Credentials struct {
	Email  string
	Secret string
	Type   vocabulary.CredentialsType
}

func (cr *Credentials) Validate() error {
	return ozzo.ValidateStruct(
		cr,
		ozzo.Field(&cr.Email, ozzo.Required),
		ozzo.Field(&cr.Secret, ozzo.Required),
		ozzo.Field(&cr.Type, ozzo.Required, ozzo.In(
			vocabulary.CredentialsPassword,
			vocabulary.CredentialsLock,
		)),
	)
}

// UserPayload interface
type UserPayload interface {
	ToUser() (*User, error)
}

// LockPayload interface
type LockPayload interface {
	ToLock() (*Lock, error)
}

// LoginPayload interface
type LoginPayload interface {
	ToCredentials() (*Credentials, error)
}
