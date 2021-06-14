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
	"github.com/webtech-fmi/phonebook/backend/go/authentication-service/pkg/web/login"
	"github.com/webtech-fmi/phonebook/backend/go/authentication-service/pkg/web/logout"
	"github.com/webtech-fmi/phonebook/backend/go/authentication-service/pkg/web/sessions"
	"github.com/webtech-fmi/phonebook/backend/go/authentication-service/pkg/web/users"

	routing "github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/content"
)

func NewUserRepository(ctx context.Context, cfg *configuration.AppConfiguration, logger *log.Logger) (domain.Repository, error) {
	switch cfg.User.Repository.Adapter {
	// case "memory":
	// 	return memory.NewRepository(ctx, cfg.Repository.Options, logger)
	case "psql":
		return storage.NewRepository(ctx, cfg.User.Repository.Options, logger)
	default:
		return nil, fmt.Errorf("unknown storage adapter: [%s]", cfg.User.Repository.Adapter)
	}
}

func NewUserService(r domain.Repository, logger *log.Logger) (*service.UserService, error) {
	return &service.UserService{
		Repository: r,
		Logger:     logger,
	}, nil
}

func NewSessionRepository(ctx context.Context, cfg *configuration.AppConfiguration) (domain.SessionRepository, error) {
	switch cfg.Session.Repository.Adapter {
	// case "memory":
	// 	return memory.NewRepository(ctx, cfg.Repository.Options, logger)
	case "redis":
		return storage.NewSessionRepository(cfg.Session.Repository.Options)
	default:
		return nil, fmt.Errorf("unknown storage adapter: [%s]", cfg.Session.Repository.Adapter)
	}
}

func NewSessionService(r domain.SessionRepository, cfg *configuration.AppConfiguration) (*service.SessionService, error) {
	return service.NewSessionService(r, cfg.Session.Options)
}

// NewRouter creates a mux with mounted routes and instantiates respective dependencies.
func NewRouter(ctx context.Context, cfg *configuration.AppConfiguration, logger *log.Logger) *routing.Router {
	userRepository, err := NewUserRepository(ctx, cfg, logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("Could not instantiate the users repository")
	}

	userService, err := NewUserService(userRepository, logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("Could not instantiate the user service")
	}

	sessionRepository, err := NewSessionRepository(ctx, cfg)
	if err != nil {
		logger.Fatal().Err(err).Msg("Could not instantiate the session repository")
	}

	sessionService, err := NewSessionService(sessionRepository, cfg)
	if err != nil {
		logger.Fatal().Err(err).Msg("Could not instantiate the session service")
	}

	r := routing.New()

	usersAPI := r.Group("/users")
	usersAPI.Use(content.TypeNegotiator(content.JSON))
	users.Handler{}.Routes(usersAPI, logger, userService)

	sessionAPI := r.Group("/sessions")
	sessionAPI.Use(content.TypeNegotiator(content.JSON))
	sessions.Handler{}.Routes(sessionAPI, logger, sessionService)

	loginAPI := r.Group("/login")
	loginAPI.Use(content.TypeNegotiator(content.JSON))
	login.Handler{}.Routes(loginAPI, logger, userService, sessionService)

	logoutAPI := r.Group("/logout")
	logoutAPI.Use(content.TypeNegotiator(content.JSON))
	logout.Handler{}.Routes(logoutAPI, logger, sessionService)

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
