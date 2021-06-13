package domain

import (
	"database/sql/driver"
	"time"

	"github.com/google/uuid"
	"github.com/webtech-fmi/phonebook/backend/go/infrastructure/storage"
)

type Repository interface {
	Add(Profile) error
	GetByID(string) (*Profile, error)
	GetByOwnerID(string) (*Profile, error)
	Edit(string, Profile) error
}

type Profile struct {
	ID           uuid.UUID  `db:"id"`
	UserID       uuid.UUID  `db:"user_id"`
	CreatedTime  *time.Time `db:"created_time"`
	ModifiedTime *time.Time `db:"modified_time"`
	Email        Contact    `db:"email"`
	Personal     Personal   `db:"personal"`
	Phone        Contact    `db:"phone"`
	Metadata     Metadata   `db:"metadata"`
}

type ProfilePayload interface {
	ToProfile() (*Profile, error)
}

type Contact struct {
	Home     []string `json:"home"`
	Work     []string `json:"work"`
	Personal []string `json:"personal"`
}

// SQL Valuer/Scan interface implementations
func (e Contact) Value() (driver.Value, error) {
	return storage.JSONBValue(e)
}

// SQL Valuer/Scan interface implementations
func (e *Contact) Scan(src interface{}) error {
	return storage.ScanJSONB(e, src)
}

type Personal struct {
	FullName  string `json:"full_name"`
	BirthDate string `json:"birth_date"`
}

// SQL Valuer/Scan interface implementations
func (p Personal) Value() (driver.Value, error) {
	return storage.JSONBValue(p)
}

// SQL Valuer/Scan interface implementations
func (p *Personal) Scan(src interface{}) error {
	return storage.ScanJSONB(p, src)
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
