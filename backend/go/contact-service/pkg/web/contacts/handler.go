package contacts

import (
	"net/http"

	routing "github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/content"

	"github.com/webtech-fmi/phonebook/backend/go/contact-service/pkg/service"
	"github.com/webtech-fmi/phonebook/backend/go/infrastructure/log"
)

// Handler is just a route collection
type Handler struct{}

// GetByOwner - load contact by owner ID
func (h Handler) GetByOwner(logger *log.Logger, ds *service.ContactService) func(c *routing.Context) error {
	return func(c *routing.Context) error {
		ID := c.Query("id")
		if ID == "" {
			return routing.NewHTTPError(http.StatusBadRequest, "passed an empty ID")
		}

		contact, err := ds.GetByOwnerID(ID)
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

func (h Handler) CreateContact(logger *log.Logger, ds *service.ContactService) func(c *routing.Context) error {
	return func(c *routing.Context) error {
		request := CreateRequest{}

		if err := c.Read(&request); err != nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := ds.CreateContact(&request); err != nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		c.Response.WriteHeader(http.StatusOK)
		return nil
	}
}

func (h Handler) EditContact(logger *log.Logger, ds *service.ContactService) func(c *routing.Context) error {
	return func(c *routing.Context) error {
		ID := c.Query("id")
		if ID == "" {
			return routing.NewHTTPError(http.StatusBadRequest, "passed an empty ID")
		}

		ownerID := c.Query("owner_id")
		if ownerID == "" {
			return routing.NewHTTPError(http.StatusBadRequest, "passed an empty owner ID")
		}

		request := EditRequest{}

		if err := c.Read(&request); err != nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := ds.EditContact(ID, ownerID, &request); err != nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		c.Response.WriteHeader(http.StatusOK)
		return nil
	}
}

// Routes for demo create/read
func (h Handler) Routes(api *routing.RouteGroup, logger *log.Logger, s *service.ContactService) {
	api.Get("/by-owner", h.GetByOwner(logger, s))
	api.Get("/by-id", h.Get(logger, s))
	api.Post("/create", h.CreateContact(logger, s))
	api.Post("/edit", h.EditContact(logger, s))
}
