package recurly

import (
	"net/http"
	"testing"
)

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
			return mockResponse(req, 200, `{"id": "abcd1234"}`)
		},
	}
	client := scenario.MockHTTPClient()

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
			body := `{ "error": { "type": "not_found", "message": "\"idontexist\" not found", "params": [{ "param": "resource_id", "message": "not found" }] }}`
			return mockResponse(req, 404, body)
		},
	}
	client := scenario.MockHTTPClient()

	resource, err := client.GetResource("idontexist")
	if resource != nil {
		t.Error("Expected Resource to be nil")
	}
	t.Assert(err.Error(), "\"idontexist\" not found", "err.Error()")
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
			return mockResponse(req, 201, body)
		},
	}
	client := scenario.MockHTTPClient()

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
			return mockResponse(req, 422, body)
		},
	}
	client := scenario.MockHTTPClient()

	body := &ResourceCreate{
		String: "hello world",
	}
	resource, err := client.CreateResource(body)
	if resource != nil {
		t.Error("Expected Resource to be nil")
	}
	t.Assert(err.Error(), "Param 'string' is bad", "err.Error()")
}

func TestClientInjectsResponseMetadataIntoResource(test *testing.T) {
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
			return mockResponse(req, 200, `{"id": "abcd1234"}`)
		},
	}
	client := scenario.MockHTTPClient()

	resource, err := client.GetResource("abcd1234")
	t.Assert(err, nil, "Error not expected")

	resp := resource.GetResponse()
	t.Assert(resp.StatusCode, 200, "resp.StatusCode")
	t.Assert(resp.Request.Method, "GET", "resp.Request.Method")
	t.Assert(resp.Request.ID, "msy-1234", "resp.Request.ID")
	t.Assert(resp.RateLimit.Limit, 2000, "resp.RateLimit.Limit")
	t.Assert(resp.RateLimit.Remaining, 1999, "resp.RateLimit.Remaining")
	t.Assert(resp.Version, "recurly."+APIVersion, "resp.Version")
}
