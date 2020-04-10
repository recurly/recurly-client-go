# Recurly Golang Client

[APIv3 Documentation](https://developers.recurly.com/api)

## Examples

### Fetching

```go
account, err := client.GetAccount(accountId)
if err != nil {
    fmt.Printf("Failed to retrieve account: %v", err)
} else {
    fmt.Printf("Fetched Account: %s", account.Id)
}
 ```

### Creating

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

### Deleting

```go
account, err := client.DeactivateAccount(accountId)
if err != nil {
    fmt.Printf("Failed to retrieve account: %v", err)
} else {
    fmt.Printf("Deactivated Account: %s", account.Id)
}
```

### Pagination

```go
listParams := &recurly.ListAccountsParams{
    Sort:  recurly.String("created_at"),
    Order: recurly.String("desc"),
    Limit: recurly.Int(200),
}

accounts := client.ListAccounts(listParams)

var err error = nil
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

