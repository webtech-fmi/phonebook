package contacts

import (
	"net/http"

	routing "github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/content"

	"github.com/webtech-fmi/phonebook/backend/go/contact-service/pkg/service"
	domain_service "github.com/webtech-fmi/phonebook/backend/go/domain/service"
	"github.com/webtech-fmi/phonebook/backend/go/infrastructure/log"
)

// Handler is just a route collection
type Handler struct{}

// GetByOwner - load contact by owner ID
func (h Handler) GetByOwner(logger *log.Logger, ds *service.ContactService, hs *domain_service.HTTPServices) func(c *routing.Context) error {
	return func(c *routing.Context) error {
		ID := c.Query("id")
		if ID == "" {
			return routing.NewHTTPError(http.StatusBadRequest, "passed an empty ID")
		}

		session, err := hs.Authentication.ResolveSessionID(ID)
		if err != nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		profile, err := hs.Profile.ResolveUserID(session.ID)
		if err != nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		contact, err := ds.GetByOwnerID(profile.ID.String())
		if err != nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		w := content.JSONDataWriter{}

		return w.Write(c.Response, NewFetchAllResponse(contact, ds))
	}
}

// Get - load contact by ID
func (h Handler) Get(logger *log.Logger, ds *service.ContactService) func(c *routing.Context) error {
	return func(c *routing.Context) error {
		ID := c.Query("id")
		if ID == "" {
			return routing.NewHTTPError(http.StatusBadRequest, "passed an empty ID")
		}

		contact, err := ds.GetByID(ID)
		if err != nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		w := content.JSONDataWriter{}

		return w.Write(c.Response, NewFetchResponse(*contact, ds))
	}
}

func (h Handler) CreateContact(logger *log.Logger, ds *service.ContactService, hs *domain_service.HTTPServices) func(c *routing.Context) error {
	return func(c *routing.Context) error {
		ID := c.Query("id")
		if ID == "" {
			return routing.NewHTTPError(http.StatusBadRequest, "passed an empty ID")
		}

		session, err := hs.Authentication.ResolveSessionID(ID)
		if err != nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		profile, err := hs.Profile.ResolveUserID(session.ID)
		if err != nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		request := CreateRequest{}

		if err := c.Read(&request); err != nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if request.OwnerID != profile.ID.String() {
			return routing.NewHTTPError(http.StatusBadRequest, "unauthorized")
		}

		if err := ds.CreateContact(&request); err != nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		c.Response.WriteHeader(http.StatusOK)
		return nil
	}
}

func (h Handler) EditContact(logger *log.Logger, ds *service.ContactService, hs *domain_service.HTTPServices) func(c *routing.Context) error {
	return func(c *routing.Context) error {
		ID := c.Query("id")
		if ID == "" {
			return routing.NewHTTPError(http.StatusBadRequest, "passed an empty ID")
		}

		session, err := hs.Authentication.ResolveSessionID(ID)
		if err != nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		profile, err := hs.Profile.ResolveUserID(session.ID)
		if err != nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		request := EditRequest{}

		if err := c.Read(&request); err != nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := ds.EditContact(request.ID, profile.ID.String(), &request); err != nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		c.Response.WriteHeader(http.StatusOK)
		return nil
	}
}

func (h Handler) MergeContacts(logger *log.Logger, ds *service.ContactService, hs *domain_service.HTTPServices) func(c *routing.Context) error {
	return func(c *routing.Context) error {
		ID := c.Query("id")
		if ID == "" {
			return routing.NewHTTPError(http.StatusBadRequest, "passed an empty ID")
		}

		session, err := hs.Authentication.ResolveSessionID(ID)
		if err != nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		profile, err := hs.Profile.ResolveUserID(session.ID)
		if err != nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		request := MergeRequest{}
		if err := c.Read(&request); err != nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := ds.MergeContacts(ID, profile.ID.String(), request.Contacts); err != nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		c.Response.WriteHeader(http.StatusOK)
		return nil
	}
}

// Routes for demo create/read
func (h Handler) Routes(api *routing.RouteGroup, logger *log.Logger, s *service.ContactService, hs *domain_service.HTTPServices) {
	api.Get("/by-owner", h.GetByOwner(logger, s, hs))
	api.Get("/by-id", h.Get(logger, s))
	api.Post("/create", h.CreateContact(logger, s, hs))
	api.Post("/edit", h.EditContact(logger, s, hs))
	api.Post("/merge", h.MergeContacts(logger, s, hs))
}
