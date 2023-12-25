package backend

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

	"github.com/itimofeev/social-network/internal/server/backend/gen/api"
)

type Config struct {
	BackendAPI apiClient
}

type apiClient interface {
	UserGetIDGet(ctx context.Context, params api.UserGetIDGetParams) (api.UserGetIDGetRes, error)
}

type Client struct {
	client apiClient
}

func NewClient(cfg Config) (*Client, error) {
	if err := validator.New().Struct(cfg); err != nil {
		return nil, fmt.Errorf("failed to validate config: %w", err)
	}

	return &Client{
		client: cfg.BackendAPI,
	}, nil
}

func (c *Client) IsUserExists(ctx context.Context, userID uuid.UUID) (bool, error) {
	slog.DebugContext(ctx, "trying to check if user exists", "userID", userID)

	res, err := c.client.UserGetIDGet(ctx, api.UserGetIDGetParams{
		ID: api.UserId(userID.String()),
	})
	if err != nil {
		return false, err
	}

	if _, ok := res.(*api.UserGetIDGetNotFound); ok {
		return false, nil
	}

	if _, ok := res.(*api.User); !ok {
		return false, fmt.Errorf("failed to convert response")
	}

	return true, nil
}
