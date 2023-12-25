package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Config struct {
	MongoDSN string
}

type Repository struct {
	client *mongo.Client
	db     *mongo.Database
}

func New(ctx context.Context, cfg Config) (*Repository, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(timeoutCtx, options.Client().ApplyURI(cfg.MongoDSN))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongo: %w", err)
	}

	timeoutCtx, cancel = context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	if err := client.Ping(timeoutCtx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("failed to ping mongo: %w", err)
	}

	return &Repository{
		client: client,
		db:     client.Database("sc"),
	}, nil
}

func (r *Repository) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return r.client.Disconnect(ctx)
}
