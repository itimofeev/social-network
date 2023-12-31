package pg

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/itimofeev/social-network/internal/entity"
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

func (r *Repository) DeletePost(ctx context.Context, postID string) error {
	query := `
DELETE FROM
    posts 
WHERE id = $1`
	args := []interface{}{postID}

	_, err := r.getTx(ctx).Exec(ctx, query, args...)
	return err
}

func (r *Repository) UpdatePost(ctx context.Context, postID string, text string) error {
	query := `
UPDATE
    posts
SET
	text = $1
WHERE id = $2`
	args := []interface{}{postID, text}

	_, err := r.getTx(ctx).Exec(ctx, query, args...)
	return err
}

func (r *Repository) GetPost(ctx context.Context, postID string) (entity.Post, error) {
	query := `
SELECT
	id, user_id, text, created_at
FROM
    posts
WHERE id = $1`
	args := []interface{}{postID}

	post := entity.Post{}
	if err := r.getTx(ctx).QueryRow(ctx, query, args...).Scan(
		&post.ID,
		&post.AuthorUserID,
		&post.Text,
		&post.CreatedAt,
	); err != nil {
		return entity.Post{}, err
	}

	return post, nil
}
