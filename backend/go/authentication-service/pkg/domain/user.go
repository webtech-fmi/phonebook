package domain

import (
	"database/sql/driver"
	"time"

	"github.com/google/uuid"
	"github.com/webtech-fmi/phonebook/backend/go/infrastructure/storage"
)

type Repository interface {
	Add(User) error
	GetUserByID(string) (*User, error)
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
	ID           uuid.UUID  `db:"id"`
	CreatedTime  *time.Time `db:"created_time"`
	ModifiedTime *time.Time `db:"modified_time"`
	Email        string     `db:"email"`
	FullName     string     `db:"full_name"`
	Password     string     `db:"password"`
	Lock         *Lock      `db:"lock,omitempty"`
	Metadata     Metadata   `db:"metadata"`
}

// UserPayload interface
type UserPayload interface {
	ToUser() (*User, error)
}