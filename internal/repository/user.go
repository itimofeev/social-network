package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/itimofeev/social-network/internal/entity"
)

func userColumns() []string {
	return []string{
		"id",
		"user_id",
		"password",
		"first_name",
		"second_name",
		"birthdate",
		"biography",
		"interests",
		"city",
	}
}

func (r *Repository) InsertUser(ctx context.Context, req entity.CreateUserRequest) (entity.User, error) {
	insertColumns := []string{
		"id",
		"user_id",
		"password",
		"first_name",
		"second_name",
		"birthdate",
		"biography",
		"interests",
		"city",
	}
	builder := sq.Insert("users").
		Columns(insertColumns...).
		Suffix("RETURNING *").
		PlaceholderFormat(sq.Dollar)

	builder = builder.Values(
		uuid.New(),
		req.UserID,
		req.Password,
		req.FirstName,
		req.SecondName,
		req.BirthDate,
		req.Biography,
		req.Interests,
		req.City)

	query, args, err := builder.ToSql()
	if err != nil {
		return entity.User{}, fmt.Errorf("failed to build query: %w", err)
	}

	row := r.getTx(ctx).QueryRowContext(ctx, query, args...)

	user, err := scanUser(row)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *Repository) GetUserByUserID(ctx context.Context, userID string) (entity.User, error) {
	builder := sq.Select(userColumns()...).
		From("users").
		Where("user_id = ?", userID).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return entity.User{}, fmt.Errorf("failed to  building sql: %w", err)
	}

	row := r.getTx(ctx).QueryRowContext(ctx, query, args...)
	if err != nil {
		return entity.User{}, fmt.Errorf("failed to query: %w", err)
	}

	user, err := scanUser(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, entity.ErrUserNotFound
		}
		return entity.User{}, err
	}

	return user, nil
}

func (r *Repository) SearchUsers(ctx context.Context, firstName string, lastName string) ([]entity.User, error) {
	builder := sq.Select(userColumns()...).
		From("users").
		//Where("first_name LIKE '?' AND second_name LIKE '?'", firstName, lastName).
		Where("first_name LIKE ?", firstName+"%").
		Where("second_name LIKE ?", lastName+"%").
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to  building sql: %w", err)
	}

	rows, err := r.getTx(ctx).QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query: %w", err)
	}
	defer rows.Close()

	resp := make([]entity.User, 0)
	for rows.Next() {
		user, err := scanUser(rows)
		if err != nil {
			return nil, fmt.Errorf("error on scan user: %w", err)
		}

		resp = append(resp, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred: %w", err)
	}

	return resp, nil
}

func scanUser(rows rowScanner) (entity.User, error) {
	var user entity.User

	if err := rows.Scan(
		&user.ID,
		&user.UserID,
		&user.Password,
		&user.FirstName,
		&user.SecondName,
		&user.BirthDate,
		&user.Biography,
		&user.Interests,
		&user.City,
	); err != nil {
		return user, fmt.Errorf("failed to scan row: %w", err)
	}

	return user, nil
}
