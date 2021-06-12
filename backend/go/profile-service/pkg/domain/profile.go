package domain

import (
	"github.com/google/uuid"
	"time"
)

type Repository interface {
	Add(Profile) error
}

type Profile struct {
	ID           uuid.UUID `db:"id"`
	CreatedTime  time.Time `db:"created_time"`
	ModifiedTime time.Time `db:"modified_time"`
	FullName     string    `db:"full_name"`
	BirthDate    string    `db:"birth_date"`
}
