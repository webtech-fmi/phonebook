package web

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	domain_services "github.com/webtech-fmi/phonebook/backend/go/domain/service"
	"github.com/webtech-fmi/phonebook/backend/go/domain/service/authentication"
	"github.com/webtech-fmi/phonebook/backend/go/domain/service/profile"
	"github.com/webtech-fmi/phonebook/backend/go/infrastructure/log"

	"github.com/webtech-fmi/phonebook/backend/go/contact-service/pkg/configuration"
	"github.com/webtech-fmi/phonebook/backend/go/contact-service/pkg/domain"
	"github.com/webtech-fmi/phonebook/backend/go/contact-service/pkg/infrastructure/storage"
	"github.com/webtech-fmi/phonebook/backend/go/contact-service/pkg/service"
	webcontact "github.com/webtech-fmi/phonebook/backend/go/contact-service/pkg/web/contacts"

	routing "github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/content"
)

func NewContactRepository(ctx context.Context, cfg *configuration.AppConfiguration, logger *log.Logger) (domain.Repository, error) {
	switch cfg.Repository.Adapter {
	// case "memory":
	// 	return memory.NewRepository(ctx, cfg.Repository.Options, logger)
	case "psql":
		return storage.NewRepository(ctx, cfg.Repository.Options, logger)
	default:
		return nil, fmt.Errorf("unknown storage adapter: [%s]", cfg.Repository.Adapter)
	}
}

func NewContactService(r domain.Repository, logger *log.Logger) (*service.ContactService, error) {
	return &service.ContactService{
		Repository: r,
		Logger:     logger,
	}, nil
}

func NewHTTPServices(ctx context.Context, cfg *configuration.AppConfiguration) (*domain_services.HTTPServices, error) {
	httpServices := domain_services.HTTPServices{}
	if cfg.Services != nil {
		for _, serv := range *cfg.Services {
			switch serv.Service {
			case "authentication":
				s, err := authentication.NewService(serv.URI)
				if err != nil {
					return nil, err
				}
				httpServices.Authentication = s
			case "profile":
				s, err := profile.NewService(serv.URI)
				if err != nil {
					return nil, err
				}
				httpServices.Profile = s
			}
		}
	}
	return &httpServices, nil
}

// NewRouter creates a mux with mounted routes and instantiates respective dependencies.
func NewRouter(ctx context.Context, cfg *configuration.AppConfiguration, logger *log.Logger) *routing.Router {
	contactRepository, err := NewContactRepository(ctx, cfg, logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("Could not instantiate the contact repository")
	}

	contactService, err := NewContactService(contactRepository, logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("Could not instantiate the contact service")
	}

	httpServices, err := NewHTTPServices(ctx, cfg)
	if err != nil {
		logger.Fatal().Err(err).Msg("Could not instantiate the profile service")
	}

	r := routing.New()

	contactsAPI := r.Group("/contacts")
	contactsAPI.Use(content.TypeNegotiator(content.JSON))
	webcontact.Handler{}.Routes(contactsAPI, logger, contactService, httpServices)

	return r
}

// LaunchServer starts a web server and propagates shutdown context.
func LaunchServer(cfg *configuration.AppConfiguration, logger *log.Logger) error {
	var err error

	c := make(chan os.Signal, 1)
	signal.Notify(
		c,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		s := <-c
		logger.Debug().Str("syscall", s.String()).Msg("Intercepted syscall")
		cancel()
	}()

	router := NewRouter(ctx, cfg, logger)
	srv := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%d", cfg.Port),
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal().Err(err).Msg("Could not launch the web server")
		}
	}()
	logger.Printf("Starting server on port: [%d]", cfg.Port)

	<-ctx.Done()

	logger.Printf("Cleaning up the server")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	if err = srv.Shutdown(ctxShutDown); err != nil {
		logger.Fatal().Err(err).Msg("Error on server shutdown")
	}

	cancel()

	logger.Printf("Server exited successfully")

	if err == http.ErrServerClosed {
		err = nil
	}
	return err
}
