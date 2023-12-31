package backend

import (
	"context"

	"github.com/itimofeev/social-network/internal/server/backend/gen/api"
)

func (h *Handler) PostCreatePost(ctx context.Context, req api.OptPostCreatePostReq) (api.PostCreatePostRes, error) {
	panic("implement me")
}

func (h *Handler) PostDeleteIDPut(ctx context.Context, params api.PostDeleteIDPutParams) (api.PostDeleteIDPutRes, error) {
	panic("implement me")
}

func (h *Handler) PostGetIDGet(ctx context.Context, params api.PostGetIDGetParams) (api.PostGetIDGetRes, error) {
	panic("implement me")
}

func (h *Handler) PostUpdatePut(ctx context.Context, req api.OptPostUpdatePutReq) (api.PostUpdatePutRes, error) {
	panic("implement me")
}
