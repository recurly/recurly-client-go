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
			return mockResponse(req, 200, String(`{"id": "abcd1234"}`))
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
			return mockResponse(req, 404, String(body))
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
			return mockResponse(req, 201, String(body))
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
			return mockResponse(req, 422, String(body))
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
		T:             t,
		AssertRequest: func(req *http.Request) {},
		MakeResponse: func(req *http.Request) *http.Response {
			return mockResponse(req, 200, String(`{"id": "abcd1234"}`))
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
	t.Assert(resp.IdempotencyPrior, true, "resp.IdempotencyPrior")

	// check that the we inject the metadata into the Empty resource as well
	scenario = &Scenario{
		T:             t,
		AssertRequest: func(req *http.Request) {},
		MakeResponse: func(req *http.Request) *http.Response {
			return mockResponse(req, 204, String(``))
		},
	}
	client = scenario.MockHTTPClient()

	empty, err := client.DeleteResource("abcd1234")

	resp = empty.GetResponse()
	t.Assert(resp.Request.ID, "msy-1234", "resp.Request.ID")
}

func TestHeadResource200(test *testing.T) {
	t := &T{test}

	scenario := &Scenario{
		T:             t,
		AssertRequest: func(req *http.Request) {},
		MakeResponse: func(req *http.Request) *http.Response {
			return mockResponse(req, 200, nil)
		},
	}
	client := scenario.MockHTTPClient()

	resp, err := client.GetResourceMetadata("abcd1234")
	t.Assert(err, nil, "Error not expected")

	t.Assert(resp.StatusCode, 200, "resp.StatusCode")
	t.Assert(resp.Request.Method, "HEAD", "resp.Request.Method")
	t.Assert(resp.Request.ID, "msy-1234", "resp.Request.ID")
	t.Assert(resp.RateLimit.Limit, 2000, "resp.RateLimit.Limit")
	t.Assert(resp.RateLimit.Remaining, 1999, "resp.RateLimit.Remaining")
	t.Assert(resp.Version, "recurly."+APIVersion, "resp.Version")
	intVal := *resp.TotalRecords
	if intVal != 100 {
		t.Errorf("resp.TotalRecords is incorrect. Expected: %v Got: %v", 100, intVal)
	}
}

func TestEncodePathParameters(test *testing.T) {
	t := &T{test}

	scenario := &Scenario{
		T: t,
		AssertRequest: func(req *http.Request) {
			t.Assert(req.Method, http.MethodGet, "HTTP Method")
			t.Assert(req.URL.String(), "https://v3.recurly.com/resources/%2F", "Request URL")
			// assert headers and other request properties
		},
		MakeResponse: func(req *http.Request) *http.Response {
			// default headers set, we may want to customize though
			return mockResponse(req, 200, String(`{"id": "abcd1234"}`))
		},
	}
	client := scenario.MockHTTPClient()

	resource, err := client.GetResource("/")
	t.Assert(err, nil, "Error not expected")
	t.Assert(resource.Id, "abcd1234", "resource.Id")
}

func TestValidatePathParameters(test *testing.T) {
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
			return mockResponse(req, 200, String(`{"id": "abcd1234"}`))
		},
	}
	client := scenario.MockHTTPClient()

	_, err := client.GetResource("")
	t.Assert(err.Error(), "Operation parameters cannot be empty strings.", "err.Error()")
}

func TestSetCustomHeader(test *testing.T) {
	t := &T{test}

	hKey := "Doug"
	hVal := "Miller"

	scenario := &Scenario{
		T: t,
		AssertRequest: func(req *http.Request) {
			t.Assert(req.Header.Get(hKey), hVal, "Set Custom Header")
		},
		MakeResponse: func(req *http.Request) *http.Response {
			// default headers set, we may want to customize though
			return mockResponse(req, 200, String(`{"id": "abcd1234"}`))
		},
	}
	client := scenario.MockHTTPClient()

	header := http.Header{hKey: []string{hVal}}
	client.GetResource("abcd1234", WithHeader(header))
}

func TestPreserveInternalHeaders(test *testing.T) {
	t := &T{test}

	scenario := &Scenario{
		T: t,
		AssertRequest: func(req *http.Request) {
			t.Assert(req.Header.Get("Content-Type"), "application/json; charset=utf-8", "Preserve Internal Headers")
		},
		MakeResponse: func(req *http.Request) *http.Response {
			// default headers set, we may want to customize though
			return mockResponse(req, 200, String(`{"id": "abcd1234"}`))
		},
	}
	client := scenario.MockHTTPClient()

	header := http.Header{"Content-Type": []string{"Custom-Value"}}
	client.GetResource("abcd1234", WithHeader(header))
}

func TestSetIdempotencyKey(test *testing.T) {
	t := &T{test}

	key := "Doug"

	scenario := &Scenario{
		T: t,
		AssertRequest: func(req *http.Request) {
			t.Assert(req.Header.Get("Idempotency-Key"), key, "Set Idempotency Key")
		},
		MakeResponse: func(req *http.Request) *http.Response {
			// default headers set, we may want to customize though
			return mockResponse(req, 200, String(`{"id": "abcd1234"}`))
		},
	}
	client := scenario.MockHTTPClient()

	client.GetResource("abcd1234", WithIdempotencyKey(key))
}

func TestBaseUrlDefaultsToUSDataCenter(test *testing.T) {
	t := &T{test}

	scenario := &Scenario{
		T:             t,
		AssertRequest: nil,
		MakeResponse:  nil,
	}
	client := scenario.MockHTTPClient()

	t.Assert(client.baseURL, "https://v3.recurly.com", "Base URL")
}

func TestBaseUrlInEUDataCenter(test *testing.T) {
	t := &T{test}

	scenario := &Scenario{
		T:             t,
		AssertRequest: nil,
		MakeResponse:  nil,
	}
	client, err := scenario.MockHTTPClientWithOptions(ClientOptions{
		Region: EU,
	})
	if err != nil {
		t.Assert(err, nil, "Error not expected")
	}

	t.Assert(client.baseURL, "https://v3.eu.recurly.com", "Base URL")
}

func TestInvalidRegionError(test *testing.T) {
	t := &T{test}

	scenario := &Scenario{
		T:             t,
		AssertRequest: nil,
		MakeResponse:  nil,
	}
	_, err := scenario.MockHTTPClientWithOptions(ClientOptions{
		Region: region("invalid-region"),
	})
	t.Assert(err.Error(), "invalid region: invalid-region", "err.Error()")

}
