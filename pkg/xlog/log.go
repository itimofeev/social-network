package xlog

import (
	"context"
	"log/slog"
	"os"

	"github.com/itimofeev/social-network/pkg/xcontext"
)

type RequestIDHandler struct {
	slog.Handler
}

func (h RequestIDHandler) Handle(ctx context.Context, r slog.Record) error {
	if requestID := xcontext.GetRequestID(ctx); requestID != "" {
		r.Add("request_id", slog.StringValue(requestID))
	}

	return h.Handler.Handle(ctx, r)
}

func InitSlog() {
	var handler slog.Handler
	handler = slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	})

	handler = RequestIDHandler{Handler: handler}

	slog.SetDefault(slog.New(handler))
}
