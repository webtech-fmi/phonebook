package storage

import (
	"context"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/webtech-fmi/phonebook/backend/go/authentication-service/pkg/domain"
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
		Insert("users", dbx.Params{
			"full_name":  p.FullName,
			"birth_date": p.BirthDate,
		}).Execute()
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
