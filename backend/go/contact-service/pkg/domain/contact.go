package domain

import (
	"database/sql/driver"
	"time"

	"github.com/google/uuid"
	"github.com/webtech-fmi/phonebook/backend/go/domain/vocabulary"
	"github.com/webtech-fmi/phonebook/backend/go/infrastructure/storage"
)

type Repository interface {
	Add(Contact) error
	GetByID(string) (*Contact, error)
	GetByOwnerID(string) ([]Contact, error)
	GetFavouritesByOwnerID(string) ([]Contact, error)
	Edit(string, string, Contact) error
	Delete(string, string) error
	AddToFavourites(string, string) error
	RemoveFromFavourites(string, string) error
	Merge(Contact, []string) error
}

type Contact struct {
	ID           uuid.UUID         `json:"id" db:"id"`
	Status       vocabulary.Status `json:"status" db:"status"`
	OwnerID      uuid.UUID         `json:"owner_id" db:"owner_id"`
	CreatedTime  *time.Time        `json:"created_time" db:"created_time"`
	ModifiedTime *time.Time        `json:"modified_time" db:"modified_time"`
	Favourite    bool              `json:"favourite" db:"favourite"`
	Email        ContactInfo       `json:"email" db:"email"`
	Personal     Personal          `json:"personal" db:"personal"`
	Phone        ContactInfo       `json:"phone" db:"phone"`
	Metadata     Metadata          `json:"metadata" db:"metadata"`
}

type ContactPayload interface {
	ToContact() (*Contact, error)
}

type ContactInfo struct {
	Home     []string `json:"home"`
	Work     []string `json:"work"`
	Personal []string `json:"personal"`
}

// SQL Valuer/Scan interface implementations
func (e ContactInfo) Value() (driver.Value, error) {
	return storage.JSONBValue(e)
}

// SQL Valuer/Scan interface implementations
func (e *ContactInfo) Scan(src interface{}) error {
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

type Metadata struct {
	Organization string `json:"organization"`
	Address      string `json:"address"`
}

// SQL Valuer/Scan interface implementations
func (m Metadata) Value() (driver.Value, error) {
	return storage.JSONBValue(m)
}

// SQL Valuer/Scan interface implementations
func (m *Metadata) Scan(src interface{}) error {
	return storage.ScanJSONB(m, src)
}
