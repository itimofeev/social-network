package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kelseyhightower/envconfig"
	"golang.org/x/sync/errgroup"

	"github.com/itimofeev/social-network/internal/app"
	"github.com/itimofeev/social-network/internal/repository"
	"github.com/itimofeev/social-network/internal/server"
)

type configuration struct {
	Port            string        `envconfig:"PORT" default:"8080"`
	ReadTimeout     time.Duration `envconfig:"READ_TIMEOUT" default:"10s"`
	WriteTimeout    time.Duration `envconfig:"WRITE_TIMEOUT" default:"10s"`
	ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" default:"10s"`

	SessionSecretKey string `envconfig:"SESSION_SECRET_KEY" default:"5468ac74e23ea5c297413a3020af91601f22c82e77aa89cca4e8fb4ec28fb300"` // this have to be passed from secured store like vault in production env

	RepositoryDSN string `envconfig:"PG_REPOSITORY_DSN" required:"true"`
}

func main() {
	cfg := mustParseConfig()

	ctx := signalContext(context.Background())

	slog.Info("service is starting")

	if err := run(ctx, cfg); err != nil {
		log.Fatalf("service is stopped with error: %s", err)
	}

	slog.Info("service is stopped")
}

func run(ctx context.Context, cfg configuration) error {
	repo, err := repository.New(ctx, repository.Config{
		DSN:          cfg.RepositoryDSN,
		MaxOpenConns: 10,
	})
	if err != nil {
		return err
	}

	application, err := app.New(app.Config{
		Repository:      repo,
		PasetoSecretKey: cfg.SessionSecretKey,
	})
	if err != nil {
		return err
	}

	srv, err := server.NewServer(server.Config{
		BaseContextFn: func(_ net.Listener) context.Context {
			return ctx
		},
		Domain:          "http://localhost:8080",
		Version:         "1.0.0",
		Port:            cfg.Port,
		ReadTimeout:     cfg.ReadTimeout,
		WriteTimeout:    cfg.WriteTimeout,
		ShutdownTimeout: cfg.ShutdownTimeout,

		App: application,
	})
	if err != nil {
		return err
	}

	errGr, errGrCtx := errgroup.WithContext(ctx)

	errGr.Go(func() error {
		slog.Info("start http server")

		return srv.Serve(errGrCtx)
	})

	return errGr.Wait()
}

func mustParseConfig() configuration {
	var cfg configuration

	if err := envconfig.Process("", &cfg); err != nil {
		panic(fmt.Sprintf("failed to load configuration: %s", err))
	}

	return cfg
}

// signalContext returns a context that is canceled if either SIGTERM or SIGINT signal is received.
func signalContext(ctx context.Context) context.Context {
	cnCtx, cancel := context.WithCancel(ctx)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		sig := <-c
		slog.Info("received signal", sig)
		cancel()
	}()

	return cnCtx
}
