package recurly

import (
	"context"
	"net/http"
	"time"
)

type GenericParams interface {
	toParams() *Params
}

type RequestParams interface {
	// URLParams returns a list of key/value pairs to URL encode in the request URL
	URLParams() []KeyValue
}

// Params contains additional request parameters for the API client
type Params struct {
	// IdempotencyKey used to prevent duplicate requests
	IdempotencyKey string `json:"-"`
	// Header contains additional request headers for unique requests
	Header http.Header `json:"-"`
	// Context passed to the HTTP request for cancelling requests
	Context context.Context `json:"-"`

	RequestParams `json:"-"`

	// Data contains an object to be JSON encoded and sent in the body of the request
	Data interface{} `json:"-"`
}

// URLParams contains additional URL parameters for querying generic lists
func (params *Params) URLParams() []KeyValue {
	if params.RequestParams != nil {
		return params.RequestParams.URLParams()
	}
	return []KeyValue{}
}

type KeyValue struct {
	Key   string
	Value string
}

// ListOrder options for ordering a list
type ListOrder string

// SortField specifies a field to sort against
type SortField string

const (
	// ListAscending sorts results in ascending order. When sorting by UpdatedAt, you really only
	// want to return results in ascending order.
	ListAscending ListOrder = "asc"
	// ListDescending sorts results in descending order
	ListDescending ListOrder = "desc"

	// SortDefault sorts results by their default field
	SortDefault SortField = ""
	// SortCreatedAt sorts results by their created_at datetime
	SortCreatedAt SortField = "created_at"
	// SortUpdatedAt sorts results by their updated_at datetime. This should only be used with ListAscending.
	SortUpdatedAt SortField = "updated_at"
)

// formatBool converts a boolean to a string
func formatBool(v bool) string {
	if v {
		return "true"
	}
	return "false"
}

// formatTime returns an ISO8601 timestamp
func formatTime(t time.Time) string {
	return t.UTC().Format("2006-01-02T15:04:05Z")
}
