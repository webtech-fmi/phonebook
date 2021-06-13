package login

import (
	"github.com/webtech-fmi/phonebook/backend/go/authentication-service/pkg/domain"
)

type Response struct {
	SessionID string `json:"session_id"`
}

// NewLoginResponse function
func NewLoginResponse(s domain.Session) *Response {
	return &Response{s.ID}
}
