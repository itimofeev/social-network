package app

import (
	"context"
	"fmt"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

	"github.com/itimofeev/social-network/internal/entity"
)

type Repository interface {
	GetUserByUserID(ctx context.Context, value string) (entity.User, error)
	InsertUser(ctx context.Context, request entity.CreateUserRequest) (entity.User, error)
	SearchUsers(ctx context.Context, firstName string, lastName string) ([]entity.User, error)
}

type MongoRepository interface {
	SendMessage(ctx context.Context, fromUser, toUser uuid.UUID, messageText string, ts time.Time) error
	ListMessages(ctx context.Context, fromUser, toUser uuid.UUID, laterThan time.Time) ([]entity.Message, error)
}

type Config struct {
	PGRepository    Repository      `validate:"required"`
	MongoRepository MongoRepository `validate:"required"`

	PasetoSecretKey string `validate:"required"`
}

type App struct {
	repo      Repository
	mongoRepo MongoRepository
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
		repo:      cfg.PGRepository,
		mongoRepo: cfg.MongoRepository,
		secretKey: secretKey,
	}, nil
}
