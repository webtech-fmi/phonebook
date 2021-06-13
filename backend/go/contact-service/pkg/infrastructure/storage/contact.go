package storage

import (
	"context"
	"time"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/google/uuid"
	"github.com/webtech-fmi/phonebook/backend/go/contact-service/pkg/domain"
	"github.com/webtech-fmi/phonebook/backend/go/domain/vocabulary"
	"github.com/webtech-fmi/phonebook/backend/go/infrastructure/log"
	"github.com/webtech-fmi/phonebook/backend/go/infrastructure/storage"
)

const (
	contactsTable = "contacts"
)

type ContactRepository struct {
	Adapter *storage.PSQL
}

func (r *ContactRepository) Add(p domain.Contact) error {
	_, err := r.Adapter.DB.
		Insert(
			contactsTable,
			dbx.Params{
				"id":            p.ID,
				"status":        string(p.Status),
				"owner_id":      p.OwnerID,
				"favourite":     p.Favourite,
				"created_time":  p.CreatedTime,
				"modified_time": p.ModifiedTime,
				"email":         p.Email,
				"personal":      p.Personal,
				"phone":         p.Phone,
				"metadata":      p.Metadata,
			}).Execute()
	return err
}

func (r *ContactRepository) GetByID(id string) (*domain.Contact, error) {
	ID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	query := r.Adapter.DB.
		Select(
			"id",
			"status",
			"owner_id",
			"favourite",
			"created_time",
			"modified_time",
			"email",
			"personal",
			"phone",
			"metadata",
		).
		From(contactsTable).
		Where(
			dbx.And(
				dbx.In("id", ID),
				dbx.In("status", string(vocabulary.Active)),
			),
		)

	var contact domain.Contact
	err = query.One(&contact)
	if err != nil {
		return nil, err
	}

	return &contact, nil
}

func (r *ContactRepository) GetByOwnerID(ownerID string) ([]domain.Contact, error) {
	ID, err := uuid.Parse(ownerID)
	if err != nil {
		return nil, err
	}

	query := r.Adapter.DB.
		Select(
			"id",
			"status",
			"owner_id",
			"favourite",
			"created_time",
			"modified_time",
			"email",
			"personal",
			"phone",
			"metadata",
		).
		From(contactsTable).
		Where(
			dbx.And(
				dbx.In("status", string(vocabulary.Active)),
				dbx.In("owner_id", ID),
			),
		)

	var contacts []domain.Contact
	err = query.All(&contacts)
	if err != nil {
		return nil, err
	}

	return contacts, nil
}

func (r *ContactRepository) Edit(id, owner_id string, newProfile domain.Contact) error {
	ID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	ownerID, err := uuid.Parse(owner_id)
	if err != nil {
		return err
	}

	_, err = r.Adapter.DB.Update(
		contactsTable,
		dbx.Params{
			"modified_time": time.Now().UTC(),
			"email":         newProfile.Email,
			"personal":      newProfile.Personal,
			"phone":         newProfile.Phone,
			"metadata":      newProfile.Metadata,
		},
		dbx.And(
			dbx.In("id", ID),
			dbx.In("owner_id", ownerID),
			dbx.In("status", string(vocabulary.Active)),
		),
	).Execute()

	return err
}

func (r *ContactRepository) editFavourites(id, owner_id string, favourite bool) error {
	ID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	ownerID, err := uuid.Parse(owner_id)
	if err != nil {
		return err
	}

	_, err = r.Adapter.DB.Update(
		contactsTable,
		dbx.Params{
			"modified_time": time.Now().UTC(),
			"favourite":     favourite,
		},
		dbx.And(
			dbx.In("id", ID),
			dbx.In("owner_id", ownerID),
			dbx.In("status", string(vocabulary.Active)),
		),
	).Execute()

	return err
}

func (r *ContactRepository) AddToFavourites(id, owner_id string) error {
	return r.editFavourites(id, owner_id, true)
}

func (r *ContactRepository) RemoveFromFavourites(id, owner_id string) error {
	return r.editFavourites(id, owner_id, false)
}

func (r *ContactRepository) Delete(id, owner_id string) error {
	ID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	ownerID, err := uuid.Parse(owner_id)
	if err != nil {
		return err
	}

	_, err = r.Adapter.DB.Update(
		contactsTable,
		dbx.Params{
			"modified_time": time.Now().UTC(),
			"status":        string(vocabulary.Active),
		},
		dbx.And(
			dbx.In("id", ID),
			dbx.In("owner_id", ownerID),
			dbx.In("status", string(vocabulary.Active)),
		),
	).Execute()

	return err
}

func (r *ContactRepository) Merge(toMerge domain.Contact, toDelete []string) error {
	return r.Adapter.DB.Transactional(func(tx *dbx.Tx) error {
		var err error

		_, err = tx.Update(
			contactsTable,
			dbx.Params{
				"modified_time": time.Now().UTC(),
				"email":         toMerge.Email,
				"personal":      toMerge.Personal,
				"phone":         toMerge.Phone,
				"metadata":      toMerge.Metadata,
			},
			dbx.And(
				dbx.In("id", toMerge.ID),
				dbx.In("status", string(vocabulary.Active)),
			),
		).Execute()
		if err != nil {
			return err
		}

		for _, current := range toDelete {
			_, err = tx.Update(
				contactsTable,
				dbx.Params{
					"status": string(vocabulary.Archived),
				},
				dbx.And(
					dbx.In("id", current),
				),
			).Execute()
			if err != nil {
				return err
			}
		}
		return err
	})
}

// NewRepository creates a PSQL implementation of a secrets repository
func NewRepository(ctx context.Context, options map[string]interface{}, logger *log.Logger) (*ContactRepository, error) {
	adapter, err := storage.NewPSQL(ctx, options, logger)

	if err != nil {
		return nil, err
	}

	return &ContactRepository{
		Adapter: adapter,
	}, nil
}
