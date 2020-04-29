# Recurly 

This repository houses the official Go library for Recurly's V3 API.

Documentation for the HTTP API and example code can be found [on our Developer Hub](https://developers.recurly.com/api/latest/index.html).

## Getting Started

### Installing

TODO

### Creating a client

A client represents a connection to the Recurly API. Every call to the server exists as a method on this struct. To initialize, you only need the private API key which can be obtained on the [API Credentials Page](https://app.recurly.com/go/integrations/api_keys).

```go
recurly.APIKey = "<apikey>"
client := recurly.DefaultClient()
```

### Operations

Every operation that can be performed against the API has a corresponding method in the Client struct. The [client_operations.go](client_operations.go) file implements these operations. This file also provides descriptions for each method and their returns. For example, to use the [get_account](https://developers.recurly.com/api/v2019-10-10/index.html#operation/get_account) endpoint, the `GetAccount` method can be called, as seen below. `GetAccount()` is used to fetch an account; it returns a pointer to the Account struct.

```go
account, err := client.GetAccount(accountID)
if e, ok := err.(*recurly.Error); ok {
  if e.Type == recurly.ErrorTypeNotFound {
    fmt.Printf("Resource not found: %v", e)
    return nil, err
  }
  fmt.Printf("Unexpected Recurly error: %v", e)
  return nil, err
}
fmt.Printf("Fetched Account: %s", account.Id)
 ```

### Creating Resources

Every Create method on the client takes a specific Request type to form the request.

```go
accountReq := &recurly.AccountCreate{
    Code:      accountCode,
    FirstName: "Isaac",
    LastName:  "Du Monde",
    Email:     "isaac@example.com",
    BillingInfo: recurly.BillingInfoCreate{
      FirstName: "Isaac",
      LastName:  "Du Monde",
      Address: recurly.Address{
        Phone:      "415-555-5555",
        Street1:    "400 Alabama St.",
        City:       "San Francisco",
        PostalCode: "94110",
        Country:    "US",
        Region:     "CA",
      },
      Number: "4111111111111111",
      Month:  "12",
      Year:   "22",
      Cvv:    "123",
    },
  }

account, err := client.CreateAccount(accountReq)
if e, ok := err.(*recurly.Error); ok {
  if e.Type == recurly.ErrorTypeValidation {
    fmt.Printf("Failed validation: %v", e)
    return nil, err
  }
  fmt.Printf("Unexpected Recurly error: %v", e)
  return nil, err
}
fmt.Printf("Created Account: %s", account.Id)
```

### Pagination

Pagination is accomplished via the `*List` types defined in [resources.go]. For example, `AccountList` allows you to paginate `Account` objects. All `List*` methods on the `Client` return pagers (pointers to these list types).

These list operations accept parameters for sorting and filtering the results (e.g. `Ids`, `Limit`, `Sort`, `Order`). Each list operation accepts its own parameters type. For example, `ListAccounts()` accepts a `ListAccountsParams` object. These types can be found in [Client operations](client_operations.go).

```go
listParams := &recurly.ListAccountsParams{
    Sort:  recurly.String("created_at"),
    Order: recurly.String("desc"),
    Limit: recurly.Int(200),
}

accounts := client.ListAccounts(listParams)
```

`NextPage()` returns the next page of resources.

```go
var err error
for accounts.HasMore {
    accounts, err = accounts.NextPage()
    if e, ok := err.(*recurly.Error); ok {
        fmt.Printf("Failed to retrieve next page: %v", err)
        break
    }
    for i, account := range accounts.Data {
        fmt.Printf("Account %3d: %s, %s\n",
            i,
            account.Id,
            account.Code,
        )
    }
}
```

### Error Handling

Errors are configured in [error.go](error.go). Common scenarios in which errors occur may involve "not found" or "validation" errors, which are included among the examples below. A complete list of error types can be found in [error.go](error.go). You can use the list to customize your case statements.

```go
account, err := client.CreateAccount(accountReq)
if e, ok := err.(*recurly.Error); ok {
  switch e.Type {
  case recurly.ErrorTypeValidation:
    fmt.Printf("Failed validation: %v", e)
    return nil, err
  case recurly.ErrorTypeNotFound:
    fmt.Printf("Resource not found: %v", e)
    return nil, err
  case recurly.ErrorTypeTransaction:
    fmt.Printf("Transaction Error: %v", e)
    return nil, err
  case recurly.ErrorTypeInvalidToken:
    fmt.Printf("Invalid Token: %v", e)
    return nil, err
  case recurly.ErrorTypeTimeout:
    fmt.Printf("Timeout: %v", e)
    return nil, err
  // If an error occurs that is not covered by a case statement,
  // we can alert the user to a generic error from the API
  default:
    fmt.Printf("Unexpected Recurly error: %v", e)
    return nil, err
  }
}
fmt.Printf("Created Account: %s", account.Id)
```

When calling a method triggers an error, a pointer to the Error struct will be output detailing the error type, complete with an error message. The `e` variable will print out the error message. For example, if `GetAccount("invalid_id")` is called, where `invalid_id` is an invalid account ID, the resulting output is as follows:
```
Body:
{"error":{"type":"not_found","message":"Couldn't find Account with id = invalid_id","params":[{"param":"account_id"}]}}
Resource not found: Couldn't find Account with id = invalid_id
```

### HTTP Metadata

Sometimes you might want additional information about the underlying HTTP request and response. Instead of returning this information directly and forcing the programmer to handle it, we inject this metadata into the top level resource that was returned. You can access the response by calling `GetResponse()` on anything that implements the `Resource` interface. This includes the resource objects that are returned from operations, as well as Recurly errors.

> *Warning*: Be careful logging or rendering ResponseMetadata objects as they may contain PII or sensitive data.

#### On a Resource

```go
account, err := client.GetAccount(accountID)
if err != nil {
  return nil, err
}

// GetResponse() returns a *ResponseMetadata object
resp := account.GetResponse()
fmt.Println(resp.Request.ID)          // "58ac04b9d9218319-ATL"
fmt.Println(resp.StatusCode)          // 201
fmt.Println(resp.Request.Method)      // "POST"
fmt.Println(resp.Request.URL)         // "https://v3.recurly.com/accounts"
fmt.Println(resp.RateLimit.Limit)     // 2000
fmt.Println(resp.RateLimit.Remaining) // 1999
```

#### On an Error

It can be helpful to inspect the metadata on errors for logging purposes. Having the Request-Id on hand
will help our support staff diagnose potential problems quickly.

```go
account, err := client.CreateAccount(accountReq)
if e, ok := err.(*recurly.Error); ok {
  resp := e.GetResponse()
  fmt.Println("Unexpected Error. Request Id: {}", resp.Request.ID)
  return nil, err
}
```

#### On any Resource

Both resources and errors implement the `Resource` interface. This allows you to extract
metadata from either in a generic way:

```go
func LogRateLimit(resource *recurly.Resource) {
  fmt.Println("Requests Remaining: {}", resource.GetResponse().RateLimit.Remaining)
}
```

```go
account, err := client.CreateAccount(accountReq)
if e, ok := err.(*recurly.Error); ok {
  LogRateLimit(e)
  return nil, err
}

LogRateLimit(account)
```

## Contributing

Please see our [Contributing Guide](CONTRIBUTING.md).
