package backend

import (
	"context"

	"github.com/google/uuid"
)

func (a *App) CreatePost(ctx context.Context, id uuid.UUID, text string) (string, error) {
	return a.repo.CreatePost(ctx, id, text)
}
