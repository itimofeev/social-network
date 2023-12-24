package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

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
	repo, err := pg.New(context.Background(), pg.Config{
		DSN: os.Getenv("PG_REPOSITORY_DSN"),
	})

	if err != nil {
		return fmt.Errorf("failed to create repository: %w", err)
	}

	f, err := os.Open(os.Getenv("PROFILES_CSV_PATH"))
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}

	var profiles = []entity.Profile{}
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

	return repo.InsertProfiles(context.Background(), profiles)
}
