package server

import (
	"context"
	"errors"

	"github.com/itimofeev/social-network/internal/app"
	"github.com/itimofeev/social-network/internal/entity"
	"github.com/itimofeev/social-network/internal/gen/api"
)

type Handler struct {
	app *app.App
}

func (h *Handler) LoginPost(ctx context.Context, req api.OptLoginPostReq) (api.LoginPostRes, error) {
	if !req.Set {
		return &api.LoginPostBadRequest{}, nil
	}
	token, err := h.app.LoginUser(ctx, req.Value.ID, req.Value.Password)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrUserNotFound):
			return &api.LoginPostNotFound{}, nil
		case errors.Is(err, entity.ErrIncorrectPassword):
			return &api.LoginPostBadRequest{}, nil
		}
		return nil, err
	}

	return &api.LoginPostOK{Token: api.NewOptString(token)}, nil
}

func (h *Handler) UserRegisterPost(ctx context.Context, req api.OptUserRegisterPostReq) (api.UserRegisterPostRes, error) {
	if !req.Set {
		return &api.UserRegisterPostBadRequest{}, nil
	}

	user, err := h.app.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}

	return &api.UserRegisterPostOK{UserID: api.NewOptString(user.UserID)}, nil
}

func (h *Handler) UserGetIDGet(ctx context.Context, params api.UserGetIDGetParams) (api.UserGetIDGetRes, error) {
	user, err := h.app.GetUser(ctx, string(params.ID))
	if err != nil {
		if errors.Is(err, entity.ErrUserNotFound) {
			return &api.UserGetIDGetNotFound{}, nil
		}
		return nil, err
	}

	return convertUserToAPI(user), nil
}

func convertUserToAPI(user entity.User) *api.User {
	return &api.User{
		ID:         api.NewOptUserId(api.UserId(user.UserID)),
		FirstName:  api.NewOptString(user.FirstName),
		SecondName: api.NewOptString(user.SecondName),
		Birthdate:  api.NewOptBirthDate(api.BirthDate(user.BirthDate)),
		Biography:  api.NewOptString(user.Biography),
		City:       api.NewOptString(user.City),
	}
}

func convertUsersToAPI(user []entity.User) []api.User {
	res := make([]api.User, len(user))
	for i := range user {
		res[i] = *convertUserToAPI(user[i])
	}
	return res
}

func (h *Handler) UserSearchGet(ctx context.Context, params api.UserSearchGetParams) (api.UserSearchGetRes, error) {
	users, err := h.app.SearchUsers(ctx, params.FirstName, params.LastName)
	if err != nil {
		return nil, err
	}

	apiUsers := api.UserSearchGetOKApplicationJSON(convertUsersToAPI(users))
	return &apiUsers, nil
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

func NewHandler(app *app.App) *Handler {
	return &Handler{app: app}
}
