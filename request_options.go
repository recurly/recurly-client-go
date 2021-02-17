package recurly

import (
	"net/http"
)

// RequestOptions contains additional request parameters for the API client
type RequestOptions struct {
	// IdempotencyKey used to prevent duplicate requests
	IdempotencyKey string `json:"-"`
	// Header contains additional request headers for unique requests
	Header http.Header `json:"-"`
}

func (o *RequestOptions) clone() *RequestOptions {
	clone := &RequestOptions{}
	clone.Header = o.Header.Clone()
	clone.IdempotencyKey = o.IdempotencyKey

	return clone
}

// Option allows for customizations to a request.
type Option func(o *RequestOptions)

// NewRequestOptions will create a new RequestOptions with the passed Options applied to it.
func NewRequestOptions(opts ...Option) *RequestOptions {
	o := &RequestOptions{}
	return o.ApplyOptions(opts...)
}

// RequestOptionsFromParams will create a RequestOptions from a Params
func RequestOptionsFromParams(p Params, opts ...Option) *RequestOptions {
	o := &RequestOptions{}
	o.Header = p.Header.Clone()
	o.IdempotencyKey = p.IdempotencyKey

	return o.ApplyOptions(opts...)
}

// ApplyOptions will apply the passed Options to the RequestOptions.
func (o *RequestOptions) ApplyOptions(opts ...Option) *RequestOptions {
	cOpts := o.clone()
	for _, opt := range opts {
		opt(cOpts)
	}
	return cOpts
}

// WithHeader provides the capability to add custom headers to an operation.
func WithHeader(h http.Header) Option {
	return func(o *RequestOptions) {
		o.Header = h
	}
}

// WithIdempotencyKey provides the capability to add an Idempotency Key to an operation.
func WithIdempotencyKey(k string) Option {
	return func(o *RequestOptions) {
		o.IdempotencyKey = k
	}
}

type OptionsApplier interface {
	applyOptions(req *http.Request) *http.Request
}

func (o *RequestOptions) applyOptions(req *http.Request) *http.Request {
	for key, v := range o.Header {
		for _, value := range v {
			req.Header.Set(key, value)
		}
	}
	// TODO: generate an idempotency key if missing?
	if o.IdempotencyKey != "" {
		req.Header.Add("Idempotency-Key", o.IdempotencyKey)
	}

	return req
}
