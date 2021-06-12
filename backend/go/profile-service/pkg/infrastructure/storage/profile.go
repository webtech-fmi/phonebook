package storage

import (
	"context"
	"github.com/go-ozzo/ozzo-dbx"
	"github.com/webtech-fmi/phonebook/backend/go/infrastructure/log"
	"github.com/webtech-fmi/phonebook/backend/go/infrastructure/storage"
	"github.com/webtech-fmi/phonebook/backend/go/profile-service/pkg/domain"
)

const (
	profilesTable = "profiles"
)

type ProfileRepository struct {
	Adapter *storage.PSQL
}

func (r *ProfileRepository) Add(p domain.Profile) error {
	_, err := r.Adapter.DB.
		Insert("profiles", dbx.Params{
			"full_name":  p.FullName,
			"birth_date": p.BirthDate,
		}).Execute()
	return err
}

// NewRepository creates a PSQL implementation of a secrets repository
func NewRepository(ctx context.Context, options map[string]interface{}, logger *log.Logger) (*ProfileRepository, error) {
	adapter, err := storage.NewPSQL(ctx, options, logger)

	if err != nil {
		return nil, err
	}

	return &ProfileRepository{
		Adapter: adapter,
	}, nil
}
