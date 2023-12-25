package xmw

import (
	"net/http"

	"github.com/itimofeev/social-network/pkg/xcontext"
)

const requestIDHeader = "X-Request-Id"

func RequestID(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		requestID := r.Header.Get(requestIDHeader)
		if requestID != "" {
			ctx = xcontext.WithRequestID(ctx, requestID)
			w.Header().Set("X-Request-Id", requestID)
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
