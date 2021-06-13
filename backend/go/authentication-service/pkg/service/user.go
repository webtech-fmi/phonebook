package service

import (
	"time"

	"github.com/google/uuid"

	"github.com/webtech-fmi/phonebook/backend/go/authentication-service/pkg/domain"
	"github.com/webtech-fmi/phonebook/backend/go/domain/auth"
	"github.com/webtech-fmi/phonebook/backend/go/infrastructure/log"
)

type UserService struct {
	Repository domain.Repository
	Logger     *log.Logger
}

func (s UserService) CreateUser(payload domain.UserPayload) (string, error) {
	newUser, err := payload.ToUser()
	if err != nil {
		return "", err
	}

	createdTime := time.Now().UTC()
	newUser.ID = uuid.New()
	newUser.CreatedTime = &createdTime
	newUser.ModifiedTime = &createdTime

	hashedPassword, err := auth.HashPassword(newUser.Password)
	if err != nil {
		return "", err
	}
	newUser.Password = hashedPassword

	err = s.Repository.Add(*newUser)
	if err != nil {
		return "", err
	}

	return newUser.ID.String(), nil
}

func (s UserService) GetByID(id string) (*domain.User, error) {
	return s.Repository.GetUserByID(id)
}

func (s UserService) GetByCredentials(credentials domain.Credentials) (*domain.User, error) {
	return s.Repository.GetUserByCredentials(credentials)
}

func (s UserService) SetPassword(id, newPassword string) error {
	user, err := s.Repository.GetUserByID(id)
	if err != nil {
		return err
	}

	hashedPassword, err := auth.HashPassword(newPassword)
	if err != nil {
		return err
	}

	return s.Repository.SetPassword(user.ID.String(), hashedPassword)
}

func (s UserService) LockUser(id string, payload domain.LockPayload) error {
	user, err := s.Repository.GetUserByID(id)
	if err != nil {
		return err
	}

	lock, err := payload.ToLock()
	if err != nil {
		return err
	}

	return s.Repository.SetLock(user.ID.String(), lock)
}
