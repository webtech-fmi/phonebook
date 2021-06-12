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
