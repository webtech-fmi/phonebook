package service

import (
	"github.com/webtech-fmi/phonebook/backend/go/infrastructure/log"
	"github.com/webtech-fmi/phonebook/backend/go/profile-service/pkg/domain"
)

type ProfileService struct {
	Repository domain.Repository
	Logger     *log.Logger
}

func (s *ProfileService) CreateProfile(payload domain.ProfilePayload) error {
	profile, err := payload.ToProfile()
	if err != nil {
		return err
	}

	return s.Repository.Add(*profile)
}

func (s *ProfileService) GetByOwnerID(ID string) (*domain.Profile, error) {
	return s.Repository.GetByOwnerID(ID)
}

func (s *ProfileService) GetByID(ID string) (*domain.Profile, error) {
	return s.Repository.GetByID(ID)
}

func (s *ProfileService) EditProfile(id string, payload domain.ProfilePayload) error {
	profile, err := payload.ToProfile()
	if err != nil {
		return err
	}

	return s.Repository.Edit(id, *profile)
}
