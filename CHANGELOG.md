# Changelog

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


