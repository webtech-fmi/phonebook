package service

import (
	"errors"
	"time"

	"github.com/webtech-fmi/phonebook/backend/go/authentication-service/pkg/domain"
	"github.com/webtech-fmi/phonebook/backend/go/domain/auth"
	"github.com/webtech-fmi/phonebook/backend/go/infrastructure/log"
)

type SessionService struct {
	Repository domain.SessionRepository
	Logger     *log.Logger
	duration   uint64
}

func (s *SessionService) Start(payload domain.SessionPayload) (*domain.Session, error) {
	sessionID, err := auth.CreateSessionID()
	if err != nil {
		return nil, err
	}

	sessionInfo, err := payload.ToSession()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()

	session := domain.Session{
		ID:          sessionID,
		CreatedTime: now,
		ExpiryTime:  now.Add(time.Duration(s.duration) * time.Second),
		Payload:     *sessionInfo,
	}

	err = s.Repository.PutSession(session)
	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (s *SessionService) Valid(id string) (*domain.Session, error) {
	session, err := s.Repository.GetSessionByID(id)
	if err != nil {
		return nil, err
	}

	if time.Now().UTC().After(session.ExpiryTime) {
		err := s.Terminate(id)
		if err != nil {
			return nil, err
		}
		return nil, errors.New("session has expired")
	}

	now := time.Now().UTC()
	session.ExpiryTime = now.Add(time.Duration(s.duration) * time.Second)

	err = s.Repository.PutSession(*session)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (s *SessionService) Terminate(id string) error {
	session, err := s.Repository.GetSessionByID(id)
	if err != nil {
		return err
	}

	return s.Repository.DeleteSession(*session)
}

func NewService(repository domain.SessionRepository, options map[string]interface{}) (*SessionService, error) {
	duration, ok := options["duration"].(int) // in seconds
	if !ok {
		return nil, errors.New("missing duration")
	}

	return &SessionService{
		Repository: repository,
		duration:   uint64(duration),
	}, nil
}
