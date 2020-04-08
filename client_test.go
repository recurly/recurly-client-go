package recurly

import (
	"bytes"
	"io"
	"io/ioutil"
	"testing"
	"time"

	"net/http"
)

// A wrapper for testing.T
type T struct {
	*testing.T
}

// A helpful method to assert variables are the expected values
func (t *T) Assert(val interface{}, expected interface{}, label string) {
	if val != expected {
		t.Errorf("%v is incorrect. Expected: %v Got: %v", label, val, expected)
	}
}

// roundTripFunc is a function used to mock the transport barrier of http.Client
type roundTripFunc func(req *http.Request) *http.Response

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// A testing "Scenario" is composed of 2 functions:
//	AssertRequest: asserts the correct request properties are being sent to the transport layer
//	MakeResponse: returns a canned response to test the response handling code
type Scenario struct {
	T             *T
	AssertRequest func(req *http.Request)
	MakeResponse  func(req *http.Request) *http.Response
}

// This method on Scenario gives you a *recurly.Client
// which implements the testing scenario
func (s *Scenario) MockHttpClient() *Client {
	roundTrip := func(req *http.Request) *http.Response {
		// Check the request has the expected properties
		s.AssertRequest(req)

		// Assert the default headers on every request
		expectedAccept := "application/vnd.recurly." + APIVersion
		header := req.Header
		s.T.Assert(header.Get("Accept"), expectedAccept, "Request Header \"Accept\"")
		s.T.Assert(header.Get("Accept-Encoding"), "gzip", "Request Header \"Accept-Encoding\"")
		s.T.Assert(header.Get("Content-Type"), "application/json; charset=utf-8", "Request Header \"Content-Type\"")

		// Return the canned Response
		return s.MakeResponse(req)
	}

	return NewClient("APIKEY", &http.Client{
		Transport: roundTripFunc(roundTrip),
	})
}

// A helpful utility method for creating default
// http.Responses with Recurly metadata
func mockResponse(statusCode int, body string) *http.Response {
	headers := make(http.Header)
	headers.Add("Content-Type", "application/json; charset=utf-8")
	headers.Add("Recurly-Version", "recurly"+APIVersion)
	headers.Add("X-RateLimit-Limit", "2000")
	headers.Add("X-RateLimit-Remaining", "1999")
	headers.Add("X-RateLimit-Reset", "1586203320")
	headers.Add("X-RateLimit-Reset", "1586203320")
	headers.Add("X-Request-Id", "msy-1234")
	return &http.Response{
		StatusCode: statusCode,
		Body:       ioutil.NopCloser(bytes.NewBufferString(body)),
		Header:     headers,
	}
}

// turns a body (as an io reader) into a string
func bodyToString(io io.ReadCloser) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(io)
	return buf.String()
}

// Resource and SubResource are used as placeholders for
// Recurly's resources. This accomplishes 2 goals:
//   1. Keep the test code from coupling to generated code
//	 2. Easily setup specific type scenarios which may or may not exist in Recurly now
type Resource struct {
	Id          string        `json:"id,omitempty"`
	Integer     int           `json:"int,omitempty"`
	Float       float64       `json:"float,omitempty"`
	DateTime    time.Time     `json:"date_time,omitempty"`
	SubResource SubResource   `json:"sub_resource,omitempty"`
	StrArray    []string      `json:"str_array,omitempty"`
	SubArray    []SubResource `json:"sub_array,omitempty"`
}

type SubResource struct {
	Id string `json:"id,omitempty"`
}

type ResourceCreate struct {
	Params `json:"-"`
	String string `json:"string,omitempty"`
}

func (attr *ResourceCreate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}

// We also implement fake CRUD operations for these fake resources
// We want to use the Client from the consuming code's perspective
func (c *Client) GetResource(resourceId string) (*Resource, error) {
	path := c.InterpolatePath("/resources/{resource_id}", resourceId)
	result := &Resource{}
	err := c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (c *Client) CreateResource(body *ResourceCreate) (*Resource, error) {
	path := c.InterpolatePath("/resources")
	result := &Resource{}
	err := c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

func TestGetResource200(test *testing.T) {
	t := &T{test}

	scenario := &Scenario{
		T: t,
		AssertRequest: func(req *http.Request) {
			t.Assert(req.Method, http.MethodGet, "HTTP Method")
			t.Assert(req.URL.String(), "https://v3.recurly.com/resources/abcd1234", "Request URL")
			// assert headers and other request properties
		},
		MakeResponse: func(req *http.Request) *http.Response {
			// default headers set, we may want to customize though
			return mockResponse(200, `{"id": "abcd1234"}`)
		},
	}
	client := scenario.MockHttpClient()

	resource, err := client.GetResource("abcd1234")
	t.Assert(err, nil, "Error not expected")
	t.Assert(resource.Id, "abcd1234", "resource.Id")
}

func TestGetResource404(test *testing.T) {
	t := &T{test}
	scenario := &Scenario{
		T: t,
		AssertRequest: func(req *http.Request) {
			t.Assert(req.Method, http.MethodGet, "HTTP Method")
			t.Assert(req.URL.String(), "https://v3.recurly.com/resources/idontexist", "Request URL")
		},
		MakeResponse: func(req *http.Request) *http.Response {
			body := `{ "error": { "type": "not_found", "params": [{ "param": "resource_id" }] }}`
			return mockResponse(404, body)
		},
	}
	client := scenario.MockHttpClient()

	resource, err := client.GetResource("idontexist")
	if resource != nil {
		t.Error("Expected Resource to be nil")
	}
	t.Assert(err.Error(), "requested object or endpoint not found", "err.Error()")
}

func TestCreateResource201(test *testing.T) {
	t := &T{test}
	scenario := &Scenario{
		T: t,
		AssertRequest: func(req *http.Request) {
			t.Assert(req.Method, http.MethodPost, "HTTP Method")
			t.Assert(req.URL.String(), "https://v3.recurly.com/resources", "Request URL")
			t.Assert(bodyToString(req.Body), `{"string":"hello world"}`, "Request Body")
		},
		MakeResponse: func(req *http.Request) *http.Response {
			body := `{ "id": "abcd1234" }`
			return mockResponse(201, body)
		},
	}
	client := scenario.MockHttpClient()

	body := &ResourceCreate{
		String: "hello world",
	}
	resource, err := client.CreateResource(body)
	t.Assert(err, nil, "Error not expected")
	t.Assert(resource.Id, "abcd1234", "resource.Id")
}

func TestCreateResource422(test *testing.T) {
	t := &T{test}
	scenario := &Scenario{
		T: t,
		AssertRequest: func(req *http.Request) {
			t.Assert(req.Method, http.MethodPost, "HTTP Method")
			t.Assert(req.URL.String(), "https://v3.recurly.com/resources", "Request URL")
			t.Assert(bodyToString(req.Body), `{"string":"hello world"}`, "Request Body")
		},
		MakeResponse: func(req *http.Request) *http.Response {
			body := `
				{
					"error": {
						"type": "validation",
						"message": "Param 'string' is bad",
						"params":[{"param":"string","message":"is bad"}]
					}
				}`
			return mockResponse(422, body)
		},
	}
	client := scenario.MockHttpClient()

	body := &ResourceCreate{
		String: "hello world",
	}
	resource, err := client.CreateResource(body)
	if resource != nil {
		t.Error("Expected Resource to be nil")
	}
	t.Assert(err.Error(), "Invalid request", "err.Error()")
}
