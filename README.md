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

Errors are configured in [error.go](error.go), including a comprehensive list of error types. Common scenarios in which errors occur may involve "not found" and "validation" errors, which are shown in the examples below.

```go
account, err := client.CreateAccount(accountReq)
if e, ok := err.(*recurly.Error); ok {
  if e.Type == recurly.ErrorTypeValidation {
    // Here we have a validation error
    fmt.Printf("Failed validation: %v", e)
    return nil, err
  }
  // If an error occurs that is not a validation error,
  // we can alert the user to a generic error from the API
  fmt.Printf("Unexpected Recurly error: %v", e)
  return nil, err
}
fmt.Printf("Created Account: %s", account.Id)
```

```go
account, err := client.GetAccount(accountID)
if e, ok := err.(*recurly.Error); ok {
  if e.Type == recurly.ErrorTypeNotFound {
    // Here we have a not-found error
    fmt.Printf("Resource not found: %v", e)
    return nil, err
  }
  fmt.Printf("Unexpected Recurly error: %v", e)
  // If an error occurs that is not a not-found error,
  // we can alert the user to a generic error from the API
  return nil, err
}
fmt.Printf("Fetched Account: %s", account.Id)
```

When calling a method triggers an error, a pointer to the Error struct will be output detailing the error type, complete with an error message. The `e` variable will print out the error message. For example, if `GetAccount("invalid_id")` is called, where `invalid_id` is an invalid account ID, the resulting output is as follows:
```
Body:
{"error":{"type":"not_found","message":"Couldn't find Account with id = invalid_id","params":[{"param":"account_id"}]}}
Resource not found: Couldn't find Account with id = invalid_id
```

## Contributing

Please see our [Contributing Guide](CONTRIBUTING.md).
