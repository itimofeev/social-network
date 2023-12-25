package xmw

import (
	"net/http"

	"github.com/itimofeev/social-network/pkg/xcontext"
)

type RequestIDClient struct {
	defaultTransport http.RoundTripper
}

func NewRequestIDClient() *RequestIDClient {
	return &RequestIDClient{
		defaultTransport: http.DefaultTransport,
	}
}

func (h RequestIDClient) Do(r *http.Request) (*http.Response, error) {
	r.Header.Set(requestIDHeader, xcontext.GetRequestID(r.Context()))
	return h.defaultTransport.RoundTrip(r)
}
