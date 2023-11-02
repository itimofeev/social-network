package server

import (
	"context"

	"github.com/itimofeev/social-network/internal/app"
	"github.com/itimofeev/social-network/internal/gen/api"
)

type Handler struct {
	app *app.App
}

func (h *Handler) HandleBearerAuth(ctx context.Context, operationName string, t api.BearerAuth) (context.Context, error) {
	panic("implement me")
}

func (h *Handler) DialogUserIDListGet(ctx context.Context, params api.DialogUserIDListGetParams) (api.DialogUserIDListGetRes, error) {
	panic("implement me")
}

func (h *Handler) DialogUserIDSendPost(ctx context.Context, req api.OptDialogUserIDSendPostReq, params api.DialogUserIDSendPostParams) (api.DialogUserIDSendPostRes, error) {
	panic("implement me")
}

func (h *Handler) FriendDeleteUserIDPut(ctx context.Context, params api.FriendDeleteUserIDPutParams) (api.FriendDeleteUserIDPutRes, error) {
	panic("implement me")
}

func (h *Handler) FriendSetUserIDPut(ctx context.Context, params api.FriendSetUserIDPutParams) (api.FriendSetUserIDPutRes, error) {
	panic("implement me")
}

func (h *Handler) LoginPost(ctx context.Context, req api.OptLoginPostReq) (api.LoginPostRes, error) {
	panic("implement me")
}

func (h *Handler) PostCreatePost(ctx context.Context, req api.OptPostCreatePostReq) (api.PostCreatePostRes, error) {
	panic("implement me")
}

func (h *Handler) PostDeleteIDPut(ctx context.Context, params api.PostDeleteIDPutParams) (api.PostDeleteIDPutRes, error) {
	panic("implement me")
}

func (h *Handler) PostFeedGet(ctx context.Context, params api.PostFeedGetParams) (api.PostFeedGetRes, error) {
	panic("implement me")
}

func (h *Handler) PostGetIDGet(ctx context.Context, params api.PostGetIDGetParams) (api.PostGetIDGetRes, error) {
	panic("implement me")
}

func (h *Handler) PostUpdatePut(ctx context.Context, req api.OptPostUpdatePutReq) (api.PostUpdatePutRes, error) {
	panic("implement me")
}

func (h *Handler) UserGetIDGet(ctx context.Context, params api.UserGetIDGetParams) (api.UserGetIDGetRes, error) {
	panic("implement me")
}

func (h *Handler) UserRegisterPost(ctx context.Context, req api.OptUserRegisterPostReq) (api.UserRegisterPostRes, error) {
	panic("implement me")
}

func (h *Handler) UserSearchGet(ctx context.Context, params api.UserSearchGetParams) (api.UserSearchGetRes, error) {
	panic("implement me")
}

func NewHandler(app *app.App) *Handler {
	return &Handler{app: app}
}
