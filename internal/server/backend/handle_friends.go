package backend

import (
	"context"

	"github.com/google/uuid"

	"github.com/itimofeev/social-network/internal/server/backend/gen/api"
)

func (h *Handler) FriendDeleteUserIDPut(ctx context.Context, params api.FriendDeleteUserIDPutParams) (api.FriendDeleteUserIDPutRes, error) {
	userID := uuid.MustParse(params.XScUserID)
	followerUserID := uuid.MustParse(string(params.UserID))
	if err := h.app.DeleteFollower(ctx, userID, followerUserID); err != nil {
		return nil, err
	}
	return &api.FriendDeleteUserIDPutOK{}, nil
}

func (h *Handler) FriendSetUserIDPut(ctx context.Context, params api.FriendSetUserIDPutParams) (api.FriendSetUserIDPutRes, error) {
	userID := uuid.MustParse(params.XScUserID)
	followerUserID := uuid.MustParse(string(params.UserID))
	if err := h.app.SetFollower(ctx, userID, followerUserID); err != nil {
		return nil, err
	}
	return &api.FriendSetUserIDPutOK{}, nil
}
