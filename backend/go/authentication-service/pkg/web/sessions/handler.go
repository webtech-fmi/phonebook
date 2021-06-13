package sessions

import (
	"net/http"

	"github.com/webtech-fmi/phonebook/backend/go/authentication-service/pkg/service"

	routing "github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/content"
	"github.com/webtech-fmi/phonebook/backend/go/infrastructure/log"
)

// Handler is just a route collection
type Handler struct{}

func (h Handler) Valid(logger *log.Logger, s *service.SessionService) func(c *routing.Context) error {
	return func(c *routing.Context) error {
		request := SessionValidRequest{}

		if err := c.Read(&request); err != nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := request.Validate(); err != nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		session, err := s.Valid(request.ID)
		if err != nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		w := content.JSONDataWriter{}
		return w.Write(c.Response, NewSessionValidResponse(*session))
	}
}

// Routes for demo create/read
func (h Handler) Routes(api *routing.RouteGroup, logger *log.Logger, s *service.SessionService) {
	api.Post("/valid", h.Valid(logger, s))
}
