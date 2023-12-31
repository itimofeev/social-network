package backend

import (
	"context"

	"github.com/go-faster/errors"
	"github.com/google/uuid"
)

func (a *App) DeleteFollower(ctx context.Context, userID uuid.UUID, followUserID uuid.UUID) error {
	return a.repo.DeleteFollower(ctx, userID, followUserID)
}

func (a *App) SetFollower(ctx context.Context, userID uuid.UUID, followUserID uuid.UUID) error {
	if userID == followUserID {
		return errors.New("user can't follow himself")
	}
	return a.repo.SetFollower(ctx, userID, followUserID)
}
