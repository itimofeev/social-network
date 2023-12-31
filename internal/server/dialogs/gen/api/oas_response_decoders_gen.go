// Code generated by ogen, DO NOT EDIT.

package api

import (
	"io"
	"mime"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

func decodeDialogUserIDListGetResponse(resp *http.Response) (res DialogUserIDListGetRes, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response DialogUserIDListGetOKApplicationJSON
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			// Validate response.
			if err := func() error {
				if err := response.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, errors.Wrap(err, "validate")
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 401:
		// Code 401.
		return &R401{}, nil
	}
	// Convenient error response.
	defRes, err := func() (res *R5xxStatusCodeWithHeaders, err error) {
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response R5xx
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			var wrapper R5xxStatusCodeWithHeaders
			wrapper.Response = response
			wrapper.StatusCode = resp.StatusCode
			h := uri.NewHeaderDecoder(resp.Header)
			// Parse "Retry-After" header.
			{
				cfg := uri.HeaderParameterDecodingConfig{
					Name:    "Retry-After",
					Explode: false,
				}
				if err := func() error {
					if err := h.HasParam(cfg); err == nil {
						if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
							var wrapperDotRetryAfterVal int
							if err := func() error {
								val, err := d.DecodeValue()
								if err != nil {
									return err
								}

								c, err := conv.ToInt(val)
								if err != nil {
									return err
								}

								wrapperDotRetryAfterVal = c
								return nil
							}(); err != nil {
								return err
							}
							wrapper.RetryAfter.SetTo(wrapperDotRetryAfterVal)
							return nil
						}); err != nil {
							return err
						}
					}
					return nil
				}(); err != nil {
					return res, errors.Wrap(err, "parse Retry-After header")
				}
			}
			return &wrapper, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	}()
	if err != nil {
		return res, errors.Wrapf(err, "default (code %d)", resp.StatusCode)
	}
	return res, errors.Wrap(defRes, "error")
}

func decodeDialogUserIDSendPostResponse(resp *http.Response) (res DialogUserIDSendPostRes, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		return &DialogUserIDSendPostOK{}, nil
	case 401:
		// Code 401.
		return &R401{}, nil
	}
	// Convenient error response.
	defRes, err := func() (res *R5xxStatusCodeWithHeaders, err error) {
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response R5xx
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			var wrapper R5xxStatusCodeWithHeaders
			wrapper.Response = response
			wrapper.StatusCode = resp.StatusCode
			h := uri.NewHeaderDecoder(resp.Header)
			// Parse "Retry-After" header.
			{
				cfg := uri.HeaderParameterDecodingConfig{
					Name:    "Retry-After",
					Explode: false,
				}
				if err := func() error {
					if err := h.HasParam(cfg); err == nil {
						if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
							var wrapperDotRetryAfterVal int
							if err := func() error {
								val, err := d.DecodeValue()
								if err != nil {
									return err
								}

								c, err := conv.ToInt(val)
								if err != nil {
									return err
								}

								wrapperDotRetryAfterVal = c
								return nil
							}(); err != nil {
								return err
							}
							wrapper.RetryAfter.SetTo(wrapperDotRetryAfterVal)
							return nil
						}); err != nil {
							return err
						}
					}
					return nil
				}(); err != nil {
					return res, errors.Wrap(err, "parse Retry-After header")
				}
			}
			return &wrapper, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	}()
	if err != nil {
		return res, errors.Wrapf(err, "default (code %d)", resp.StatusCode)
	}
	return res, errors.Wrap(defRes, "error")
}
