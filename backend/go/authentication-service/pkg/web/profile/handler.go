package profile

import (
	//	"net/http"

	"fmt"

	"github.com/webtech-fmi/phonebook/backend/go/authentication-service/pkg/service"

	routing "github.com/go-ozzo/ozzo-routing"
	//	"github.com/go-ozzo/ozzo-routing/content"
	"github.com/webtech-fmi/phonebook/backend/go/infrastructure/log"
)

const (
	// You decide if you want to wrap errors or
	// will use values.
	ErrCreateUserParam = "create_user_param"
)

// Handler is just a route collection
type Handler struct{}

// // GetDemo Load a specific demo by ID - only "demo" will be found
// func (h Handler) GetDemo(logger *zerolog.Logger, ds *service.DemoService) func(c *routing.Context) error {
// 	return func(c *routing.Context) error {
// 		ID := c.Query("id")
// 		if ID == "" {
// 			return routing.NewHTTPError(http.StatusBadRequest, "passed an empty ID")
// 		}

// 		demo, err := ds.GetByID(ID)
// 		if err != nil {
// 			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
// 		}
// 		w := content.JSONDataWriter{}

// 		return w.Write(c.Response, NewFetchResponse(*demo, ds))
// 	}
// }

func (h Handler) CreateUser(logger *log.Logger, ds *service.UserService) func(c *routing.Context) error {
	return func(c *routing.Context) error {
		request := CreateRequest{}

		if err := c.Read(&request); err != nil {
			return err
		}

		if err := ds.CreateUser(request.Name, request.Date); err != nil {
			return err
		}

		fmt.Println(request)
		return nil
	}
}

// Routes for demo create/read
func (h Handler) Routes(api *routing.RouteGroup, logger *log.Logger, s *service.UserService) {
	//	api.Get("/demo", h.GetDemo(logger, ds))
	api.Post("/create", h.CreateUser(logger, s))
}
