package service

import (
	"github.com/webtech-fmi/phonebook/backend/go/contact-service/pkg/domain"
	"github.com/webtech-fmi/phonebook/backend/go/infrastructure/log"
)

type ContactService struct {
	Repository domain.Repository
	Logger     *log.Logger
}

func (s *ContactService) CreateContact(payload domain.ContactPayload) error {
	contact, err := payload.ToContact()
	if err != nil {
		return err
	}

	return s.Repository.Add(*contact)
}

func (s *ContactService) GetFavouritesByOwnerID(ID string) ([]domain.Contact, error) {
	return s.Repository.GetFavouritesByOwnerID(ID)
}

func (s *ContactService) GetByOwnerID(ID string) ([]domain.Contact, error) {
	return s.Repository.GetByOwnerID(ID)
}

func (s *ContactService) GetByID(ID string) (*domain.Contact, error) {
	return s.Repository.GetByID(ID)
}

func (s *ContactService) EditContact(id string, owner_id string, payload domain.ContactPayload) error {
	contact, err := payload.ToContact()
	if err != nil {
		return err
	}

	return s.Repository.Edit(id, owner_id, *contact)
}

func (s *ContactService) Favourite(id, owner_id string) error {
	return s.Repository.AddToFavourites(id, owner_id)
}

func (s *ContactService) Unfavourite(id, owner_id string) error {
	return s.Repository.RemoveFromFavourites(id, owner_id)
}

func (s *ContactService) DeleteContact(id, owner_id string) error {
	return s.Repository.Delete(id, owner_id)
}


func (s *ContactService) MergeContacts(id, owner_id string, to_merge []string) error {
	newContact, err := s.GetByID(id)
	if err != nil {
		return err
	}

	uniqueMapPhones := uniqueMapContact{}
	uniquePhones := uniqueContact{}

	uniqueMapEmails := uniqueMapContact{}
	uniqueEmails := uniqueContact{}

	err = fillUniqueForContact(&uniqueMapPhones, &uniquePhones, "phone", newContact)
	if err != nil {
		return err
	}

	err = fillUniqueForContact(&uniqueMapEmails, &uniqueEmails, "email", newContact)
	if err != nil {
		return err
	}

	for _, currentID := range to_merge {
		currentContact, err := s.GetByID(currentID)
		if err != nil {
			return err
		}

		err = fillUniqueForContact(&uniqueMapPhones, &uniquePhones, "phone", currentContact)
		if err != nil {
			return err
		}

		err = fillUniqueForContact(&uniqueMapEmails, &uniqueEmails, "email", currentContact)
		if err != nil {
			return err
		}
	}

	// metadata & personal - no

	newContact.Email = domain.ContactInfo{
		Work:     uniqueEmails.Work,
		Home:     uniqueEmails.Home,
		Personal: uniqueEmails.Personal,
	}

	newContact.Phone = domain.ContactInfo{
		Work:     uniquePhones.Work,
		Home:     uniquePhones.Home,
		Personal: uniquePhones.Personal,
	}

	return s.Repository.Merge(*newContact, to_merge)
}
