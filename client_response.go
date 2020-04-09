package recurly

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

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
	// RequestID uniquely identifies the Recurly API request for debugging
	RequestID string
	// Version indicates the response version
	Version string
}

func parseResponseMetadata(res *http.Response) ResponseMetadata {
	deprecated := res.Header.Get("Recurly-Deprecated") == "TRUE"

	return ResponseMetadata{
		Deprecated:      deprecated,
		DeprecationDate: res.Header.Get("Recurly-Sunset-Date"),
		RateLimit:       parseRateLimit(res),
		StatusCode:      res.StatusCode,
		RequestID:       res.Header.Get("X-Request-Id"),
		Version:         res.Header.Get("Recurly-Version"),
	}
}

func (meta ResponseMetadata) String() string {
	line := fmt.Sprintf("status code: %d, request ID: %s, version: %s, limit: %s",
		meta.StatusCode,
		meta.RequestID,
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
