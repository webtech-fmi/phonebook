package web

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/webtech-fmi/phonebook/backend/go/infrastructure/log"

	"github.com/webtech-fmi/phonebook/backend/go/authentication-service/pkg/configuration"
	"github.com/webtech-fmi/phonebook/backend/go/authentication-service/pkg/domain"
	"github.com/webtech-fmi/phonebook/backend/go/authentication-service/pkg/infrastructure/storage"
	"github.com/webtech-fmi/phonebook/backend/go/authentication-service/pkg/service"
	webprofile "github.com/webtech-fmi/phonebook/backend/go/authentication-service/pkg/web/profile"

	routing "github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/content"
)

func NewProfileRepository(ctx context.Context, cfg *configuration.AppConfiguration, logger *log.Logger) (domain.Repository, error) {
	switch cfg.Repository.Adapter {
	// case "memory":
	// 	return memory.NewRepository(ctx, cfg.Repository.Options, logger)
	case "psql":
		return storage.NewRepository(ctx, cfg.Repository.Options, logger)
	default:
		return nil, fmt.Errorf("unknown storage adapter: [%s]", cfg.Repository.Adapter)
	}
}


func NewProfileService(r domain.Repository, logger *log.Logger) (*service.UserService, error) {
	return &service.UserService{
		Repository: r,
		Logger:     logger,
	}, nil
}

// NewRouter creates a mux with mounted routes and instantiates respective dependencies.
func NewRouter(ctx context.Context, cfg *configuration.AppConfiguration, logger *log.Logger) *routing.Router {
	profileRepository, err := NewProfileRepository(ctx, cfg, logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("Could not instantiate the users repository")
	}

	profileService, err := NewProfileService(profileRepository, logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("Could not instantiate the demo service")
	}

	r := routing.New()

	profilesAPI := r.Group("/profiles")
	profilesAPI.Use(content.TypeNegotiator(content.JSON))
	webprofile.Handler{}.Routes(profilesAPI, logger, profileService)

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
