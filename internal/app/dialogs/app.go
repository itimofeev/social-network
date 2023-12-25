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

type Config struct {
	MongoRepo repository `validate:"required"`
}

type App struct {
	repo repository
}

func NewApp(cfg Config) (*App, error) {
	if err := validator.New().Struct(cfg); err != nil {
		return nil, fmt.Errorf("failed to validate config: %w", err)
	}

	return &App{
		repo: cfg.MongoRepo,
	}, nil
}
