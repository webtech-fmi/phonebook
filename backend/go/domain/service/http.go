package service

import (
	"github.com/webtech-fmi/phonebook/backend/go/domain/service/authentication"
	"github.com/webtech-fmi/phonebook/backend/go/domain/service/profile"
)

type HTTPServices struct {
	Authentication authentication.Service
	Profile        profile.Service
}
