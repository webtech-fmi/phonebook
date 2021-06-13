package users

import (
	"net/http"

	"github.com/webtech-fmi/phonebook/backend/go/authentication-service/pkg/service"
	"github.com/webtech-fmi/phonebook/backend/go/authentication-service/pkg/domain"

	routing "github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/content"
	"github.com/webtech-fmi/phonebook/backend/go/infrastructure/log"
)

const (
	ErrCreateUserParam = "create_user_param"
)

// Handler is just a route collection
type Handler struct{}

func (h Handler) GetUserByID(logger *log.Logger, ds *service.UserService) func(c *routing.Context) error {
	return func(c *routing.Context) error {
		ID := c.Query("id")
		if ID == "" {
			return routing.NewHTTPError(http.StatusBadRequest, "passed an empty ID")
		}

		user, err := ds.GetByID(ID)
		if err != nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		w := content.JSONDataWriter{}
		return w.Write(c.Response, NewFetchResponse(*user, ds))
	}
}

func (h Handler) GetUserByCredentials(logger *log.Logger, ds *service.UserService) func(c *routing.Context) error {
	return func(c *routing.Context) error {
		request := domain.Credentials{}

		if err := c.Read(&request); err != nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := request.Validate(); err!=nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		user, err := ds.GetByCredentials(request)
		if err != nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		w := content.JSONDataWriter{}
		return w.Write(c.Response, NewFetchResponse(*user, ds))
	}
}

func (h Handler) CreateUser(logger *log.Logger, ds *service.UserService) func(c *routing.Context) error {
	return func(c *routing.Context) error {
		request := CreateRequest{}

		if err := c.Read(&request); err != nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		_, err := ds.CreateUser(&request)
		if err != nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		c.Response.WriteHeader(http.StatusOK)
		return nil
	}
}

// Routes for demo create/read
func (h Handler) Routes(api *routing.RouteGroup, logger *log.Logger, s *service.UserService) {
	api.Get("/get", h.GetUserByID(logger, s))
	api.Post("/get", h.GetUserByCredentials(logger, s))
	api.Post("/create", h.CreateUser(logger, s))
}
