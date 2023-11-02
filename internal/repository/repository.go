package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/go-playground/validator/v10"
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
	db  *sql.DB
	cfg Config
}

func New(ctx context.Context, cfg Config) (*Repository, error) {
	slog.Debug("try to validate repository config")

	if err := validator.New().Struct(cfg); err != nil {
		return nil, fmt.Errorf("failed to validate repository config: %w", err)
	}

	slog.Debug("try to open sql connection")

	conn, err := sql.Open("postgres", cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("failed to open sql connection: %w", err)
	}

	slog.Info("trying to ping db")

	if err := conn.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping sql connection: %w", err)
	}

	slog.Info("repository initialized")

	return &Repository{
		db:  conn,
		cfg: cfg,
	}, nil
}

func (r *Repository) Close() error {
	return r.db.Close()
}

type tx interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
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
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return fmt.Errorf("failed to db.BeginTx: %w", err)
	}

	defer func() {
		if errRollback := tx.Rollback(); errRollback != nil && !errors.Is(errRollback, sql.ErrTxDone) {
			slog.Error("failed to tx.Rollback", errRollback)
		}
	}()

	// run callback
	if err = tFunc(injectTx(ctx, tx)); err != nil {
		return err
	}
	// if no error, commit
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to tx.Commit: %w", err)
	}
	return nil
}

type txKey struct{}

// injectTx injects transaction to context
//
//nolint:unused // will be used in future
func injectTx(ctx context.Context, tx *sql.Tx) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

// extractTx extracts transaction from context
func extractTx(ctx context.Context) *sql.Tx {
	if tx, ok := ctx.Value(txKey{}).(*sql.Tx); ok {
		return tx
	}
	return nil
}

type rowScanner interface {
	Scan(dest ...any) error
	Err() error
}
