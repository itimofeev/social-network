package pg

import (
	"context"
	"time"

	"github.com/google/uuid"
)

func (r *Repository) CreatePost(ctx context.Context, userID uuid.UUID, text string) (string, error) {
	query := `
INSERT INTO
    posts (id, user_id, text, created_at)
VALUES ($1, $2, $3, $4)
RETURNING id`
	args := []interface{}{uuid.New(), userID, text, time.Now()}

	var postID string

	return postID, r.getTx(ctx).QueryRow(ctx, query, args...).Scan(&postID)
}
