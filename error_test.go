package recurly

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

func bodyToBytes(io io.ReadCloser) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(io)
	return buf.Bytes()
}

func mockResponseWithBody(statusCode int, body string) *http.Response {
	req := &http.Request{Method: "GET"}
	resp := mockResponse(req, statusCode, String(body))
	return resp
}
func mockResponseWithoutBody(statusCode int) *http.Response {
	req := &http.Request{Method: "GET"}
	resp := mockResponse(req, statusCode, String(""))
	resp.Header.Del("Content-Type")
	return resp
}

func parseError(resp *http.Response) error {
	return parseResponseToError(resp, bodyToBytes(resp.Body))
}

func Test401UnauthorizedError(test *testing.T) {
	t := &T{test}

	resp := mockResponseWithBody(401, `
		{
			 "error": {
				"type": "unauthorized",
				"message": "Invalid API key"
			}
		}`)
	err := parseError(resp)
	t.Assert(err.Error(), "Invalid API key", "err.Error()")
	if e, ok := err.(*Error); ok {
		t.Assert(e.Type, ErrorTypeUnauthorized, "e.Type")
		t.Assert(e.Class, ErrorClassClient, "e.Class")
	}

	// test fallback without body
	resp = mockResponseWithoutBody(401)
	err = parseError(resp)
	t.Assert(err.Error(), "Unauthorized", "err.Error()")
	if e, ok := err.(*Error); ok {
		t.Assert(e.Type, ErrorTypeUnauthorized, "e.Type")
		t.Assert(e.Class, ErrorClassClient, "e.Class")
	}
}

func Test403ForbiddenError(test *testing.T) {
	t := &T{test}
	resp := mockResponseWithBody(403, `
		{
			 "error": {
				"type": "forbidden",
				"message": "API key cannot access this resource"
			}
		}`)
	err := parseError(resp)
	t.Assert(err.Error(), "API key cannot access this resource", "err.Error()")
	if e, ok := err.(*Error); ok {
		t.Assert(e.Type, ErrorTypeForbidden, "e.Type")
		t.Assert(e.Class, ErrorClassClient, "e.Class")
	}

	// test fallback without body
	resp = mockResponseWithoutBody(403)
	err = parseError(resp)
	t.Assert(err.Error(), "The API key is not authorized for this resource", "err.Error()")
	if e, ok := err.(*Error); ok {
		t.Assert(e.Type, ErrorTypeForbidden, "e.Type")
		t.Assert(e.Class, ErrorClassClient, "e.Class")
	}
}

func Test404NotFoundError(test *testing.T) {
	t := &T{test}

	resp := mockResponseWithBody(404, `
		{
			 "error": {
				"type": "not_found",
				"message": "It's Not found"
			}
		}`)
	err := parseError(resp)
	t.Assert(err.Error(), "It's Not found", "err.Error()")
	if e, ok := err.(*Error); ok {
		t.Assert(e.Type, ErrorTypeNotFound, "e.Type")
		t.Assert(e.Class, ErrorClassClient, "e.Class")
	}

	// test fallback without body
	resp = mockResponseWithoutBody(404)
	err = parseError(resp)
	t.Assert(err.Error(), "Requested object or endpoint not found", "err.Error()")
	if e, ok := err.(*Error); ok {
		t.Assert(e.Type, ErrorTypeNotFound, "e.Type")
		t.Assert(e.Class, ErrorClassClient, "e.Class")
	}
}
func Test422ValidationError(test *testing.T) {
	t := &T{test}

	resp := mockResponseWithBody(422, `
		{
			 "error": {
				"type": "validation",
				"message": "It's Invalid",
				"params": [
					{"param":"property","message":"is invalid"}
			    ]
			}
		}`)
	err := parseError(resp)
	t.Assert(err.Error(), "It's Invalid", "err.Error()")
	if e, ok := err.(*Error); ok {
		t.Assert(e.Type, ErrorTypeValidation, "e.Type")
		t.Assert(e.Class, ErrorClassClient, "e.Class")
		t.Assert(len(e.Params), 1, "e")
	}

	// test fallback without body
	resp = mockResponseWithoutBody(422)
	err = parseError(resp)
	t.Assert(err.Error(), "Invalid request", "err.Error()")
	if e, ok := err.(*Error); ok {
		t.Assert(e.Type, ErrorTypeValidation, "e.Type")
		t.Assert(e.Class, ErrorClassClient, "e.Class")
	}
}

