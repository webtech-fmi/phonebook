package domain

import (
	"time"
)

type SessionRepository interface {
	PutSession(Session) error
	GetSessionByID(ID string) (*Session, error)
	DeleteSession(Session) error
}

type SessionInfo struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type Session struct {
	ID          string      `json:"id"`
	CreatedTime time.Time   `json:"created_time"`
	ExpiryTime  time.Time   `json:"expiry_time"`
	Payload     SessionInfo `json:"payload"`
}

type SessionPayload interface {
	ToSession() (*SessionInfo, error)
}
