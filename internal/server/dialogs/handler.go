package dialogs

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/ogen-go/ogen/ogenerrors"

	"github.com/itimofeev/social-network/internal/app/dialogs"
	"github.com/itimofeev/social-network/internal/server/dialogs/gen/api"
	"github.com/itimofeev/social-network/pkg/xcontext"
)

type Handler struct {
	app *dialogs.App
}

func (h *Handler) NewError(ctx context.Context, err error) *api.R5xxStatusCodeWithHeaders {
	requestID := xcontext.GetRequestID(ctx)
	return &api.R5xxStatusCodeWithHeaders{
		StatusCode: http.StatusInternalServerError,
		Response: api.R5xx{
			Message: err.Error(),
			RequestID: api.OptString{
				Value: requestID,
				Set:   requestID != "",
			},
			Code: api.NewOptInt(http.StatusInternalServerError),
		},
	}
}

func NewHandler(app *dialogs.App) *Handler {
	return &Handler{app: app}
}

func (h *Handler) ErrorHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	slog.WarnContext(ctx, "Error on handling request", "err", err)
	ogenerrors.DefaultErrorHandler(ctx, w, r, err)
}