func Test429RateLimitedError(test *testing.T) {
	t := &T{test}

	resp := mockResponseWithBody(429, `
		{
			 "error": {
				"type": "rate_limited",
				"message": "Too many requests"
			}
		}`)
	err := parseError(resp)
	t.Assert(err.Error(), "Too many requests", "err.Error()")
	if e, ok := err.(*Error); ok {
		t.Assert(e.Type, ErrorTypeRateLimited, "e.Type")
		t.Assert(e.Class, ErrorClassClient, "e.Class")
	}

	// test fallback without body
	resp = mockResponseWithoutBody(429)
	err = parseError(resp)
	t.Assert(err.Error(), "You made too many API requests", "err.Error()")
	if e, ok := err.(*Error); ok {
		t.Assert(e.Type, ErrorTypeRateLimited, "e.Type")
		t.Assert(e.Class, ErrorClassClient, "e.Class")
	}
}

func Test500InternalServerError(test *testing.T) {
	t := &T{test}

	resp := mockResponseWithBody(500, `
		{
			 "error": {
				"type": "internal_server_error",
				"message": "Unknown Server Error"
			}
		}`)
	err := parseError(resp)
	t.Assert(err.Error(), "Unknown Server Error", "err.Error()")
	if e, ok := err.(*Error); ok {
		t.Assert(e.Type, ErrorTypeInternalServer, "e.Type")
		t.Assert(e.Class, ErrorClassServer, "e.Class")
	}

	// test fallback without body
	resp = mockResponseWithoutBody(500)
	err = parseError(resp)
	t.Assert(err.Error(), "Server experienced an error", "err.Error()")
	if e, ok := err.(*Error); ok {
		t.Assert(e.Type, ErrorTypeInternalServer, "e.Type")
		t.Assert(e.Class, ErrorClassServer, "e.Class")
	}
}

func Test502BadGatewatError(test *testing.T) {
	t := &T{test}

	resp := mockResponseWithBody(502, `
		{
			 "error": {
				"type": "bad_gateway",
				"message": "A bad gateway"
			}
		}`)
	err := parseError(resp)
	t.Assert(err.Error(), "A bad gateway", "err.Error()")
	if e, ok := err.(*Error); ok {
		t.Assert(e.Type, ErrorTypeBadGateway, "e.Type")
		t.Assert(e.Class, ErrorClassServer, "e.Class")
	}

	// test fallback without body
	resp = mockResponseWithoutBody(502)
	err = parseError(resp)
	t.Assert(err.Error(), "Error contacting server", "err.Error()")
	if e, ok := err.(*Error); ok {
		t.Assert(e.Type, ErrorTypeBadGateway, "e.Type")
		t.Assert(e.Class, ErrorClassServer, "e.Class")
	}
}

func Test503ServiceUnavailableError(test *testing.T) {
	t := &T{test}

	resp := mockResponseWithBody(503, `
		{
			 "error": {
				"type": "service_unavailable",
				"message": "The service is unavailable"
			}
		}`)
	err := parseError(resp)
	t.Assert(err.Error(), "The service is unavailable", "err.Error()")
	if e, ok := err.(*Error); ok {
		t.Assert(e.Type, ErrorTypeServiceUnavailable, "e.Type")
		t.Assert(e.Class, ErrorClassServer, "e.Class")
	}

	// test fallback without body
	resp = mockResponseWithoutBody(503)
	err = parseError(resp)
	t.Assert(err.Error(), "Service unavailable", "err.Error()")
	if e, ok := err.(*Error); ok {
		t.Assert(e.Type, ErrorTypeServiceUnavailable, "e.Type")
		t.Assert(e.Class, ErrorClassServer, "e.Class")
	}
}

func Test504GatewayTimeoutError(test *testing.T) {
	t := &T{test}

	resp := mockResponseWithBody(504, `
		{
			 "error": {
				"type": "timeout",
				"message": "The request timed out"
			}
		}`)
	err := parseError(resp)
	t.Assert(err.Error(), "The request timed out", "err.Error()")
	if e, ok := err.(*Error); ok {
		t.Assert(e.Type, ErrorTypeTimeout, "e.Type")
		t.Assert(e.Class, ErrorClassServer, "e.Class")
	}

	// test fallback without body
	resp = mockResponseWithoutBody(504)
	err = parseError(resp)
	t.Assert(err.Error(), "Request timed out", "err.Error()")
	if e, ok := err.(*Error); ok {
		t.Assert(e.Type, ErrorTypeTimeout, "e.Type")
		t.Assert(e.Class, ErrorClassServer, "e.Class")
	}
}
