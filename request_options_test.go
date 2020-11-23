package recurly

import (
	"context"
	"net/http"
	"testing"
)

func TestWithOption(test *testing.T) {
	type testCase struct {
		name           string
		idempotencyKey string
		context        context.Context
		header         http.Header
	}

	testCases := [...]testCase{
		{
			name:           "WithIdempotencyKey()",
			idempotencyKey: "special-key",
		},
		{
			name:    "WithContext()",
			context: context.Background(),
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
				WithContext(tCase.context),
				WithHeader(tCase.header))
			t.Assert(o.IdempotencyKey, tCase.idempotencyKey, tCase.name)
			t.Assert(o.Context, tCase.context, tCase.name)
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

	ctx := context.Background()
	iKey := "special-key"
	header := http.Header{"Doug": []string{"Miller"}}

	o := &RequestOptions{
		IdempotencyKey: iKey,
		Context:        ctx,
		Header:         header,
	}
	oClone := o.clone()

	// Changing original RequestOptions shouldn't impact the clone
	o.IdempotencyKey = "different-key"
	o.Header = http.Header{}
	o.Context = context.TODO()

	t.Assert(oClone.IdempotencyKey, iKey, "RequestOptions.clone()")
	t.Assert(oClone.Context, ctx, "RequestOptions.clone()")
	for key, v := range header {
		for _, value := range v {
			t.Assert(oClone.Header.Get(key), value, "RequestOptions.clone()")
		}
	}
}

func TestRequestOptionsFromParams(test *testing.T) {
	t := &T{test}

	ctx := context.Background()
	iKey := "special-key"
	header := http.Header{"Doug": []string{"Miller"}}

	p := Params{
		IdempotencyKey: iKey,
		Context:        ctx,
		Header:         header,
	}
	o := RequestOptionsFromParams(p)

	// Changing original RequestOptions shouldn't impact the clone
	p.IdempotencyKey = "different-key"
	p.Header = http.Header{}
	p.Context = context.TODO()

	t.Assert(o.IdempotencyKey, iKey, "RequestOptionsFromParams()")
	t.Assert(o.Context, ctx, "RequestOptionsFromParams()")
	for key, v := range header {
		for _, value := range v {
			t.Assert(o.Header.Get(key), value, "RequestOptionsFromParams()")
		}
	}
}

func TestApplyOptions(test *testing.T) {
	t := &T{test}

	ctx := context.TODO()
	iKey := "special-key"
	header := http.Header{"Doug": []string{"Miller"}}

	o := &RequestOptions{
		IdempotencyKey: iKey,
		Context:        ctx,
		Header:         header,
	}

	req, err := http.NewRequest(http.MethodGet, "/resources", nil)
	if err != nil {
		t.Fail()
	}

	req = o.applyOptions(req)

	t.Assert(req.Context(), ctx, "RequestOptions.applyOptions()")
	for key, v := range header {
		for _, value := range v {
			t.Assert(req.Header.Get(key), value, "RequestOptions.applyOptions()")
		}
	}
	t.Assert(req.Header.Get("Idempotency-Key"), iKey, "RequestOptions.applyOptions()")
}
