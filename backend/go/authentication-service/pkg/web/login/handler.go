package login

import (
	"net/http"

	"github.com/webtech-fmi/phonebook/backend/go/authentication-service/pkg/service"

	routing "github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/content"
	"github.com/webtech-fmi/phonebook/backend/go/infrastructure/log"
)

// Handler is just a route collection
type Handler struct{}

func (h Handler) Login(logger *log.Logger, us *service.UserService, ss *service.SessionService) func(c *routing.Context) error {
	return func(c *routing.Context) error {
		request := Request{}

		if err := c.Read(&request); err != nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := request.Validate(); err != nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		user, err := us.Login(request)
		if err != nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		session, err := ss.Start(user)
		if err != nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		w := content.JSONDataWriter{}
		return w.Write(c.Response, NewLoginResponse(*session))
	}
}

// Routes for demo create/read
func (h Handler) Routes(api *routing.RouteGroup, logger *log.Logger, us *service.UserService, ss *service.SessionService) {
	api.Post("/", h.Login(logger, us, ss))
}
