package service

import (
	// "time"

	// "github.com/google/uuid"

	"github.com/webtech-fmi/phonebook/backend/go/infrastructure/log"
	"github.com/webtech-fmi/phonebook/backend/go/profile-service/pkg/domain"
	// "github.com/webtech-fmi/phonebook/backend/go/domain/vocabulary"
)

type ProfileService struct {
	Repository domain.Repository
	Logger     *log.Logger
}

func (s ProfileService) CreateProfile(name, date string) error {
	// // hashing password
	// now := time.Now().UTC()
	// userID := uuid.New()
	// // create lock
	// newLock := &domain.Lock{}

	return s.Repository.Add(domain.Profile{
		FullName:  name,
		BirthDate: date,
	})
}
