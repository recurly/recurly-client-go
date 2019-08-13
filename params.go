package recurly

import "net/http"

// Params contains additional request parameters for the API client
type Params interface {
	// IdempotencyKey used to prevent duplicate requests
	IdempotencyKey() string
	// Headers contains additional request headers for unique requests
	Headers() http.Header
	// URLParams contains additional URL parameters for querying lists
	URLParams() map[string]interface{}
}
