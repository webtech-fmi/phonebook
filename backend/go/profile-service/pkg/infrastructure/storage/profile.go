package storage

import (
	"context"
	"time"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/google/uuid"
	"github.com/webtech-fmi/phonebook/backend/go/infrastructure/log"
	"github.com/webtech-fmi/phonebook/backend/go/infrastructure/storage"
	"github.com/webtech-fmi/phonebook/backend/go/profile-service/pkg/domain"
)

const (
	profilesTable = "profiles"
)

// GetByID(string) (*Profile, error)
// GetByOwnerID(string) (*Profile, error)
// Edit(Profile) error

type ProfileRepository struct {
	Adapter *storage.PSQL
}

func (r *ProfileRepository) Add(p domain.Profile) error {
	_, err := r.Adapter.DB.
		Insert(
			profilesTable,
			dbx.Params{
				"id":            p.ID,
				"user_id":       p.UserID,
				"consent":       p.Consent,
				"created_time":  p.CreatedTime,
				"modified_time": p.ModifiedTime,
				"email":         p.Email,
				"personal":      p.Personal,
				"phone":         p.Phone,
				"metadata":      p.Metadata,
			}).Execute()
	return err
}

func (r *ProfileRepository) GetByID(id string) (*domain.Profile, error) {
	ID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	query := r.Adapter.DB.
		Select(
			"id",
			"user_id",
			"consent",
			"created_time",
			"modified_time",
			"email",
			"personal",
			"phone",
			"metadata",
		).
		From(profilesTable).
		Where(dbx.In("id", ID))

	var profile domain.Profile
	err = query.One(&profile)
	if err != nil {
		return nil, err
	}

	return &profile, nil
}

func (r *ProfileRepository) GetByOwnerID(ownerID string) (*domain.Profile, error) {
	ID, err := uuid.Parse(ownerID)
	if err != nil {
		return nil, err
	}

	query := r.Adapter.DB.
		Select(
			"id",
			"user_id",
			"consent",
			"created_time",
			"modified_time",
			"email",
			"personal",
			"phone",
			"metadata",
		).
		From(profilesTable).
		Where(dbx.In("user_id", ID))

	var profile domain.Profile
	err = query.One(&profile)
	if err != nil {
		return nil, err
	}

	return &profile, nil
}

func (r *ProfileRepository) Edit(id string, newProfile domain.Profile) error {
	ID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	_, err = r.Adapter.DB.Update(
		profilesTable,
		dbx.Params{
			"modified_time": time.Now().UTC(),
			"email":         newProfile.Email,
			"personal":      newProfile.Personal,
			"phone":         newProfile.Phone,
			"metadata":      newProfile.Metadata,
		},
		dbx.In("id", ID),
	).Execute()

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
