package backend

import (
	"context"

	"github.com/itimofeev/social-network/internal/entity"
)

func (a *App) SearchUsers(ctx context.Context, firstName string, lastName string) ([]entity.User, error) {
	return a.repoReplica.SearchUsers(ctx, firstName, lastName)
}
