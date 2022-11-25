// Code generated by ogen, DO NOT EDIT.

package api

import (
	"io"
	"mime"
	"net/http"
	"net/url"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.uber.org/multierr"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

func (s *Server) decodeNullableStringsRequest(r *http.Request) (
	req NilString,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = multierr.Append(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}

		var request NilString
		buf, err := io.ReadAll(r.Body)
		if err != nil {
			return req, close, err
		}

		if len(buf) == 0 {
			return req, close, validate.ErrBodyRequired
		}

		d := jx.DecodeBytes(buf)
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"application/json\"")
		}
		if err := d.Skip(); err != io.EOF {
			return req, close, errors.New("unexpected trailing data")
		}
		if err := func() error {
			if err := (validate.String{
				MinLength:    0,
				MinLengthSet: false,
				MaxLength:    0,
				MaxLengthSet: false,
				Email:        false,
				Hostname:     false,
				Regex:        regexMap["^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$"],
			}).Validate(string(request.Value)); err != nil {
				return errors.Wrap(err, "string")
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "validate")
		}
		return request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeObjectsWithConflictingArrayPropertyRequest(r *http.Request) (
	req ObjectsWithConflictingArrayPropertyReq,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = multierr.Append(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}

		var request ObjectsWithConflictingArrayPropertyReq
		buf, err := io.ReadAll(r.Body)
		if err != nil {
			return req, close, err
		}

		if len(buf) == 0 {
			return req, close, validate.ErrBodyRequired
		}

		d := jx.DecodeBytes(buf)
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"application/json\"")
		}
		if err := d.Skip(); err != io.EOF {
			return req, close, errors.New("unexpected trailing data")
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "validate")
		}
		return request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeObjectsWithConflictingPropertiesRequest(r *http.Request) (
	req ObjectsWithConflictingPropertiesReq,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = multierr.Append(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}

		var request ObjectsWithConflictingPropertiesReq
		buf, err := io.ReadAll(r.Body)
		if err != nil {
			return req, close, err
		}

		if len(buf) == 0 {
			return req, close, validate.ErrBodyRequired
		}

		d := jx.DecodeBytes(buf)
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"application/json\"")
		}
		if err := d.Skip(); err != io.EOF {
			return req, close, errors.New("unexpected trailing data")
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "validate")
		}
		return request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeReferencedAllofRequest(r *http.Request) (
	req ReferencedAllofReq,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = multierr.Append(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}

		var request ReferencedAllofApplicationJSON
		buf, err := io.ReadAll(r.Body)
		if err != nil {
			return req, close, err
		}

		if len(buf) == 0 {
			return req, close, validate.ErrBodyRequired
		}

		d := jx.DecodeBytes(buf)
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"application/json\"")
		}
		if err := d.Skip(); err != io.EOF {
			return req, close, errors.New("unexpected trailing data")
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "validate")
		}
		return &request, close, nil
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		form := url.Values(r.MultipartForm.Value)

		var request ReferencedAllofMultipartFormData
		{
			var unwrapped Robot
			q := uri.NewQueryDecoder(form)
			{
				cfg := uri.QueryParameterDecodingConfig{
					Name:    "state",
					Style:   uri.QueryStyleForm,
					Explode: true,
				}
				if err := q.HasParam(cfg); err == nil {
					if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						unwrapped.State = RobotState(c)
						return nil
					}); err != nil {
						return req, close, errors.Wrap(err, "decode \"state\"")
					}
					if err := func() error {
						if err := unwrapped.State.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						return req, close, errors.Wrap(err, "validate")
					}
				} else {
					return req, close, errors.Wrap(err, "query")
				}
			}
			{
				cfg := uri.QueryParameterDecodingConfig{
					Name:    "id",
					Style:   uri.QueryStyleForm,
					Explode: true,
				}
				if err := q.HasParam(cfg); err == nil {
					if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToUUID(val)
						if err != nil {
							return err
						}

						unwrapped.ID = c
						return nil
					}); err != nil {
						return req, close, errors.Wrap(err, "decode \"id\"")
					}
				} else {
					return req, close, errors.Wrap(err, "query")
				}
			}
			{
				cfg := uri.QueryParameterDecodingConfig{
					Name:    "location",
					Style:   uri.QueryStyleForm,
					Explode: true,
					Fields:  []uri.QueryParameterObjectField{{"lat", true}, {"lon", true}},
				}
				if err := q.HasParam(cfg); err == nil {
					if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
						return unwrapped.Location.DecodeURI(d)
					}); err != nil {
						return req, close, errors.Wrap(err, "decode \"location\"")
					}
					if err := func() error {
						if err := unwrapped.Location.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						return req, close, errors.Wrap(err, "validate")
					}
				} else {
					return req, close, errors.Wrap(err, "query")
				}
			}
			request = ReferencedAllofMultipartFormData(unwrapped)
		}
		return &request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeReferencedAllofOptionalRequest(r *http.Request) (
	req ReferencedAllofOptionalReq,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = multierr.Append(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()
	req = &ReferencedAllofOptionalReqEmptyBody{}
	if _, ok := r.Header["Content-Type"]; !ok && r.ContentLength == 0 {
		return req, close, nil
	}
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		if r.ContentLength == 0 {
			return req, close, nil
		}

		var request ReferencedAllofOptionalApplicationJSON
		buf, err := io.ReadAll(r.Body)
		if err != nil {
			return req, close, err
		}

		if len(buf) == 0 {
			return req, close, nil
		}

		d := jx.DecodeBytes(buf)
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"application/json\"")
		}
		if err := d.Skip(); err != io.EOF {
			return req, close, errors.New("unexpected trailing data")
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "validate")
		}
		return &request, close, nil
	case ct == "multipart/form-data":
		if r.ContentLength == 0 {
			return req, close, nil
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, close, errors.Wrap(err, "parse multipart form")
		}
		// Remove all temporary files created by ParseMultipartForm when the request is done.
		//
		// Notice that the closers are called in reverse order, to match defer behavior, so
		// any opened file will be closed before RemoveAll call.
		closers = append(closers, r.MultipartForm.RemoveAll)
		form := url.Values(r.MultipartForm.Value)

		var request ReferencedAllofOptionalMultipartFormData
		{
			var unwrapped Robot
			q := uri.NewQueryDecoder(form)
			{
				cfg := uri.QueryParameterDecodingConfig{
					Name:    "state",
					Style:   uri.QueryStyleForm,
					Explode: true,
				}
				if err := q.HasParam(cfg); err == nil {
					if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						unwrapped.State = RobotState(c)
						return nil
					}); err != nil {
						return req, close, errors.Wrap(err, "decode \"state\"")
					}
					if err := func() error {
						if err := unwrapped.State.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						return req, close, errors.Wrap(err, "validate")
					}
				} else {
					return req, close, errors.Wrap(err, "query")
				}
			}
			{
				cfg := uri.QueryParameterDecodingConfig{
					Name:    "id",
					Style:   uri.QueryStyleForm,
					Explode: true,
				}
				if err := q.HasParam(cfg); err == nil {
					if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToUUID(val)
						if err != nil {
							return err
						}

						unwrapped.ID = c
						return nil
					}); err != nil {
						return req, close, errors.Wrap(err, "decode \"id\"")
					}
				} else {
					return req, close, errors.Wrap(err, "query")
				}
			}
			{
				cfg := uri.QueryParameterDecodingConfig{
					Name:    "location",
					Style:   uri.QueryStyleForm,
					Explode: true,
					Fields:  []uri.QueryParameterObjectField{{"lat", true}, {"lon", true}},
				}
				if err := q.HasParam(cfg); err == nil {
					if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
						return unwrapped.Location.DecodeURI(d)
					}); err != nil {
						return req, close, errors.Wrap(err, "decode \"location\"")
					}
					if err := func() error {
						if err := unwrapped.Location.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						return req, close, errors.Wrap(err, "validate")
					}
				} else {
					return req, close, errors.Wrap(err, "query")
				}
			}
			request = ReferencedAllofOptionalMultipartFormData(unwrapped)
		}
		return &request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeSimpleIntegerRequest(r *http.Request) (
	req int,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = multierr.Append(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}

		var request int
		buf, err := io.ReadAll(r.Body)
		if err != nil {
			return req, close, err
		}

		if len(buf) == 0 {
			return req, close, validate.ErrBodyRequired
		}

		d := jx.DecodeBytes(buf)
		if err := func() error {
			v, err := d.Int()
			request = int(v)
			if err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"application/json\"")
		}
		if err := d.Skip(); err != io.EOF {
			return req, close, errors.New("unexpected trailing data")
		}
		if err := func() error {
			if err := (validate.Int{
				MinSet:        true,
				Min:           -5,
				MaxSet:        true,
				Max:           5,
				MinExclusive:  false,
				MaxExclusive:  false,
				MultipleOfSet: false,
				MultipleOf:    0,
			}).Validate(int64(request)); err != nil {
				return errors.Wrap(err, "int")
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "validate")
		}
		return request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeSimpleObjectsRequest(r *http.Request) (
	req SimpleObjectsReq,
	close func() error,
	rerr error,
) {
	var closers []func() error
	close = func() error {
		var merr error
		// Close in reverse order, to match defer behavior.
		for i := len(closers) - 1; i >= 0; i-- {
			c := closers[i]
			merr = multierr.Append(merr, c())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()
	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch {
	case ct == "application/json":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}

		var request SimpleObjectsReq
		buf, err := io.ReadAll(r.Body)
		if err != nil {
			return req, close, err
		}

		if len(buf) == 0 {
			return req, close, validate.ErrBodyRequired
		}

		d := jx.DecodeBytes(buf)
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"application/json\"")
		}
		if err := d.Skip(); err != io.EOF {
			return req, close, errors.New("unexpected trailing data")
		}
		return request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}