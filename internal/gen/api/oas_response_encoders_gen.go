// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/uri"
)

func encodeDialogUserIDListGetResponse(response DialogUserIDListGetRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *DialogUserIDListGetOKApplicationJSON:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *R400:
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		return nil

	case *R401:
		w.WriteHeader(401)
		span.SetStatus(codes.Error, http.StatusText(401))

		return nil

	case *DialogUserIDListGetInternalServerError:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Retry-After" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Retry-After",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.RetryAfter.Get(); ok {
						return e.EncodeValue(conv.IntToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Retry-After header")
				}
			}
		}
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		e := jx.GetEncoder()
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *DialogUserIDListGetServiceUnavailable:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Retry-After" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Retry-After",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.RetryAfter.Get(); ok {
						return e.EncodeValue(conv.IntToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Retry-After header")
				}
			}
		}
		w.WriteHeader(503)
		span.SetStatus(codes.Error, http.StatusText(503))

		e := jx.GetEncoder()
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeDialogUserIDSendPostResponse(response DialogUserIDSendPostRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *DialogUserIDSendPostOK:
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		return nil

	case *R400:
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		return nil

	case *R401:
		w.WriteHeader(401)
		span.SetStatus(codes.Error, http.StatusText(401))

		return nil

	case *DialogUserIDSendPostInternalServerError:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Retry-After" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Retry-After",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.RetryAfter.Get(); ok {
						return e.EncodeValue(conv.IntToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Retry-After header")
				}
			}
		}
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		e := jx.GetEncoder()
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *DialogUserIDSendPostServiceUnavailable:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Retry-After" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Retry-After",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.RetryAfter.Get(); ok {
						return e.EncodeValue(conv.IntToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Retry-After header")
				}
			}
		}
		w.WriteHeader(503)
		span.SetStatus(codes.Error, http.StatusText(503))

		e := jx.GetEncoder()
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeFriendDeleteUserIDPutResponse(response FriendDeleteUserIDPutRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *FriendDeleteUserIDPutOK:
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		return nil

	case *R400:
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		return nil

	case *R401:
		w.WriteHeader(401)
		span.SetStatus(codes.Error, http.StatusText(401))

		return nil

	case *FriendDeleteUserIDPutInternalServerError:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Retry-After" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Retry-After",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.RetryAfter.Get(); ok {
						return e.EncodeValue(conv.IntToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Retry-After header")
				}
			}
		}
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		e := jx.GetEncoder()
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *FriendDeleteUserIDPutServiceUnavailable:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Retry-After" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Retry-After",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.RetryAfter.Get(); ok {
						return e.EncodeValue(conv.IntToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Retry-After header")
				}
			}
		}
		w.WriteHeader(503)
		span.SetStatus(codes.Error, http.StatusText(503))

		e := jx.GetEncoder()
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeFriendSetUserIDPutResponse(response FriendSetUserIDPutRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *FriendSetUserIDPutOK:
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		return nil

	case *R400:
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		return nil

	case *R401:
		w.WriteHeader(401)
		span.SetStatus(codes.Error, http.StatusText(401))

		return nil

	case *FriendSetUserIDPutInternalServerError:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Retry-After" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Retry-After",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.RetryAfter.Get(); ok {
						return e.EncodeValue(conv.IntToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Retry-After header")
				}
			}
		}
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		e := jx.GetEncoder()
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *FriendSetUserIDPutServiceUnavailable:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Retry-After" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Retry-After",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.RetryAfter.Get(); ok {
						return e.EncodeValue(conv.IntToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Retry-After header")
				}
			}
		}
		w.WriteHeader(503)
		span.SetStatus(codes.Error, http.StatusText(503))

		e := jx.GetEncoder()
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeLoginPostResponse(response LoginPostRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *LoginPostOK:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *LoginPostBadRequest:
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		return nil

	case *LoginPostNotFound:
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		return nil

	case *LoginPostInternalServerError:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Retry-After" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Retry-After",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.RetryAfter.Get(); ok {
						return e.EncodeValue(conv.IntToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Retry-After header")
				}
			}
		}
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		e := jx.GetEncoder()
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *LoginPostServiceUnavailable:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Retry-After" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Retry-After",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.RetryAfter.Get(); ok {
						return e.EncodeValue(conv.IntToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Retry-After header")
				}
			}
		}
		w.WriteHeader(503)
		span.SetStatus(codes.Error, http.StatusText(503))

		e := jx.GetEncoder()
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodePostCreatePostResponse(response PostCreatePostRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PostId:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *R400:
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		return nil

	case *R401:
		w.WriteHeader(401)
		span.SetStatus(codes.Error, http.StatusText(401))

		return nil

	case *PostCreatePostInternalServerError:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Retry-After" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Retry-After",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.RetryAfter.Get(); ok {
						return e.EncodeValue(conv.IntToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Retry-After header")
				}
			}
		}
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		e := jx.GetEncoder()
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *PostCreatePostServiceUnavailable:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Retry-After" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Retry-After",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.RetryAfter.Get(); ok {
						return e.EncodeValue(conv.IntToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Retry-After header")
				}
			}
		}
		w.WriteHeader(503)
		span.SetStatus(codes.Error, http.StatusText(503))

		e := jx.GetEncoder()
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodePostDeleteIDPutResponse(response PostDeleteIDPutRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PostDeleteIDPutOK:
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		return nil

	case *R400:
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		return nil

	case *R401:
		w.WriteHeader(401)
		span.SetStatus(codes.Error, http.StatusText(401))

		return nil

	case *PostDeleteIDPutInternalServerError:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Retry-After" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Retry-After",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.RetryAfter.Get(); ok {
						return e.EncodeValue(conv.IntToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Retry-After header")
				}
			}
		}
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		e := jx.GetEncoder()
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *PostDeleteIDPutServiceUnavailable:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Retry-After" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Retry-After",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.RetryAfter.Get(); ok {
						return e.EncodeValue(conv.IntToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Retry-After header")
				}
			}
		}
		w.WriteHeader(503)
		span.SetStatus(codes.Error, http.StatusText(503))

		e := jx.GetEncoder()
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodePostFeedGetResponse(response PostFeedGetRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PostFeedGetOKApplicationJSON:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *R400:
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		return nil

	case *R401:
		w.WriteHeader(401)
		span.SetStatus(codes.Error, http.StatusText(401))

		return nil

	case *PostFeedGetInternalServerError:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Retry-After" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Retry-After",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.RetryAfter.Get(); ok {
						return e.EncodeValue(conv.IntToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Retry-After header")
				}
			}
		}
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		e := jx.GetEncoder()
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *PostFeedGetServiceUnavailable:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Retry-After" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Retry-After",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.RetryAfter.Get(); ok {
						return e.EncodeValue(conv.IntToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Retry-After header")
				}
			}
		}
		w.WriteHeader(503)
		span.SetStatus(codes.Error, http.StatusText(503))

		e := jx.GetEncoder()
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodePostGetIDGetResponse(response PostGetIDGetRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *Post:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *R400:
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		return nil

	case *R401:
		w.WriteHeader(401)
		span.SetStatus(codes.Error, http.StatusText(401))

		return nil

	case *PostGetIDGetInternalServerError:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Retry-After" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Retry-After",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.RetryAfter.Get(); ok {
						return e.EncodeValue(conv.IntToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Retry-After header")
				}
			}
		}
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		e := jx.GetEncoder()
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *PostGetIDGetServiceUnavailable:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Retry-After" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Retry-After",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.RetryAfter.Get(); ok {
						return e.EncodeValue(conv.IntToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Retry-After header")
				}
			}
		}
		w.WriteHeader(503)
		span.SetStatus(codes.Error, http.StatusText(503))

		e := jx.GetEncoder()
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodePostUpdatePutResponse(response PostUpdatePutRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PostUpdatePutOK:
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		return nil

	case *R400:
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		return nil

	case *R401:
		w.WriteHeader(401)
		span.SetStatus(codes.Error, http.StatusText(401))

		return nil

	case *PostUpdatePutInternalServerError:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Retry-After" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Retry-After",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.RetryAfter.Get(); ok {
						return e.EncodeValue(conv.IntToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Retry-After header")
				}
			}
		}
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		e := jx.GetEncoder()
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *PostUpdatePutServiceUnavailable:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Retry-After" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Retry-After",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.RetryAfter.Get(); ok {
						return e.EncodeValue(conv.IntToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Retry-After header")
				}
			}
		}
		w.WriteHeader(503)
		span.SetStatus(codes.Error, http.StatusText(503))

		e := jx.GetEncoder()
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeUserGetIDGetResponse(response UserGetIDGetRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *User:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *UserGetIDGetBadRequest:
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		return nil

	case *UserGetIDGetNotFound:
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		return nil

	case *UserGetIDGetInternalServerError:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Retry-After" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Retry-After",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.RetryAfter.Get(); ok {
						return e.EncodeValue(conv.IntToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Retry-After header")
				}
			}
		}
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		e := jx.GetEncoder()
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *UserGetIDGetServiceUnavailable:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Retry-After" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Retry-After",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.RetryAfter.Get(); ok {
						return e.EncodeValue(conv.IntToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Retry-After header")
				}
			}
		}
		w.WriteHeader(503)
		span.SetStatus(codes.Error, http.StatusText(503))

		e := jx.GetEncoder()
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeUserRegisterPostResponse(response UserRegisterPostRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *UserRegisterPostOK:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *UserRegisterPostBadRequest:
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		return nil

	case *UserRegisterPostInternalServerError:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Retry-After" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Retry-After",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.RetryAfter.Get(); ok {
						return e.EncodeValue(conv.IntToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Retry-After header")
				}
			}
		}
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		e := jx.GetEncoder()
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *UserRegisterPostServiceUnavailable:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Retry-After" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Retry-After",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.RetryAfter.Get(); ok {
						return e.EncodeValue(conv.IntToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Retry-After header")
				}
			}
		}
		w.WriteHeader(503)
		span.SetStatus(codes.Error, http.StatusText(503))

		e := jx.GetEncoder()
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeUserSearchGetResponse(response UserSearchGetRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *UserSearchGetOKApplicationJSON:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *UserSearchGetBadRequest:
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		return nil

	case *UserSearchGetInternalServerError:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Retry-After" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Retry-After",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.RetryAfter.Get(); ok {
						return e.EncodeValue(conv.IntToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Retry-After header")
				}
			}
		}
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		e := jx.GetEncoder()
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *UserSearchGetServiceUnavailable:
		w.Header().Set("Content-Type", "application/json")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Retry-After" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Retry-After",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.RetryAfter.Get(); ok {
						return e.EncodeValue(conv.IntToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Retry-After header")
				}
			}
		}
		w.WriteHeader(503)
		span.SetStatus(codes.Error, http.StatusText(503))

		e := jx.GetEncoder()
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}
