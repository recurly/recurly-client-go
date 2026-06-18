# Changelog

## [v6.0.0](https://github.com/recurly/recurly-client-go/tree/v6.0.0) (2026-06-18)


# Major Version Release

The 6.x major version of the client pairs with the `v2021-02-25` API version. While there are no breaking changes in the API, the client code does include breaking changes.

## Breaking Changes in the Client
- Adds new `params *DeactivateAccountParams` parameter to the `DeactivateAccount` function. Any existing call sites will need to be updated.


