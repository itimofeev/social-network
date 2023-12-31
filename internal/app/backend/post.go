package backend

import (
	"context"

	"github.com/google/uuid"

	"github.com/itimofeev/social-network/internal/entity"
)

func (a *App) CreatePost(ctx context.Context, id uuid.UUID, text string) (string, error) {
	return a.repo.CreatePost(ctx, id, text)
}

func (a *App) DeletePost(ctx context.Context, postID string) error {
	return a.repo.DeletePost(ctx, postID)
}

func (a *App) UpdatePost(ctx context.Context, id string, s string) error {
	return a.repo.UpdatePost(ctx, id, s)
}

func (a *App) GetPost(ctx context.Context, id string) (entity.Post, error) {
	return a.repoReplica.GetPost(ctx, id)
}
