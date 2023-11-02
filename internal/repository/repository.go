package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/itimofeev/social-network/internal/app"
)

type Config struct {
	DSN string `validate:"required,url"`
}

var (
	_ app.Repository = (*Repository)(nil)
)

type Repository struct {
	db  *sqlx.DB
	cfg Config
}

func New(ctx context.Context, cfg Config) (*Repository, error) {
	slog.Debug("try to validate repository config")

	if err := validator.New().Struct(cfg); err != nil {
		return nil, fmt.Errorf("failed to validate repository config: %w", err)
	}

	slog.Debug("try to open sql connection")

	conn, err := sqlx.Open("postgres", cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("failed to open sql connection: %w", err)
	}

	slog.Info("trying to ping db")

	if err := conn.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping sql connection: %w", err)
	}

	slog.Info("repository initialised")

	return &Repository{
		db:  conn,
		cfg: cfg,
	}, nil
}

func (r *Repository) Close() error {
	return r.db.Close()
}

//nolint:unused // maybe useful in future
func (r *Repository) doInTx(ctx context.Context, queries func(context.Context, *sql.Tx) error) error {
	slog.Debug("try to begin sql transaction")

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin sql transaction: %w", err)
	}
	defer func() {
		slog.Debug("try to rollback sql transaction")

		if _err := tx.Rollback(); _err != nil && !errors.Is(_err, sql.ErrTxDone) {
			slog.Error("failed to rollback transaction", _err)
		}
	}()

	slog.Debug("try to execute queries")

	if err = queries(ctx, tx); err != nil {
		return err
	}

	slog.Debug("try to commit sql transaction")

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit sql transaction: %w", err)
	}

	return nil
}
