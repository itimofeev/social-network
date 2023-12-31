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
	if err := h.app.DeletePost(ctx, string(params.ID)); err != nil {
		return nil, err
	}
	return &api.PostDeleteIDPutOK{}, nil
}

func (h *Handler) PostGetIDGet(ctx context.Context, params api.PostGetIDGetParams) (api.PostGetIDGetRes, error) {
	post, err := h.app.GetPost(ctx, string(params.ID))
	if err != nil {
		return nil, err
	}

	return &api.Post{
		ID:           api.PostId(post.ID),
		Text:         api.PostText(post.Text),
		AuthorUserID: api.UserId(post.AuthorUserID),
	}, nil
}

func (h *Handler) PostUpdatePut(ctx context.Context, req *api.PostUpdatePutReq, params api.PostUpdatePutParams) (api.PostUpdatePutRes, error) {
	if err := h.app.UpdatePost(ctx, string(req.ID), string(req.Text)); err != nil {
		return nil, err
	}
	return &api.PostUpdatePutOK{}, nil
}
