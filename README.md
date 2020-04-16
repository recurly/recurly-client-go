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
account, err := client.GetAccount(accountId)
if err != nil {
    fmt.Printf("Failed to retrieve account: %v", err)
} else {
    fmt.Printf("Fetched Account: %s", account.Id)
}
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
if err != nil {
    fmt.Printf("Failed to create an account: %v", err)
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
    if err != nil {
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

Errors are configured in [error.go](error.go). 

```go
account, err := client.GetAccount(accountID)
if err != nil {
  fmt.Printf("Failed to retrieve account: %v", err)
  os.Exit(1)
} else {
  fmt.Printf("Fetched Account: %s", accountID)
}
```

When calling a method triggers an error, a pointer to the Error struct will be output detailing the error type, complete with an error message. The `err` variable will print out the error message. For example, if `GetAccount()` is called with an invalid `accountID`, the resulting output is as follows:
```
Body:
{"error":{"type":"not_found","message":"Couldn't find Account with id = invalid_id","params":[{"param":"account_id"}]}}
Failed to retrieve account: Couldn't find Account with id = invalid_id
```

## Contributing

Please see our [Contributing Guide](CONTRIBUTING.md).
