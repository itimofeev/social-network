package dialogs

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/ogen-go/ogen/ogenerrors"

	"github.com/itimofeev/social-network/internal/app/dialogs"
	"github.com/itimofeev/social-network/internal/server/dialogs/gen/api"
)

type Handler struct {
	app *dialogs.App
}

func (h *Handler) NewError(ctx context.Context, err error) *api.R5xxStatusCodeWithHeaders {
	//TODO implement me
	panic("implement me")
}

func NewHandler(app *dialogs.App) *Handler {
	return &Handler{app: app}
}

func (h *Handler) ErrorHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	slog.Warn("Error on handling request", "err", err)
	ogenerrors.DefaultErrorHandler(ctx, w, r, err)
}
