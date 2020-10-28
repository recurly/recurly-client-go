package recurly

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Error contains basic information about the error
type Error struct {
	recurlyResponse *ResponseMetadata

	Message          string
	Class            ErrorClass
	Type             ErrorType
	Params           []ErrorParam
	TransactionError *TransactionError
}

// GetResponse returns the ResponseMetadata that generated this error
func (resource *Error) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the response metadata
func (resource *Error) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

func (e *Error) Error() string {
	return e.Message
}

type ErrorType string
type ErrorClass string

type TransactionErrorCategory string

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
	TransactionID             string                   `json:"transaction_id"`
	Category                  TransactionErrorCategory `json:"category"`
	Code                      string                   `json:"code"`
	Message                   string                   `json:"message"`
	MerchantAdvice            string                   `json:"merchant_advice"`
	ThreeDSecureActionTokenId string                   `json:"three_d_secure_action_token_id"`
}

type ErrorParam struct {
	Property string `json:"param"`
	Message  string `json:"message"`
}

const (
	ErrorClassServer = ErrorClass("server")
	ErrorClassClient = ErrorClass("client")

	ErrorTypeUnknown = ErrorType("unknown")

	TransactionErrorCategorySoft          = TransactionErrorCategory("soft")
	TransactionErrorCategoryFraud         = TransactionErrorCategory("fraud")
	TransactionErrorCategoryHard          = TransactionErrorCategory("hard")
	TransactionErrorCategoryCommunication = TransactionErrorCategory("communication")
	TransactionErrorCategoryUnknown       = TransactionErrorCategory("unknown")
)

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
				recurlyResponse:  parseResponseMetadata(res),
			}
		}
	}

	// If we don't have a body, construct the details from the status code

	errMessage := fmt.Sprintf("An unexpected %d error has occurred.", res.StatusCode)
	errType := ErrorFromStatusCode(res)

	// This is here to preserve existing behavior from hard coded error
	// handling that existed prior to the introduction of ErrorFromStatusCode
	switch res.StatusCode {
	case http.StatusUnprocessableEntity: // 422
		errType = ErrorTypeValidation
	case http.StatusTooManyRequests: // 429
		errType = ErrorTypeRateLimited
	}

	return &Error{
		Message:         errMessage,
		Class:           errorClass,
		Type:            errType,
		recurlyResponse: parseResponseMetadata(res),
	}
}
