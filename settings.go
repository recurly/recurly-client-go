// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type Settings struct {
	recurlyResponse *ResponseMetadata

	// - full:      Full Address (Street, City, State, Postal Code and Country)
	// - streetzip: Street and Postal Code only
	// - zip:       Postal Code only
	// - none:      No Address
	BillingAddressRequirement string `json:"billing_address_requirement,omitempty"`

	AcceptedCurrencies []string `json:"accepted_currencies,omitempty"`

	// The default 3-letter ISO 4217 currency code.
	DefaultCurrency string `json:"default_currency,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *Settings) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *Settings) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type settingsList struct {
	ListMetadata
	Data            []Settings `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *settingsList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *settingsList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// SettingsList allows you to paginate Settings objects
type SettingsList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []Settings
}

type SettingsLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []Settings
	HasMore() bool
	Next() string
}

func NewSettingsList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *SettingsList {
	return &SettingsList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

func (list *SettingsList) HasMore() bool {
	return list.hasMore
}

func (list *SettingsList) Next() string {
	return list.nextPagePath
}

func (list *SettingsList) Data() []Settings {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *SettingsList) FetchWithContext(ctx context.Context) error {
	resources := &settingsList{}
	err := list.client.Call(ctx, http.MethodGet, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return err
	}
	// copy over properties from the response
	list.nextPagePath = resources.Next
	list.hasMore = resources.HasMore
	list.data = resources.Data
	return nil
}

// Fetch fetches the next page of data into the `Data` property
func (list *SettingsList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *SettingsList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &settingsList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *SettingsList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
