package backend

import (
	"context"

	"github.com/itimofeev/social-network/internal/entity"
)

func (a *App) GetUser(ctx context.Context, userID string) (entity.User, error) {
	user, err := a.repoReplica.GetUserByUserID(ctx, userID)
	return user, err
}
