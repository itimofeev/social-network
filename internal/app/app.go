package app

import (
	"context"
	"fmt"

	"aidanwoods.dev/go-paseto"
	"github.com/go-playground/validator/v10"

	"github.com/itimofeev/social-network/internal/entity"
)

type Repository interface {
	GetUserByUserID(ctx context.Context, value string) (entity.User, error)
	InsertUser(ctx context.Context, request entity.CreateUserRequest) (entity.User, error)
}

type Config struct {
	Repository Repository `validate:"required"`

	PasetoSecretKey string `validate:"required"`
}

type App struct {
	repo      Repository
	secretKey paseto.V4SymmetricKey
}

func New(cfg Config) (*App, error) {
	if err := validator.New().Struct(cfg); err != nil {
		return nil, fmt.Errorf("failed to validate repository config: %w", err)
	}

	secretKey, err := paseto.V4SymmetricKeyFromHex(cfg.PasetoSecretKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse paseto secret key: %w", err)
	}

	return &App{
		repo:      cfg.Repository,
		secretKey: secretKey,
	}, nil
}
