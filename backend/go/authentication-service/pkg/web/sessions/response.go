package sessions

import (
	"github.com/webtech-fmi/phonebook/backend/go/authentication-service/pkg/domain"
)

type SessionValidResponse struct {
	domain.Session
}

// NewSessionValidResponse function
func NewSessionValidResponse(s domain.Session) *SessionValidResponse {
	return &SessionValidResponse{s}
}
