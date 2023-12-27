package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"

	"github.com/go-faster/errors"

	"github.com/itimofeev/social-network/internal/entity"
	"github.com/itimofeev/social-network/internal/repository/pg"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx := signalContext(context.Background())

	repo, err := pg.New(ctx, pg.Config{
		DSN:          os.Getenv("PG_REPOSITORY_DSN"),
		MaxOpenConns: 10,
	})

	if err != nil {
		return fmt.Errorf("failed to create repository: %w", err)
	}

	f, err := os.Open(os.Getenv("PROFILES_CSV_PATH"))
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}

	var profiles []entity.Profile
	csvReader := csv.NewReader(f)

	for {
		record, err := csvReader.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read record: %w", err)
		}

		split := strings.Split(record[0], " ")
		if len(split) != 2 {
			return fmt.Errorf("failed to split record: %s", record[0])
		}
		age, err := strconv.ParseInt(record[1], 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse age: %w", err)
		}

		profiles = append(profiles, entity.Profile{
			FirstName: split[1],
			LastName:  split[0],
			Age:       age,
			City:      record[2],
		})
	}

	fmt.Println("loaded profiles:", len(profiles))

	totalInserted := atomic.Int64{}
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(workerNumber int) {
			defer wg.Done()
			for j := 0; j < len(profiles); j++ {
				if err := repo.InsertProfiles(ctx, []entity.Profile{profiles[j]}); err != nil {
					slog.Info("error on inserting profile", "worker", workerNumber, "profile", j, "error", err)
					return
				}

				if newInserted := totalInserted.Add(1); newInserted%1000 == 0 {
					slog.Info("inserted profiles", "inserted", newInserted)
				}
			}
		}(i)
	}

	wg.Wait()

	slog.Info("inserting of profiles finished", "inserted", totalInserted.Load())

	return nil
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
