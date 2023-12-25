package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kelseyhightower/envconfig"
	"golang.org/x/sync/errgroup"

	"github.com/itimofeev/social-network/internal/app/dialogs"
	"github.com/itimofeev/social-network/internal/repository/mongo"
	dialogs2 "github.com/itimofeev/social-network/internal/server/dialogs"
	"github.com/itimofeev/social-network/pkg/xlog"
)

type configuration struct {
	Port            string        `envconfig:"PORT" default:"8080"`
	ReadTimeout     time.Duration `envconfig:"READ_TIMEOUT" default:"10s"`
	WriteTimeout    time.Duration `envconfig:"WRITE_TIMEOUT" default:"10s"`
	ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" default:"10s"`

	MongoRepositoryDSN string `envconfig:"MONGO_REPOSITORY_DSN" required:"true"`
}

func main() {
	cfg := mustParseConfig()

	ctx := signalContext(context.Background())

	slog.InfoContext(ctx, "service is starting")

	if err := run(ctx, cfg); err != nil {
		log.Fatalf("service is stopped with error: %s", err)
	}

	slog.InfoContext(ctx, "service is stopped")
}

func run(ctx context.Context, cfg configuration) error {
	xlog.InitSlog()

	mongoRepo, err := mongo.New(ctx, mongo.Config{
		MongoDSN: cfg.MongoRepositoryDSN,
	})
	if err != nil {
		return err
	}

	app, err := dialogs.NewApp(dialogs.Config{
		MongoRepo: mongoRepo,
	})
	if err != nil {
		return err
	}

	srv, err := dialogs2.NewServer(dialogs2.Config{
		Domain:          "http://localhost:8080",
		Version:         "1.0.0",
		Port:            cfg.Port,
		ReadTimeout:     cfg.ReadTimeout,
		WriteTimeout:    cfg.WriteTimeout,
		ShutdownTimeout: cfg.ShutdownTimeout,

		App: app,
	})
	if err != nil {
		return err
	}

	errGr, errGrCtx := errgroup.WithContext(ctx)

	errGr.Go(func() error {
		slog.InfoContext(ctx, "start http server")

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
		slog.InfoContext(ctx, "received signal", sig)
		cancel()
	}()

	return cnCtx
}
