package recurly

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type RequestMetadata struct {
	// ID is the request id assigned by Recurly to this request. Save this for support and debugging purposes.
	ID string
	// URL is the URL created for the request
	URL *url.URL
	// Method is the HTTP method used for the request. Ex (GET, POST, etc)
	Method string
}

// ResponseMetadata is the response from Recurly's API
type ResponseMetadata struct {
	// Deprecated is true if the version or endpoint is now deprecated and should not be used.
	Deprecated bool
	// DeprecationDate indicates the version or endpoint will sunset and should not be used after this date.
	DeprecationDate string
	// RateLimit indicates the remaining API requests before the rate limit is exceeded.
	RateLimit RateLimit
	// StatusCode contains the response's HTTP status code
	StatusCode int
	// Version indicates the response version
	Version string
	// TotalRecords is the header corresponding to the `Recurly-Total-Records`. The count of records matching this pager.
	TotalRecords *int64
	// Request is the metadata describing the request for this response
	Request RequestMetadata
}

func parseIntPtr(str string) *int64 {
	total, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return nil
	}
	return &total
}

func parseResponseMetadata(res *http.Response) *ResponseMetadata {
	deprecated := res.Header.Get("Recurly-Deprecated") == "TRUE"
	total := parseIntPtr(res.Header.Get("Recurly-Total-Records"))
	return &ResponseMetadata{
		Deprecated:      deprecated,
		DeprecationDate: res.Header.Get("Recurly-Sunset-Date"),
		RateLimit:       parseRateLimit(res),
		StatusCode:      res.StatusCode,
		Version:         res.Header.Get("Recurly-Version"),
		TotalRecords:    total,
		Request: RequestMetadata{
			ID:     res.Header.Get("X-Request-Id"),
			Method: res.Request.Method,
			URL:    res.Request.URL,
		},
	}
}

func (meta *ResponseMetadata) String() string {
	line := fmt.Sprintf("status code: %d, request ID: %s, version: %s, limit: %s",
		meta.StatusCode,
		meta.Request.ID,
		meta.Version,
		meta.RateLimit,
	)
	if meta.Deprecated {
		line = "DEPRECATED! " + line
	}
	if meta.DeprecationDate != "" {
		line = fmt.Sprintf("%s sunset date: %s", line, meta.DeprecationDate)
	}
	return line
}

// RateLimit contains information about the number of API requests available within the rate limit
type RateLimit struct {
	Limit          int
	Remaining      int
	resetTimestamp int64
}

// String formats the rate limit
func (limit RateLimit) String() string {
	return fmt.Sprintf("%d remaining of %d", limit.Remaining, limit.Limit)
}

// ResetDate return the date when the API rate limit fully resets
// See https://dev.recurly.com/docs/rate-limits for more information.
func (limit RateLimit) ResetDate() *time.Time {
	if limit.resetTimestamp == 0 {
		return nil
	}

	reset := time.Unix(limit.resetTimestamp, 0)
	return &reset
}

func parseRateLimit(res *http.Response) RateLimit {
	limitStr := res.Header.Get("X-RateLimit-Limit")
	if limitStr == "" {
		return RateLimit{}
	}
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		return RateLimit{}
	}
	remaining, _ := strconv.ParseInt(res.Header.Get("X-RateLimit-Remaining"), 10, 64)
	reset, _ := strconv.ParseInt(res.Header.Get("X-RateLimit-Reset"), 10, 64)

	return RateLimit{
		Limit:          int(limit),
		Remaining:      int(remaining),
		resetTimestamp: reset,
	}
}
