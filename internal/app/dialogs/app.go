package dialogs

import (
	"context"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

	"github.com/itimofeev/social-network/internal/entity"
)

type repository interface {
	SendMessage(ctx context.Context, fromUser, toUser uuid.UUID, messageText string, ts time.Time) error
	ListMessages(ctx context.Context, fromUser, toUser uuid.UUID, laterThan time.Time) ([]entity.Message, error)
}

type backendClient interface {
	IsUserExists(ctx context.Context, userID uuid.UUID) (bool, error)
}

type Config struct {
	MongoRepo     repository    `validate:"required"`
	BackendClient backendClient `validate:"required"`
}

type App struct {
	repo          repository
	backendClient backendClient
}

func NewApp(cfg Config) (*App, error) {
	if err := validator.New().Struct(cfg); err != nil {
		return nil, fmt.Errorf("failed to validate config: %w", err)
	}

	return &App{
		repo:          cfg.MongoRepo,
		backendClient: cfg.BackendClient,
	}, nil
}

func (a *App) SendMessage(ctx context.Context, fromUser, toUser uuid.UUID, messageText string, ts time.Time) error {
	toUserExists, err := a.backendClient.IsUserExists(ctx, toUser)
	if err != nil {
		return fmt.Errorf("failed to check if user exists: %w", err)
	}
	if !toUserExists {
		return fmt.Errorf("user with id %s does not exist", toUser)
	}

	return a.repo.SendMessage(ctx, fromUser, toUser, messageText, ts)
}

func (a *App) ListMessages(ctx context.Context, fromUser, toUser uuid.UUID, laterThan time.Time) ([]entity.Message, error) {
	toUserExists, err := a.backendClient.IsUserExists(ctx, toUser)
	if err != nil {
		return nil, fmt.Errorf("failed to check if user exists: %w", err)
	}
	if !toUserExists {
		return nil, fmt.Errorf("user with id %s does not exist", toUser)
	}

	return a.repo.ListMessages(ctx, fromUser, toUser, laterThan)
}
