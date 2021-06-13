package logout

import (
	"net/http"

	"github.com/webtech-fmi/phonebook/backend/go/authentication-service/pkg/service"

	routing "github.com/go-ozzo/ozzo-routing"
	"github.com/webtech-fmi/phonebook/backend/go/infrastructure/log"
)

// Handler is just a route collection
type Handler struct{}

func (h Handler) Logout(logger *log.Logger, ss *service.SessionService) func(c *routing.Context) error {
	return func(c *routing.Context) error {
		request := Request{}

		if err := c.Read(&request); err != nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := request.Validate(); err != nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		err := ss.Terminate(request.SessionID)
		if err != nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		c.Response.WriteHeader(http.StatusOK)
		return nil
	}
}

// Routes for demo create/read
func (h Handler) Routes(api *routing.RouteGroup, logger *log.Logger, ss *service.SessionService) {
	api.Post("/", h.Logout(logger, ss))
}
