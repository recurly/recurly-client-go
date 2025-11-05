# Changelog

## [v5.7.0](https://github.com/recurly/recurly-client-go/tree/v5.7.0) (2025-11-05)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v5.6.0...v5.7.0)





## [v5.6.0](https://github.com/recurly/recurly-client-go/tree/v5.6.0) (2025-11-03)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v5.5.0...v5.6.0)


**Merged Pull Requests**

- Generated Latest Changes for v2021-02-25 [#250](https://github.com/recurly/recurly-client-go/pull/250) ([recurly-integrations](https://github.com/recurly-integrations))



## [v5.5.0](https://github.com/recurly/recurly-client-go/tree/v5.5.0) (2025-10-09)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v5.4.0...v5.5.0)


**Merged Pull Requests**

- Generated Latest Changes for v2021-02-25 [#247](https://github.com/recurly/recurly-client-go/pull/247) ([recurly-integrations](https://github.com/recurly-integrations))



## [v5.4.0](https://github.com/recurly/recurly-client-go/tree/v5.4.0) (2025-08-13)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v5.3.0...v5.4.0)


**Merged Pull Requests**

- Generated Latest Changes for v2021-02-25 [#246](https://github.com/recurly/recurly-client-go/pull/246) ([recurly-integrations](https://github.com/recurly-integrations))



## [v5.3.0](https://github.com/recurly/recurly-client-go/tree/v5.3.0) (2025-07-22)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v5.2.0...v5.3.0)


**Merged Pull Requests**

- Generated Latest Changes for v2021-02-25 [#245](https://github.com/recurly/recurly-client-go/pull/245) ([recurly-integrations](https://github.com/recurly-integrations))



## [v5.2.0](https://github.com/recurly/recurly-client-go/tree/v5.2.0) (2025-07-09)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v5.1.0...v5.2.0)


**Merged Pull Requests**

- Generated Latest Changes for v2021-02-25 [#243](https://github.com/recurly/recurly-client-go/pull/243) ([recurly-integrations](https://github.com/recurly-integrations))



## [v5.1.0](https://github.com/recurly/recurly-client-go/tree/v5.1.0) (2025-06-11)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v5.0.0...v5.1.0)


**Merged Pull Requests**

- Generated Latest Changes for v2021-02-25 [#242](https://github.com/recurly/recurly-client-go/pull/242) ([recurly-integrations](https://github.com/recurly-integrations))



## [v5.0.0](https://github.com/recurly/recurly-client-go/tree/v5.0.0) (2025-05-16)


# Major Version Release

The 5.x major version of the client pairs with the `v2021-02-25` API version. While there are no breaking changes in the API, the client code does include breaking changes.

## Breaking Changes in the Client
- All request structs have been updated to use pointers to slices instead of the slice itself.
    ```diff
    -       AddOns []SubscriptionAddOnUpdate `json:"add_ons,omitempty"`
    +       AddOns *[]SubscriptionAddOnUpdate `json:"add_ons,omitempty"`
    ```
- All references to `time.Time` have been updated to be pointers to allow for nil values.
- The `UnitAmount` of the `SubscriptionRampInterval` has been updated from `*int` to `*float64` to match the API expectations.


