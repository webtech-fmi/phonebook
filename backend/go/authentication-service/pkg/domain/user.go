package domain

import (
	"time"

	"github.com/google/uuid"
)

type Repository interface {
	Add(User) error
}

type User struct {
	ID           uuid.UUID `db:"id"`
	CreatedTime  time.Time `db:"created_time"`
	ModifiedTime time.Time `db:"modified_time"`
	Email        string    `db:"email"`
	FullName     string    `db:"full_name"`
	Password     string    `db:"password"`
}

// UserPayload interface
type UserPayload interface {
	ToUser() (*User, error)
}
