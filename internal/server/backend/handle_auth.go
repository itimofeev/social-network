package backend

import (
	"context"

	"github.com/go-faster/errors"

	"github.com/itimofeev/social-network/internal/entity"
	"github.com/itimofeev/social-network/internal/server/backend/gen/api"
	"github.com/itimofeev/social-network/pkg/xcontext"
)

func (h *Handler) HandleBearerAuth(ctx context.Context, operationName string, t api.BearerAuth) (context.Context, error) {
	userID, err := h.app.CheckAuth(ctx, t.Token)
	if err != nil {
		return nil, err
	}

	return xcontext.WithUserID(ctx, userID), nil
}

func (h *Handler) AuthGet(ctx context.Context) (r api.AuthGetRes, _ error) {
	userID := xcontext.GetUserID(ctx)
	if userID == "" {
		return nil, errors.New("invalid token")
	}
	return &api.AuthGetOK{XScUserID: userID}, nil
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
