package recurly

import (
	"context"
	"net/http"
	"testing"
	"time"
)

func TestCreateAccount(test *testing.T) {
	t := &T{test}
	scenario := &Scenario{
		T: t,
		AssertRequest: func(req *http.Request) {
			t.Assert(req.Method, http.MethodPost, "HTTP Method")
			t.Assert(req.URL.String(), "https://v3.recurly.com/accounts", "Request URL")
			t.Assert(bodyToString(req.Body), `{"code":"new_account"}`, "Request Body")
		},
		MakeResponse: func(req *http.Request) *http.Response {
			return mockResponse(req, 201, String(`{"id": "abcd1234"}`))
		},
	}
	client, err := scenario.MockHTTPClient(nil)
	if err != nil {
	    t.Fatalf("failed to initialize HTTP Client: %v", err)
	}

	body := &AccountCreate{
		Code: String("new_account"),
	}

	account, _ := client.CreateAccount(body)
	t.Assert(account.Id, "abcd1234", "account.Id")
}

func TestCreateAccountWithContext(test *testing.T) {
	t := &T{test}
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	scenario := &Scenario{
		T: t,
		AssertRequest: func(req *http.Request) {
			t.Assert(req.Method, http.MethodPost, "HTTP Method")
			t.Assert(req.Context(), ctx, "Request Context")
			t.Assert(req.URL.String(), "https://v3.recurly.com/accounts", "Request URL")
			t.Assert(bodyToString(req.Body), `{"code":"new_account"}`, "Request Body")
		},
		MakeResponse: func(req *http.Request) *http.Response {
			return mockResponse(req, 201, String(`{"id": "abcd1234"}`))
		},
	}
	client, err := scenario.MockHTTPClient(nil)
	if err != nil {
	    t.Fatalf("failed to initialize HTTP Client: %v", err)
	}

	body := &AccountCreate{
		Code: String("new_account"),
	}

	account, _ := client.CreateAccountWithContext(ctx, body)
	t.Assert(account.Id, "abcd1234", "account.Id")
}

func TestListAccounts(test *testing.T) {
	t := &T{test}

	scenario := &Scenario{
		T: t,
		AssertRequest: func(req *http.Request) {
			t.Assert(req.Method, http.MethodGet, "HTTP Method")
			t.Assert(req.URL.String(), "https://v3.recurly.com/accounts?limit=1&order=desc&sort=created_at", "Request URL")
		},
		MakeResponse: func(req *http.Request) *http.Response {
			return mockResponse(req, 200, String(`{
				"object": "list",
				"has_more": true,
				"next": "/accounts?cursor=efgh5678%3A1588803986.0&limit=1&order=desc&sort=created_at",
				"data": [
					{
						"id": "abcd1234", 
						"first_name": "marigold",
						"last_name": "sunflower"
					},
					{
						"id": "efgh5678", 
						"first_name": "juniper",
						"last_name": "pinecone"
					}
				]
			}`))
		},
	}
	client, err := scenario.MockHTTPClient(nil)
	if err != nil {
	    t.Fatalf("failed to initialize HTTP Client: %v", err)
	}

	params := &ListAccountsParams{
		Sort:  String("created_at"),
		Order: String("desc"),
		Limit: Int(1),
	}
	accounts, err := client.ListAccounts(params)
	t.Assert(err, nil, "Error not expected")
	accounts.Fetch()

	t.Assert(accounts.Data()[0].Id, "abcd1234", "list.Data[0].Id")
	t.Assert(accounts.Data()[1].FirstName, "juniper", "list.Data[1].FirstName")
	t.Assert(accounts.HasMore(), true, "list.HasMore")
	t.Assert(accounts.Next(), "/accounts?cursor=efgh5678%3A1588803986.0&limit=1&order=desc&sort=created_at", "accounts.nextPagePath")
}

func TestListAccountsWithContext(test *testing.T) {
	t := &T{test}
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	scenario := &Scenario{
		T: t,
		AssertRequest: func(req *http.Request) {
			t.Assert(req.Method, http.MethodGet, "HTTP Method")
			t.Assert(req.Context(), ctx, "Request Context")
			t.Assert(req.URL.String(), "https://v3.recurly.com/accounts?limit=1&order=desc&sort=created_at", "Request URL")
		},
		MakeResponse: func(req *http.Request) *http.Response {
			return mockResponse(req, 200, String(`{
				"object": "list",
				"has_more": true,
				"next": "/accounts?cursor=efgh5678%3A1588803986.0&limit=1&order=desc&sort=created_at",
				"data": [
					{
						"id": "abcd1234", 
						"first_name": "marigold",
						"last_name": "sunflower"
					},
					{
						"id": "efgh5678", 
						"first_name": "juniper",
						"last_name": "pinecone"
					}
				]
			}`))
		},
	}
	client, err := scenario.MockHTTPClient(nil)
	if err != nil {
	    t.Fatalf("failed to initialize HTTP Client: %v", err)
	}

	params := &ListAccountsParams{
		Sort:  String("created_at"),
		Order: String("desc"),
		Limit: Int(1),
	}
	accounts, err := client.ListAccounts(params)
	t.Assert(err, nil, "Error not expected")
	accounts.FetchWithContext(ctx)

	t.Assert(accounts.Data()[0].Id, "abcd1234", "list.Data[0].Id")
	t.Assert(accounts.Data()[1].FirstName, "juniper", "list.Data[1].FirstName")
	t.Assert(accounts.HasMore(), true, "list.HasMore")
	t.Assert(accounts.Next(), "/accounts?cursor=efgh5678%3A1588803986.0&limit=1&order=desc&sort=created_at", "accounts.nextPagePath")
}

func TestCreateAccountWithNilContext(test *testing.T) {
	t := &T{test}
	scenario := &Scenario{
		T: t,
		AssertRequest: func(req *http.Request) {
			t.Assert(req.Method, http.MethodPost, "HTTP Method")
			t.Assert(req.Context(), context.Background(), "Request Context")
			t.Assert(req.URL.String(), "https://v3.recurly.com/accounts", "Request URL")
			t.Assert(bodyToString(req.Body), `{"code":"new_account"}`, "Request Body")
		},
		MakeResponse: func(req *http.Request) *http.Response {
			return mockResponse(req, 201, String(`{"id": "abcd1234"}`))
		},
	}
	client, err := scenario.MockHTTPClient(nil)
	if err != nil {
	    t.Fatalf("failed to initialize HTTP Client: %v", err)
	}

	body := &AccountCreate{
		Code: String("new_account"),
	}

	account, _ := client.CreateAccountWithContext(nil, body)
	t.Assert(account.Id, "abcd1234", "account.Id")
}
