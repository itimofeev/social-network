package backend

import (
	"context"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/itimofeev/social-network/internal/server/backend/gen/api"
)

func (h *Handler) PostCreatePost(ctx context.Context, req *api.PostCreatePostReq, params api.PostCreatePostParams) (api.PostCreatePostRes, error) {
	userID := uuid.MustParse(params.XScUserID)
	postID, err := h.app.CreatePost(ctx, userID, string(req.Text))
	if err != nil {
		return nil, err
	}
	return lo.ToPtr(api.PostId(postID)), nil
}

func (h *Handler) PostDeleteIDPut(ctx context.Context, params api.PostDeleteIDPutParams) (api.PostDeleteIDPutRes, error) {
	panic("implement me")
}

func (h *Handler) PostGetIDGet(ctx context.Context, params api.PostGetIDGetParams) (api.PostGetIDGetRes, error) {
	panic("implement me")
}

func (h *Handler) PostUpdatePut(ctx context.Context, req *api.PostUpdatePutReq, params api.PostUpdatePutParams) (api.PostUpdatePutRes, error) {
	panic("implement me")
}
