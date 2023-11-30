package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/itimofeev/social-network/internal/app"
)

type Config struct {
	DSN          string `validate:"required,url"`
	MaxOpenConns int    `validate:"gte=0"`
}

var (
	_ app.Repository = (*Repository)(nil)
)

type Repository struct {
	db  *pgxpool.Pool
	cfg Config
}

func New(ctx context.Context, cfg Config) (*Repository, error) {
	slog.Debug("try to validate repository config")

	if err := validator.New().Struct(cfg); err != nil {
		return nil, fmt.Errorf("failed to validate repository config: %w", err)
	}

	slog.Debug("try to open sql connection")

	config, err := pgxpool.ParseConfig(cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("failed to parse DSN: %w", err)
	}
	config.MaxConns = int32(cfg.MaxOpenConns)

	conn, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed to open sql connection: %w", err)
	}

	slog.Info("trying to ping db")

	if err := conn.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping sql connection: %w", err)
	}

	slog.Info("repository initialized")

	return &Repository{
		db:  conn,
		cfg: cfg,
	}, nil
}

func (r *Repository) Close() {
	r.Close()
}

type tx interface {
	Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
}

func (r *Repository) getTx(ctx context.Context) tx {
	if tx := extractTx(ctx); tx != nil {
		return tx
	}
	return r.db
}

//nolint:unused // will be used in future
func (r *Repository) withinTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error {
	// begin transaction
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return fmt.Errorf("failed to db.BeginTx: %w", err)
	}

	defer func() {
		if errRollback := tx.Rollback(context.Background()); errRollback != nil && !errors.Is(errRollback, sql.ErrTxDone) {
			slog.Error("failed to tx.Rollback", errRollback)
		}
	}()

	// run callback
	if err = tFunc(injectTx(ctx, tx)); err != nil {
		return err
	}
	// if no error, commit
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to tx.Commit: %w", err)
	}
	return nil
}

type txKey struct{}

// injectTx injects transaction to context
//
//nolint:unused // will be used in future
func injectTx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

// extractTx extracts transaction from context
func extractTx(ctx context.Context) pgx.Tx {
	if tx, ok := ctx.Value(txKey{}).(pgx.Tx); ok {
		return tx
	}
	return nil
}

type rowScanner interface {
	Scan(dest ...any) error
}
