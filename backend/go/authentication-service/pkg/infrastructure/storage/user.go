package storage

import (
	"context"
	"fmt"
	"time"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/google/uuid"
	"github.com/webtech-fmi/phonebook/backend/go/authentication-service/pkg/domain"
	"github.com/webtech-fmi/phonebook/backend/go/domain/vocabulary"
	"github.com/webtech-fmi/phonebook/backend/go/infrastructure/log"
	"github.com/webtech-fmi/phonebook/backend/go/infrastructure/storage"
)

const (
	usersTable = "users"
)

type UserRepository struct {
	Adapter *storage.PSQL
}

func (r *UserRepository) Add(p domain.User) error {
	_, err := r.Adapter.DB.
		Insert(
			usersTable,
			dbx.Params{
				"id":            p.ID,
				"created_time":  p.CreatedTime,
				"modified_time": p.ModifiedTime,
				"email":         p.Email,
				"password":      p.Password,
				"full_name":     p.FullName,
			},
		).Execute()

	return err
}

func (r *UserRepository) GetUserByID(id string) (*domain.User, error) {
	ID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	query := r.Adapter.DB.
		Select("id", "email", "full_name", "lock").
		From(usersTable).
		Where(dbx.In("id", ID))

	var user domain.User
	err = query.One(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUserByCredentials(credentials domain.Credentials) (*domain.User, error) {
	query := r.Adapter.DB.
		Select("id", "email", "full_name", "lock").
		From(usersTable)

	switch credentials.Type {
	case vocabulary.CredentialsPassword:
		query = query.Where(
			dbx.NewExp("lock->>'reason' IS NULL AND lock->>'code' IS NULL"),
		)
	case vocabulary.CredentialsLock:
		query = query.Where(
			dbx.In("lock->>'code'", credentials.Secret),
		)
	default:
		return nil, fmt.Errorf("invalid credentials type: [%s]", credentials.Type)
	}

	var user domain.User
	err := query.One(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) SetPassword(id, password string) error {
	_, err := r.Adapter.DB.Update(
		usersTable,
		dbx.Params{
			"modified_time": time.Now().UTC(),
			"password":      password,
		},
		dbx.In("id", id),
	).Execute()

	return err
}

func (r *UserRepository) SetLock(id string, lock *domain.Lock) error {
	_, err := r.Adapter.DB.Update(
		usersTable,
		dbx.Params{
			"modified_time": time.Now().UTC(),
			"lock":          lock,
		},
		dbx.In("id", id),
	).Execute()

	return err
}

// NewRepository creates a PSQL implementation of a secrets repository
func NewRepository(ctx context.Context, options map[string]interface{}, logger *log.Logger) (*UserRepository, error) {
	adapter, err := storage.NewPSQL(ctx, options, logger)

	if err != nil {
		return nil, err
	}

	return &UserRepository{
		Adapter: adapter,
	}, nil
}
