package recurly

import (
	"encoding/json"
	"net/http"
	"strings"
)

// Error contains basic information about the error
type Error struct {
	Message          string
	Type             ErrorType
	Params           ErrorParams
	TransactionError *TransactionError
	Response         ResponseMetadata
}

func (e *Error) Error() string {
	return e.Message
}

type ErrorType string

type TransactionErrorCategory string

const (
	ErrorTypeConnection  = ErrorType("connection")
	ErrorTypeNotFound    = ErrorType("not_found")
	ErrorTypeRateLimited = ErrorType("rate_limited")
	ErrorTypeTimeout     = ErrorType("timeout")

	TransactionErrorCategorySoft          = TransactionErrorCategory("soft")
	TransactionErrorCategoryFraud         = TransactionErrorCategory("fraud")
	TransactionErrorCategoryHard          = TransactionErrorCategory("hard")
	TransactionErrorCategoryCommunication = TransactionErrorCategory("communication")
	TransactionErrorCategoryUnknown       = TransactionErrorCategory("unknown")
)

type errorResponse struct {
	Error errorDetails `json:"error"`
}

type errorDetails struct {
	Type             string            `json:"type"`
	Message          string            `json:"message"`
	Params           ErrorParams       `json:"params"`
	TransactionError *TransactionError `json:"transaction_error"`
}

type TransactionError struct {
	TransactionID  string                   `json:"transaction_id"`
	Category       TransactionErrorCategory `json:"category"`
	Code           string                   `json:"code"`
	Message        string                   `json:"message"`
	MerchantAdvice string                   `json:"merchant_advice"`
}

// ErrorParams map of fields to error messages
type ErrorParams map[string]string

// parseResponseToError converts an http.Response to the appropriate error
func parseResponseToError(res *http.Response, body []byte) error {
	if strings.HasPrefix(res.Header.Get("Content-type"), "application/json") {
		// Return an error formatted from the JSON response

		var errResp errorResponse
		if err := json.Unmarshal(body, &errResp); err == nil {
			return &Error{
				Message:          errResp.Error.Message,
				Type:             ErrorType(errResp.Error.Type),
				Params:           errResp.Error.Params,
				TransactionError: errResp.Error.TransactionError,
				Response:         parseResponseMetadata(res),
			}
		}
	}

	errMessage := ""
	errType := ErrorTypeConnection

	switch res.StatusCode {
	case http.StatusUnauthorized: // 401
		errMessage = "Unauthorized: invalid API Key"
	case http.StatusForbidden: // 403
		errMessage = "Forbidden: The API key is not authorized for this resource"
	case http.StatusNotFound: // 404
		errMessage = "requested object or endpoint not found"
		errType = ErrorTypeNotFound
	case http.StatusUnprocessableEntity: // 422
		errMessage = "Invalid request"
		errType = ErrorTypeNotFound
	case http.StatusTooManyRequests: // 429
		errMessage = "You made too many API requests"
		errType = ErrorTypeRateLimited

	case http.StatusInternalServerError: // 500
		errMessage = "Server experienced an error"
	case http.StatusBadGateway: // 502
		errMessage = "Error contacting server"
	case http.StatusServiceUnavailable: // 503
		errMessage = "Service unavailable"
	case http.StatusRequestTimeout: // 408
	case http.StatusGatewayTimeout: // 504
		errMessage = "Request timed out"
		errType = ErrorTypeTimeout
	}

	return &Error{
		Message:  errMessage,
		Type:     errType,
		Response: parseResponseMetadata(res),
	}
}
