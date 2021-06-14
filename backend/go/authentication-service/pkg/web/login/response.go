package login

import (
	"github.com/webtech-fmi/phonebook/backend/go/authentication-service/pkg/domain"
)

type Response struct {
	domain.Session
}

// NewLoginResponse function
func NewLoginResponse(s domain.Session) *Response {
	return &Response{s}
}
