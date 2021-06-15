package service

import (
	"errors"

	"github.com/webtech-fmi/phonebook/backend/go/contact-service/pkg/domain"
)

type uniqueMapContact struct {
	Home     map[string]struct{}
	Work     map[string]struct{}
	Personal map[string]struct{}
}

type uniqueContact struct {
	Home     []string
	Work     []string
	Personal []string
}

func addUnique(uniqueMap *map[string]struct{}, unique *[]string, candidates []string) {
	if unique == nil || uniqueMap == nil {
		return
	}

	if uniqueMap != nil && *uniqueMap == nil {
		*uniqueMap = make(map[string]struct{})
	}

	for _, current := range candidates {
		if _, exists := (*uniqueMap)[current]; !exists {
			(*uniqueMap)[current] = struct{}{}
			*unique = append(*unique, current)
		}
	}
}

func fillUniqueForContact(uniqueMap *uniqueMapContact, unique *uniqueContact, contactType string, contact *domain.Contact) error {
	if uniqueMap == nil || unique == nil {
		return errors.New("nil unique struct passed to fillUnique")
	}

	if contact == nil {
		return errors.New("nil contact passed to fillUnique")
	}

	switch contactType {
	case "phone":
		addUnique(&uniqueMap.Home, &unique.Home, contact.Phone.Home)
		addUnique(&uniqueMap.Personal, &unique.Personal, contact.Phone.Personal)
		addUnique(&uniqueMap.Work, &unique.Work, contact.Phone.Work)
		break
	case "email":
		addUnique(&uniqueMap.Home, &unique.Home, contact.Email.Home)
		addUnique(&uniqueMap.Personal, &unique.Personal, contact.Email.Personal)
		addUnique(&uniqueMap.Work, &unique.Work, contact.Email.Work)
		break
	default:
		return errors.New("invalid contact type")
	}
	return nil
}
