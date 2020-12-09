package recurly

import "time"

// The QueryParams interface is used to identify optional query string parameters.
type QueryParams interface {
	// URLParams returns a list of key/value pairs to URL encode in the request URL
	URLParams() []KeyValue
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
