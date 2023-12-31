package redis

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	RedisDSN string `validate:"required"`
}

type Client struct {
	client *redis.Client
}

func New(cfg Config) (*Client, error) {
	if err := validator.New().Struct(cfg); err != nil {
		return nil, fmt.Errorf("failed to validate config: %w", err)
	}

	opt, err := redis.ParseURL(cfg.RedisDSN)
	if err != nil {
		return nil, fmt.Errorf("failed to parse redis dsn: %w", err)
	}

	client := redis.NewClient(opt)

	return &Client{
		client: client,
	}, nil
}
