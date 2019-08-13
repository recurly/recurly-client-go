package recurly

import (
	"context"
	"net/http"
	"strconv"
	"strings"
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
	IdempotencyKey string
	// Header contains additional request headers for unique requests
	Header http.Header
	// Context passed to the HTTP request for cancelling requests
	Context context.Context

	RequestParams
}

// URLParams contains additional URL parameters for querying generic lists
func (params *Params) URLParams() []KeyValue {
	if params.RequestParams != nil {
		return params.RequestParams.URLParams()
	}
	return []KeyValue{}
}

// ListParams inform a list query
type ListParams struct {
	Params

	IDs       []string
	Limit     int
	Order     ListOrder
	Sort      SortField
	BeginTime *time.Time
	EndTime   *time.Time
}

// URLParams contains additional URL parameters for querying generic lists
func (list *ListParams) URLParams() []KeyValue {
	var options []KeyValue

	if list.IDs != nil && len(list.IDs) > 0 {
		options = append(options, KeyValue{Key: "ids", Value: strings.Join(list.IDs, ",")})
	}
	if list.Limit > 0 {
		options = append(options, KeyValue{Key: "limit", Value: strconv.Itoa(list.Limit)})
	}
	options = append(options, KeyValue{Key: "order", Value: string(list.Order)})
	options = append(options, KeyValue{Key: "sort", Value: string(list.Sort)})
	if list.BeginTime != nil {
		options = append(options, KeyValue{Key: "begin_time", Value: formatTime(*list.BeginTime)})
	}
	if list.EndTime != nil {
		options = append(options, KeyValue{Key: "end_time", Value: formatTime(*list.EndTime)})
	}

	return options
}

func (list *ListParams) toParams() *Params {
	return &Params{
		IdempotencyKey: list.IdempotencyKey,
		Header:         list.Header,
		Context:        list.Context,
		RequestParams:  list,
	}
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
