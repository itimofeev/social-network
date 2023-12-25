package dialogs

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/pprof"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/go-playground/validator/v10"

	oapi "github.com/itimofeev/social-network/api"
	"github.com/itimofeev/social-network/internal/app/dialogs"
	"github.com/itimofeev/social-network/internal/server/dialogs/gen/api"
	"github.com/itimofeev/social-network/pkg/xmw"
)

type Config struct {
	Domain  string `validate:"required"`
	Version string `validate:"required"`

	Port            string        `validate:"required"`
	ReadTimeout     time.Duration `validate:"required"`
	WriteTimeout    time.Duration `validate:"required"`
	ShutdownTimeout time.Duration `validate:"required"`

	App *dialogs.App `validate:"required"`
}

type Server struct {
	cfg       *Config
	srv       *http.Server
	app       *dialogs.App
	apiServer *api.Server
}

func NewServer(cfg Config) (*Server, error) {
	err := validator.New().Struct(cfg)
	if err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}

	handler := NewHandler(cfg.App)

	opts := []api.ServerOption{
		api.WithPathPrefix("/api/v1"),
		api.WithErrorHandler(handler.ErrorHandler),
	}
	apiServer, err := api.NewServer(handler, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create api server: %w", err)
	}

	s := &Server{
		cfg:       &cfg,
		srv:       srv,
		app:       cfg.App,
		apiServer: apiServer,
	}

	s.init()
	return s, nil
}

func (s *Server) Serve(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		ctxShutdown, cancel := context.WithTimeout(context.Background(), s.cfg.ShutdownTimeout)
		defer cancel()

		//nolint:contextcheck // intentionally separate context for shutting down as ctx is already closed
		if err := s.srv.Shutdown(ctxShutdown); err != nil {
			slog.WarnContext(ctx, "failed to shutdown server: %w", err)
		}
	}()

	if err := s.srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("failed to listen and serve: %w", err)
	}

	slog.DebugContext(ctx, "server gracefully stopped")

	return nil
}

func (s *Server) init() {
	r := chi.NewRouter()
	r.Use(
		middleware.Recoverer,
	)
	r.NotFound(s.notFoundHandler)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://backoffice.test.env", "https://backoffice.exness.io"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Group(func(router chi.Router) {
		router.Get("/metrics", promhttp.Handler().ServeHTTP)
		router.Get("/ping", s.pingHandler)
		router.Get("/version", s.versionHandler)

		router.Get("/swagger.yaml", oapi.BackendOapiSchemaHandler)
		router.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, r.RequestURI+"/index.html", http.StatusMovedPermanently)
		})
		router.Get("/docs/*", httpSwagger.Handler(
			httpSwagger.URL(fmt.Sprintf("https://%s/swagger.yaml", s.cfg.Domain)),
		))

		router.Group(func(telemetry chi.Router) {
			telemetry.HandleFunc("/pprof", func(w http.ResponseWriter, r *http.Request) {
				http.Redirect(w, r, r.RequestURI+"/", http.StatusMovedPermanently)
			})
			telemetry.HandleFunc("/pprof/*", pprof.Index)
			telemetry.HandleFunc("/pprof/cmdline", pprof.Cmdline)
			telemetry.HandleFunc("/pprof/profile", pprof.Profile)
			telemetry.HandleFunc("/pprof/symbol", pprof.Symbol)
			telemetry.HandleFunc("/pprof/trace", pprof.Trace)

			telemetry.Handle("/pprof/goroutine", pprof.Handler("goroutine"))
			telemetry.Handle("/pprof/threadcreate", pprof.Handler("threadcreate"))
			telemetry.Handle("/pprof/mutex", pprof.Handler("mutex"))
			telemetry.Handle("/pprof/heap", pprof.Handler("heap"))
			telemetry.Handle("/pprof/block", pprof.Handler("block"))
			telemetry.Handle("/pprof/allocs", pprof.Handler("allocs"))
		})
	})

	r.Group(func(r chi.Router) {
		r.Use(xmw.RequestID)
		r.Use(middleware.DefaultLogger)

		r.Route("/api/v1", func(api chi.Router) {
			api.Handle("/*", s.apiServer)
		})
	})

	s.srv.Handler = r
}

func (s *Server) notFoundHandler(w http.ResponseWriter, _ *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func (s *Server) pingHandler(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, "pong")
}

func (s *Server) versionHandler(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, s.cfg.Version)
}
