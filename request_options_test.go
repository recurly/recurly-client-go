package recurly

import (
	"net/http"
	"testing"
)

func TestWithOption(test *testing.T) {
	type testCase struct {
		name           string
		idempotencyKey string
		header         http.Header
	}

	testCases := [...]testCase{
		{
			name:           "WithIdempotencyKey()",
			idempotencyKey: "special-key",
		},
		{
			name:   "WithHeader()",
			header: http.Header{"Doug": []string{"Miller"}},
		},
	}

	for _, tCase := range testCases {
		test.Run(tCase.name, func(test *testing.T) {
			t := &T{test}
			o := NewRequestOptions(
				WithIdempotencyKey(tCase.idempotencyKey),
				WithHeader(tCase.header))
			t.Assert(o.IdempotencyKey, tCase.idempotencyKey, tCase.name)
			for key, v := range tCase.header {
				for _, value := range v {
					t.Assert(o.Header.Get(key), value, tCase.name)
				}
			}
		})
	}
}

func TestRequestOptionsClone(test *testing.T) {
	t := &T{test}

	iKey := "special-key"
	header := http.Header{"Doug": []string{"Miller"}}

	o := &RequestOptions{
		IdempotencyKey: iKey,
		Header:         header,
	}
	oClone := o.clone()

	// Changing original RequestOptions shouldn't impact the clone
	o.IdempotencyKey = "different-key"
	o.Header = http.Header{}

	t.Assert(oClone.IdempotencyKey, iKey, "RequestOptions.clone()")
	for key, v := range header {
		for _, value := range v {
			t.Assert(oClone.Header.Get(key), value, "RequestOptions.clone()")
		}
	}
}

func TestRequestOptionsFromParams(test *testing.T) {
	t := &T{test}

	iKey := "special-key"
	header := http.Header{"Doug": []string{"Miller"}}

	p := Params{
		IdempotencyKey: iKey,
		Header:         header,
	}
	o := RequestOptionsFromParams(p)

	// Changing original RequestOptions shouldn't impact the clone
	p.IdempotencyKey = "different-key"
	p.Header = http.Header{}

	t.Assert(o.IdempotencyKey, iKey, "RequestOptionsFromParams()")
	for key, v := range header {
		for _, value := range v {
			t.Assert(o.Header.Get(key), value, "RequestOptionsFromParams()")
		}
	}
}

func TestApplyOptions(test *testing.T) {
	t := &T{test}

	iKey := "special-key"
	header := http.Header{"Doug": []string{"Miller"}}

	o := &RequestOptions{
		IdempotencyKey: iKey,
		Header:         header,
	}

	req, err := http.NewRequest(http.MethodGet, "/resources", nil)
	if err != nil {
		t.Fail()
	}

	req = o.applyOptions(req)

	for key, v := range header {
		for _, value := range v {
			t.Assert(req.Header.Get(key), value, "RequestOptions.applyOptions()")
		}
	}
	t.Assert(req.Header.Get("Idempotency-Key"), iKey, "RequestOptions.applyOptions()")
}
