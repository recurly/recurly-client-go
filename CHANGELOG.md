# Changelog

## [v4.19.0](https://github.com/recurly/recurly-client-go/tree/v4.19.0) (2022-07-11)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v4.18.0...v4.19.0)


**Merged Pull Requests**

- Generated Latest Changes for v2021-02-25 [#155](https://github.com/recurly/recurly-client-go/pull/155) ([recurly-integrations](https://github.com/recurly-integrations))



## [v4.18.0](https://github.com/recurly/recurly-client-go/tree/v4.18.0) (2022-06-16)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v4.17.0...v4.18.0)


**Merged Pull Requests**

- Generated Latest Changes for v2021-02-25 [#154](https://github.com/recurly/recurly-client-go/pull/154) ([recurly-integrations](https://github.com/recurly-integrations))



## [v4.17.0](https://github.com/recurly/recurly-client-go/tree/v4.17.0) (2022-04-15)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v4.16.0...v4.17.0)


**Merged Pull Requests**

- Generated Latest Changes for v2021-02-25 [#148](https://github.com/recurly/recurly-client-go/pull/148) ([recurly-integrations](https://github.com/recurly-integrations))



## [v4.16.0](https://github.com/recurly/recurly-client-go/tree/v4.16.0) (2022-03-24)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v4.15.0...v4.16.0)


**Merged Pull Requests**

- Generated Latest Changes for v2021-02-25 (Percentage tiers feature) [#145](https://github.com/recurly/recurly-client-go/pull/145) ([recurly-integrations](https://github.com/recurly-integrations))



## [v4.15.0](https://github.com/recurly/recurly-client-go/tree/v4.15.0) (2022-03-14)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v4.14.0...v4.15.0)


**Merged Pull Requests**

- Generated Latest Changes for v2021-02-25 (Percentage tiers feature)  [#140](https://github.com/recurly/recurly-client-go/pull/140) ([recurly-integrations](https://github.com/recurly-integrations))



## [v4.14.0](https://github.com/recurly/recurly-client-go/tree/v4.14.0) (2022-03-03)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v4.13.0...v4.14.0)


**Merged Pull Requests**

- Generated Latest Changes for v2021-02-25 [#139](https://github.com/recurly/recurly-client-go/pull/139) ([recurly-integrations](https://github.com/recurly-integrations))



## [v4.13.0](https://github.com/recurly/recurly-client-go/tree/v4.13.0) (2022-01-31)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v4.12.0...v4.13.0)

A breaking change has been introduced in the 4.13.0 release. 

While [adding support for the EU](https://github.com/recurly/recurly-client-go/pull/132), the `ClientOptions` struct and `NewClientWithOptions` function were added. Part of this process necessitated adding region validations to the `NewClient*` functions.

When creating an instance of a `Client` with the existing `NewClient` function, you will need to handle the potential error response:

Original:
```go
client := recurly.NewClient("<apikey>")
```

Updated:
```go
client, err := recurly.NewClient("<apikey>")
if err != nil {
  // Custom error condition handling
}
```

Using the newly added `NewClientWithOptions`:
```go
client, err := recurly.NewClientWithOptions("<apikey>", recurly.ClientOptions{
    Region: recurly.EU,
})
if err != nil {
  // Custom error condition handling
}
```

While we make every effort to support semantic versioning in our modules, we determined that this breaking change was necessary in the 4.x version of the client.


**Merged Pull Requests**

- Generated Latest Changes for v2021-02-25 [#136](https://github.com/recurly/recurly-client-go/pull/136) ([recurly-integrations](https://github.com/recurly-integrations))
- Add region argument to client to connect in EU data center [#132](https://github.com/recurly/recurly-client-go/pull/132) ([FabricioCoutinho](https://github.com/FabricioCoutinho))



## [v4.12.0](https://github.com/recurly/recurly-client-go/tree/v4.12.0) (2022-01-28)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v4.11.0...v4.12.0)


**Merged Pull Requests**

- Generated Latest Changes for v2021-02-25 (Invoice Customization) [#133](https://github.com/recurly/recurly-client-go/pull/133) ([recurly-integrations](https://github.com/recurly-integrations))
- Generated Latest Changes for v2021-02-25 [#131](https://github.com/recurly/recurly-client-go/pull/131) ([recurly-integrations](https://github.com/recurly-integrations))



## [v4.11.0](https://github.com/recurly/recurly-client-go/tree/v4.11.0) (2021-12-29)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v4.10.0...v4.11.0)


**Merged Pull Requests**

- Generated Latest Changes for v2021-02-25 (Tax Inclusive Pricing) [#130](https://github.com/recurly/recurly-client-go/pull/130) ([recurly-integrations](https://github.com/recurly-integrations))



## [v4.10.0](https://github.com/recurly/recurly-client-go/tree/v4.10.0) (2021-11-22)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v4.9.0...v4.10.0)


**Merged Pull Requests**

- Generated Latest Changes for v2021-02-25 [#128](https://github.com/recurly/recurly-client-go/pull/128) ([recurly-integrations](https://github.com/recurly-integrations))
- Generated Latest Changes for v2021-02-25 [#126](https://github.com/recurly/recurly-client-go/pull/126) ([recurly-integrations](https://github.com/recurly-integrations))
- Generated Latest Changes for v2021-02-25 [#125](https://github.com/recurly/recurly-client-go/pull/125) ([recurly-integrations](https://github.com/recurly-integrations))
- Generated Latest Changes for v2021-02-25 [#122](https://github.com/recurly/recurly-client-go/pull/122) ([recurly-integrations](https://github.com/recurly-integrations))



## [v4.9.0](https://github.com/recurly/recurly-client-go/tree/v4.9.0) (2021-09-16)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v4.8.0...v4.9.0)


**Merged Pull Requests**

- Generated Latest Changes for v2021-02-25 (Support to new subscription fields and response) [#120](https://github.com/recurly/recurly-client-go/pull/120) ([recurly-integrations](https://github.com/recurly-integrations))



## [v4.8.0](https://github.com/recurly/recurly-client-go/tree/v4.8.0) (2021-09-01)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v4.7.0...v4.8.0)


**Merged Pull Requests**

- Generated Latest Changes for v2021-02-25 (Dunning Campaigns feature) [#119](https://github.com/recurly/recurly-client-go/pull/119) ([recurly-integrations](https://github.com/recurly-integrations))



## [v4.7.0](https://github.com/recurly/recurly-client-go/tree/v4.7.0) (2021-08-19)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v4.6.0...v4.7.0)


**Merged Pull Requests**

- Generated Latest Changes for v2021-02-25 (get_preview_renewal) [#117](https://github.com/recurly/recurly-client-go/pull/117) ([recurly-integrations](https://github.com/recurly-integrations))



## [v4.6.0](https://github.com/recurly/recurly-client-go/tree/v4.6.0) (2021-08-11)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v4.5.0...v4.6.0)


**Merged Pull Requests**

- Generated Latest Changes for v2021-02-25 [#115](https://github.com/recurly/recurly-client-go/pull/115) ([recurly-integrations](https://github.com/recurly-integrations))



## [v4.5.0](https://github.com/recurly/recurly-client-go/tree/v4.5.0) (2021-08-02)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v4.4.0...v4.5.0)


**Merged Pull Requests**

- Generated Latest Changes for v2021-02-25 [#110](https://github.com/recurly/recurly-client-go/pull/110) ([recurly-integrations](https://github.com/recurly-integrations))
- feat: add Idempotency-Prior to ResponseMetadata [#106](https://github.com/recurly/recurly-client-go/pull/106) ([speza](https://github.com/speza))



## [v4.4.0](https://github.com/recurly/recurly-client-go/tree/v4.4.0) (2021-06-15)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v4.3.0...v4.4.0)


**Merged Pull Requests**

- Generated Latest Changes for v2021-02-25 [#102](https://github.com/recurly/recurly-client-go/pull/102) ([recurly-integrations](https://github.com/recurly-integrations))



## [v4.3.0](https://github.com/recurly/recurly-client-go/tree/v4.3.0) (2021-06-04)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v4.2.0...v4.3.0)


**Merged Pull Requests**

- Generated Latest Changes for v2021-02-25 [#101](https://github.com/recurly/recurly-client-go/pull/101) ([recurly-integrations](https://github.com/recurly-integrations))



## [v4.2.0](https://github.com/recurly/recurly-client-go/tree/v4.2.0) (2021-04-21)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v4.1.0...v4.2.0)


**Merged Pull Requests**

- Generated Latest Changes for v2021-02-25 [#97](https://github.com/recurly/recurly-client-go/pull/97) ([recurly-integrations](https://github.com/recurly-integrations))



## [v4.1.0](https://github.com/recurly/recurly-client-go/tree/v4.1.0) (2021-04-15)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v4.0.1...v4.1.0)


**Merged Pull Requests**

- Generated Latest Changes for v2021-02-25 (Backup Payment Method) [#96](https://github.com/recurly/recurly-client-go/pull/96) ([recurly-integrations](https://github.com/recurly-integrations))
- Generated Latest Changes for v2021-02-25 [#93](https://github.com/recurly/recurly-client-go/pull/93) ([recurly-integrations](https://github.com/recurly-integrations))
- Generated Latest Changes for v2021-02-25 (Usage Percentage on Tiers) [#92](https://github.com/recurly/recurly-client-go/pull/92) ([recurly-integrations](https://github.com/recurly-integrations))



## [v4.0.1](https://github.com/recurly/recurly-client-go/tree/v4.0.1) (2021-03-23)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v4.0.0...v4.0.1)


**Merged Pull Requests**

- Release 4.0.1 [#91](https://github.com/recurly/recurly-client-go/pull/91) ([douglasmiller](https://github.com/douglasmiller))
- Generated Latest Changes for v2021-02-25 [#90](https://github.com/recurly/recurly-client-go/pull/90) ([recurly-integrations](https://github.com/recurly-integrations))
- Export optionsApplier as OptionsApplier [#81](https://github.com/recurly/recurly-client-go/pull/81) ([jguidry-recurly](https://github.com/jguidry-recurly))

**Closed Issues**

- FetchWithContext not exposed within ClientInterface [#77](https://github.com/recurly/recurly-client-go/issues/77)


## [v4.0.0](https://github.com/recurly/recurly-client-go/tree/v4.0.0) (2021-03-01)


# Major Version Release

The 4.x major version of the client pairs with the `v2021-02-25` API version. This version of the client and the API contain breaking changes that should be considered before upgrading your integration.

## Breaking Changes in the API
All changes to the core API are documented in the [Developer Portal changelog](https://developers.recurly.com/api/changelog.html#v2021-02-25---current-ga-version)

## Breaking Changes in Client

- Empty path parameters are now explicitly invalid and cause an error to be returned.  [#67] [#68]
- `Params` has been removed in favor of `Context`/`RequestOptions`.  [#70]

    ### 3.x
    
    ```go
    headers := http.Header{"Accept-Language": []string{"fr"}}
    
    accountReq := &recurly.AccountCreate{
	    Params: recurly.Params{
		    Header: headers,
	    },
    }
    
    account, err := client.CreateAccount(accountReq)
    ```
    
    ### 4.x
    
    ```go
    headers := http.Header{"Accept-Language": []string{"fr"}}
    
    accountReq := &recurly.AccountCreate{}
    
    account, err := client.CreateAccount(accountReq, recurly.WithHeader(headers))
    ```

- List operations return a pager interface instead of struct.  [#80]

    ### 3.x
    
    ```go
    account, err := client.ListAccounts()

    for accounts.HasMore {
	    err := accounts.Fetch()
	    for i, account := range accounts.Data {
		    fmt.Printf("Account %3d", i)
	    }
    }
    ```
    
    ### 4.x
    
    ```go
    account, err := client.ListAccounts()
    
    for accounts.HasMore() {
	    err := accounts.Fetch()
	    for i, account := range accounts.Data() {
		    fmt.Printf("Account %3d", i)
	    }
    }
    ```


