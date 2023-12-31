package pg

import (
	"context"

	"github.com/google/uuid"
)

func (r *Repository) SetFollower(ctx context.Context, userID uuid.UUID, followUserID uuid.UUID) error {
	query := `
INSERT INTO
    followers (user_id, follows_user_id)
VALUES ($1, $2)
ON CONFLICT (user_id, follows_user_id) DO NOTHING`
	args := []interface{}{userID, followUserID}
	_, err := r.getTx(ctx).Exec(ctx, query, args...)

	return err
}

func (r *Repository) DeleteFollower(ctx context.Context, userID uuid.UUID, followUserID uuid.UUID) error {
	query := `
DELETE FROM
    followers
WHERE user_id = $1 AND follows_user_id = $2`
	args := []interface{}{userID, followUserID}
	_, err := r.getTx(ctx).Exec(ctx, query, args...)

	return err
}
