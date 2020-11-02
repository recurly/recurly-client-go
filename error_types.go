// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"net/http"
)

const (
	ErrorTypeResponse                = ErrorType("response_error")
	ErrorTypeServer                  = ErrorType("server_error")
	ErrorTypeInternalServer          = ErrorType("internal_server_error")
	ErrorTypeServiceNotAvailable     = ErrorType("service_not_available")
	ErrorTypeBadGateway              = ErrorType("bad_gateway")
	ErrorTypeServiceUnavailable      = ErrorType("service_unavailable")
	ErrorTypeTimeout                 = ErrorType("timeout")
	ErrorTypeRedirection             = ErrorType("redirection")
	ErrorTypeNotModified             = ErrorType("not_modified")
	ErrorTypeClient                  = ErrorType("client_error")
	ErrorTypeBadRequest              = ErrorType("bad_request")
	ErrorTypeInvalidContentType      = ErrorType("invalid_content_type")
	ErrorTypeUnauthorized            = ErrorType("unauthorized")
	ErrorTypePaymentRequired         = ErrorType("payment_required")
	ErrorTypeForbidden               = ErrorType("forbidden")
	ErrorTypeInvalidApiKey           = ErrorType("invalid_api_key")
	ErrorTypeInvalidPermissions      = ErrorType("invalid_permissions")
	ErrorTypeNotFound                = ErrorType("not_found")
	ErrorTypeNotAcceptable           = ErrorType("not_acceptable")
	ErrorTypeUnknownApiVersion       = ErrorType("unknown_api_version")
	ErrorTypeUnavailableInApiVersion = ErrorType("unavailable_in_api_version")
	ErrorTypeInvalidApiVersion       = ErrorType("invalid_api_version")
	ErrorTypePreconditionFailed      = ErrorType("precondition_failed")
	ErrorTypeUnprocessableEntity     = ErrorType("unprocessable_entity")
	ErrorTypeValidation              = ErrorType("validation")
	ErrorTypeMissingFeature          = ErrorType("missing_feature")
	ErrorTypeTransaction             = ErrorType("transaction")
	ErrorTypeSimultaneousRequest     = ErrorType("simultaneous_request")
	ErrorTypeImmutableSubscription   = ErrorType("immutable_subscription")
	ErrorTypeInvalidToken            = ErrorType("invalid_token")
	ErrorTypeTooManyRequests         = ErrorType("too_many_requests")
	ErrorTypeRateLimited             = ErrorType("rate_limited")
)

func ErrorFromStatusCode(res *http.Response) ErrorType {
	var errType ErrorType

	switch res.StatusCode {
	case 500:
		errType = ErrorTypeInternalServer
	case 502:
		errType = ErrorTypeBadGateway
	case 503:
		errType = ErrorTypeServiceUnavailable
	case 504:
		errType = ErrorTypeTimeout
	case 304:
		errType = ErrorTypeNotModified
	case 400:
		errType = ErrorTypeBadRequest
	case 401:
		errType = ErrorTypeUnauthorized
	case 402:
		errType = ErrorTypePaymentRequired
	case 403:
		errType = ErrorTypeForbidden
	case 404:
		errType = ErrorTypeNotFound
	case 406:
		errType = ErrorTypeNotAcceptable
	case 412:
		errType = ErrorTypePreconditionFailed
	case 422:
		errType = ErrorTypeUnprocessableEntity
	case 429:
		errType = ErrorTypeTooManyRequests
	default:
		errType = ErrorTypeUnknown
	}

	return errType
}
