package recurly

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

// A wrapper for testing.T
type T struct {
	*testing.T
}

// A helpful method to assert variables are the expected values
func (t *T) Assert(val interface{}, expected interface{}, label string) {
	if val != expected {
		t.Errorf("%v is incorrect. Expected: %v Got: %v", label, expected, val)
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
func (s *Scenario) MockHTTPClient() *Client {
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

	client := newClient("APIKEY", &http.Client{
		Transport: roundTripFunc(roundTrip),
	})

	// override the loger to keep noise down
	client.Log = NewLogger(LevelWarn)

	return client
}

func bodyReader(body *string) io.ReadCloser {
	if body != nil {
		return ioutil.NopCloser(bytes.NewBufferString(*body))
	}
	return ioutil.NopCloser(bytes.NewBuffer(nil))
}

// A helpful utility method for creating default
// http.Responses with Recurly metadata
func mockResponse(req *http.Request, statusCode int, body *string) *http.Response {
	headers := make(http.Header)
	headers.Add("Content-Type", "application/json; charset=utf-8")
	headers.Add("Recurly-Version", "recurly."+APIVersion)
	headers.Add("X-RateLimit-Limit", "2000")
	headers.Add("X-RateLimit-Remaining", "1999")
	headers.Add("X-RateLimit-Reset", "1586203320")
	headers.Add("X-Request-Id", "msy-1234")
	headers.Add("Recurly-Total-Records", "100")
	return &http.Response{
		StatusCode: statusCode,
		Body:       bodyReader(body),
		Header:     headers,
		Request:    req,
	}
}

// turns a body (as an io reader) into a string
func bodyToString(io io.ReadCloser) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(io)
	return buf.String()
}

// RecurlyResource and SubResource are used as placeholders for
// Recurly's resources. This accomplishes 2 goals:
//   1. Keep the test code from coupling to generated code
//	 2. Easily setup specific type scenarios which may or may not exist in Recurly now
type RecurlyResource struct {
	recurlyResponse *ResponseMetadata

	Id          string        `json:"id,omitempty"`
	Integer     int           `json:"int,omitempty"`
	Float       float64       `json:"float,omitempty"`
	DateTime    time.Time     `json:"date_time,omitempty"`
	SubResource SubResource   `json:"sub_resource,omitempty"`
	StrArray    []string      `json:"str_array,omitempty"`
	SubArray    []SubResource `json:"sub_array,omitempty"`
}

func (resource *RecurlyResource) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

func (resource *RecurlyResource) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

type SubResource struct {
	recurlyResponse *ResponseMetadata

	Id string `json:"id,omitempty"`
}

func (resource *SubResource) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

func (resource *SubResource) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
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
func (c *Client) GetResource(resourceId string) (*RecurlyResource, error) {
	path, err := c.InterpolatePath("/resources/{resource_id}", resourceId)
	if err != nil {
		return nil, err
	}
	result := &RecurlyResource{}
	err = c.Call(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (c *Client) CreateResource(body *ResourceCreate) (*RecurlyResource, error) {
	path, err := c.InterpolatePath("/resources")
	if err != nil {
		return nil, err
	}
	result := &RecurlyResource{}
	err = c.Call(http.MethodPost, path, body, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (c *Client) DeleteResource(resourceId string) (*Empty, error) {
	path, err := c.InterpolatePath("/resources")
	if err != nil {
		return nil, err
	}
	result := &Empty{}
	err = c.Call(http.MethodDelete, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// ResourceMetada calls the HEAD endpoint and returns the response metadata
func (c *Client) GetResourceMetadata(resourceId string) (*ResponseMetadata, error) {
	path := c.InterpolatePath("/resources")
	result := &Empty{}
	err := c.Call(http.MethodHead, path, nil, result)
	if err != nil {
		return nil, err
	}
	return result.GetResponse(), err
}

func TestPointerHelpers(test *testing.T) {
	t := &T{test}

	str := "Hello"
	strPtr := String("Hello")
	// verify that recurly.String(s) returns a pointer to s
	t.Assert(*strPtr, str, "recurly.String()")

	i := 100
	intPtr := Int(100)
	// verify that recurly.Int(s) returns a pointer to s
	t.Assert(*intPtr, i, "recurly.Int()")

	flt := 100.0
	fltPtr := Float(100.0)
	// verify that recurly.Float(s) returns a pointer to s
	t.Assert(*fltPtr, flt, "recurly.Float()")

	boo := true
	booPtr := Bool(true)
	// verify that recurly.Bool(s) returns a pointer to s
	t.Assert(*booPtr, boo, "recurly.Bool()")

	z := time.Now()
	zPtr := Time(z)
	// verify that recurly.Time(s) returns a pointer to s
	t.Assert(*zPtr, z, "recurly.Time()")

	strSlice := []string{"apple", "orange", "banana"}
	slicePtr := StringSlice(strSlice)
	// verify that recurly.StringSlice(s) returns a pointer to elements in slice s
	t.Assert(*slicePtr[0], strSlice[0], "recurly.StringSlice()")
}
