package profile

import (
	"net/http"

	routing "github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/content"

	"github.com/webtech-fmi/phonebook/backend/go/infrastructure/log"
	"github.com/webtech-fmi/phonebook/backend/go/profile-service/pkg/service"
)

// Handler is just a route collection
type Handler struct{}

// GetByOwner - load profile by owner ID
func (h Handler) GetByOwner(logger *log.Logger, ds *service.ProfileService) func(c *routing.Context) error {
	return func(c *routing.Context) error {
		ID := c.Query("id")
		if ID == "" {
			return routing.NewHTTPError(http.StatusBadRequest, "passed an empty ID")
		}

		profile, err := ds.GetByOwnerID(ID)
		if err != nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		w := content.JSONDataWriter{}

		return w.Write(c.Response, NewFetchResponse(*profile, ds))
	}
}

// Get - load profile by ID
func (h Handler) Get(logger *log.Logger, ds *service.ProfileService) func(c *routing.Context) error {
	return func(c *routing.Context) error {
		ID := c.Query("id")
		if ID == "" {
			return routing.NewHTTPError(http.StatusBadRequest, "passed an empty ID")
		}

		profile, err := ds.GetByID(ID)
		if err != nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		w := content.JSONDataWriter{}

		return w.Write(c.Response, NewFetchResponse(*profile, ds))
	}
}

func (h Handler) CreateProfile(logger *log.Logger, ds *service.ProfileService) func(c *routing.Context) error {
	return func(c *routing.Context) error {
		request := CreateRequest{}

		if err := c.Read(&request); err != nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := ds.CreateProfile(&request); err != nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		c.Response.WriteHeader(http.StatusOK)
		return nil
	}
}

// Routes for demo create/read
func (h Handler) Routes(api *routing.RouteGroup, logger *log.Logger, s *service.ProfileService) {
	api.Get("/by-owner", h.GetByOwner(logger, s))
	api.Get("/by-id", h.Get(logger, s))
	api.Post("/create", h.CreateProfile(logger, s))
}
