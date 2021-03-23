package demo

import (
	"net/http"

	"github.com/webtech-fmi/phonebook/backend/contact-service/pkg/service"

	routing "github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/content"
	"github.com/rs/zerolog"
)

const (
	// You decide if you want to wrap errors or
	// will use values.
	ErrGetDemoParam    = "get_demo_param"
	ErrGetDemoLoad     = "get_demo_load"
	ErrCreateDemoParam = "create_demo_param"
	ErrCreateDemoStore = "create_demo_store"
)

// Handler is just a route collection
type Handler struct {}

// GetDemo Load a specific demo by ID - only "demo" will be found
func (h Handler) GetDemo(logger *zerolog.Logger,ds *service.DemoService) func (c *routing.Context) error {
	return func(c *routing.Context) error {
		ID := c.Query("id")
		if ID == "" {
			return routing.NewHTTPError(http.StatusBadRequest, "passed an empty ID")
		}

		demo, err := ds.GetByID(ID)
		if err != nil {
			return routing.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		w := content.JSONDataWriter{}

		return w.Write(c.Response, NewFetchResponse(*demo, ds))
	}
}

// Routes for demo create/read
func (h Handler) Routes(api *routing.RouteGroup, logger *zerolog.Logger, ds *service.DemoService) {
	api.Get("/demo", h.GetDemo(logger, ds))
}
