package recurly

import (
	"encoding/json"
	"net/http"
	"strings"
)

// Error contains basic information about the error
type Error struct {
	Message          string
	Class            ErrorClass
	Type             ErrorType
	Params           []ErrorParam
	TransactionError *TransactionError
	Response         ResponseMetadata
}

func (e *Error) Error() string {
	return e.Message
}

type ErrorType string
type ErrorClass string

type TransactionErrorCategory string

const (
	ErrorClassServer = ErrorClass("server")
	ErrorClassClient = ErrorClass("client")

	ErrorTypeUnknown                 = ErrorType("unknown")
	ErrorTypeRateLimited             = ErrorType("rate_limited")
	ErrorTypeTimeout                 = ErrorType("timeout")
	ErrorTypeValidation              = ErrorType("validation")
	ErrorTypeTransaction             = ErrorType("transaction")
	ErrorTypeNotFound                = ErrorType("not_found")
	ErrorTypeBadRequest              = ErrorType("bad_request")
	ErrorTypeInternalServer          = ErrorType("internal_server_error")
	ErrorTypeImmutableSubscription   = ErrorType("immutable_subscription")
	ErrorTypeInvalidApiKey           = ErrorType("invalid_api_key")
	ErrorTypeInvalidContentType      = ErrorType("invalid_content_type")
	ErrorTypeInvalidApiVersion       = ErrorType("invalid_api_version")
	ErrorTypeInvalidPermissions      = ErrorType("invalid_permissions")
	ErrorTypeInvalidToken            = ErrorType("invalid_token")
	ErrorTypeSimulaneousRequest      = ErrorType("simultaneous_request")
	ErrorTypeUnavailableInApiVersion = ErrorType("unavailable_in_api_version")
	ErrorTypeUnknownApiVersion       = ErrorType("unknown_api_version")
	ErrorTypeMissingFeature          = ErrorType("missing_feature")
	ErrorTypeUnauthorized            = ErrorType("unauthorized")
	ErrorTypeForbidden               = ErrorType("forbidden")
	ErrorTypeBadGateway              = ErrorType("bad_gateway")
	ErrorTypeServiceUnavailable      = ErrorType("service_unavailable")

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
	Params           []ErrorParam      `json:"params"`
	TransactionError *TransactionError `json:"transaction_error"`
}

type TransactionError struct {
	TransactionID  string                   `json:"transaction_id"`
	Category       TransactionErrorCategory `json:"category"`
	Code           string                   `json:"code"`
	Message        string                   `json:"message"`
	MerchantAdvice string                   `json:"merchant_advice"`
}

type ErrorParam struct {
	Property string `json:"property"`
	Message  string `json:"message"`
}

// parseResponseToError converts an http.Response to the appropriate error
func parseResponseToError(res *http.Response, body []byte) error {
	// Is this a client error or a server error?
	var errorClass ErrorClass
	if res.StatusCode >= 400 && res.StatusCode < 500 {
		errorClass = ErrorClassClient
	} else {
		errorClass = ErrorClassServer
	}

	// If we have a body, use the details from the body
	if strings.HasPrefix(res.Header.Get("Content-type"), "application/json") {
		// Return an error formatted from the JSON response

		var errResp errorResponse
		if err := json.Unmarshal(body, &errResp); err == nil {
			return &Error{
				Message:          errResp.Error.Message,
				Class:            errorClass,
				Type:             ErrorType(errResp.Error.Type),
				Params:           errResp.Error.Params,
				TransactionError: errResp.Error.TransactionError,
				Response:         parseResponseMetadata(res),
			}
		}
	}

	// If we don't have a body, construct the details from the status code

	errMessage := "An unknown error has occurred. Please try again later."
	errType := ErrorTypeUnknown

	switch res.StatusCode {
	case http.StatusUnauthorized: // 401
		errMessage = "Unauthorized"
		errType = ErrorTypeUnauthorized
	case http.StatusForbidden: // 403
		errMessage = "The API key is not authorized for this resource"
		errType = ErrorTypeForbidden
	case http.StatusNotFound: // 404
		errMessage = "Requested object or endpoint not found"
		errType = ErrorTypeNotFound
	case http.StatusUnprocessableEntity: // 422
		errMessage = "Invalid request"
		errType = ErrorTypeValidation
	case http.StatusTooManyRequests: // 429
		errMessage = "You made too many API requests"
		errType = ErrorTypeRateLimited
	case http.StatusInternalServerError: // 500
		errMessage = "Server experienced an error"
		errType = ErrorTypeInternalServer
	case http.StatusBadGateway: // 502
		errMessage = "Error contacting server"
		errType = ErrorTypeBadGateway
	case http.StatusServiceUnavailable: // 503
		errMessage = "Service unavailable"
		errType = ErrorTypeServiceUnavailable
	case http.StatusRequestTimeout: // 408
	case http.StatusGatewayTimeout: // 504
		errMessage = "Request timed out"
		errType = ErrorTypeTimeout
	}

	return &Error{
		Message:  errMessage,
		Class:    errorClass,
		Type:     errType,
		Response: parseResponseMetadata(res),
	}
}
