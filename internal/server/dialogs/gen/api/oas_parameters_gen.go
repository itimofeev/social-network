// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"net/url"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// DialogUserIDListGetParams is parameters of GET /dialog/{user_id}/list operation.
type DialogUserIDListGetParams struct {
	UserID    UserId
	XScUserID string
}

func unpackDialogUserIDListGetParams(packed middleware.Parameters) (params DialogUserIDListGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "user_id",
			In:   "path",
		}
		params.UserID = packed[key].(UserId)
	}
	{
		key := middleware.ParameterKey{
			Name: "X-Sc-User-Id",
			In:   "header",
		}
		params.XScUserID = packed[key].(string)
	}
	return params
}

func decodeDialogUserIDListGetParams(args [1]string, argsEscaped bool, r *http.Request) (params DialogUserIDListGetParams, _ error) {
	h := uri.NewHeaderDecoder(r.Header)
	// Decode path: user_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "user_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				var paramsDotUserIDVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotUserIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.UserID = UserId(paramsDotUserIDVal)
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "user_id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode header: X-Sc-User-Id.
	if err := func() error {
		cfg := uri.HeaderParameterDecodingConfig{
			Name:    "X-Sc-User-Id",
			Explode: false,
		}
		if err := h.HasParam(cfg); err == nil {
			if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.XScUserID = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "X-Sc-User-Id",
			In:   "header",
			Err:  err,
		}
	}
	return params, nil
}

// DialogUserIDSendPostParams is parameters of POST /dialog/{user_id}/send operation.
type DialogUserIDSendPostParams struct {
	UserID    UserId
	XScUserID string
}

func unpackDialogUserIDSendPostParams(packed middleware.Parameters) (params DialogUserIDSendPostParams) {
	{
		key := middleware.ParameterKey{
			Name: "user_id",
			In:   "path",
		}
		params.UserID = packed[key].(UserId)
	}
	{
		key := middleware.ParameterKey{
			Name: "X-Sc-User-Id",
			In:   "header",
		}
		params.XScUserID = packed[key].(string)
	}
	return params
}

func decodeDialogUserIDSendPostParams(args [1]string, argsEscaped bool, r *http.Request) (params DialogUserIDSendPostParams, _ error) {
	h := uri.NewHeaderDecoder(r.Header)
	// Decode path: user_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "user_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				var paramsDotUserIDVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotUserIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.UserID = UserId(paramsDotUserIDVal)
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "user_id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode header: X-Sc-User-Id.
	if err := func() error {
		cfg := uri.HeaderParameterDecodingConfig{
			Name:    "X-Sc-User-Id",
			Explode: false,
		}
		if err := h.HasParam(cfg); err == nil {
			if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.XScUserID = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "X-Sc-User-Id",
			In:   "header",
			Err:  err,
		}
	}
	return params, nil
}
