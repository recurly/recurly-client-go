package recurly

import (
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
	client       HttpCaller
	nextPagePath string

	HasMore bool
	Data    []Settings
}

func NewSettingsList(client HttpCaller, nextPagePath string) *SettingsList {
	return &SettingsList{
		client:       client,
		nextPagePath: nextPagePath,
		HasMore:      true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *SettingsList) Fetch() error {
	resources := &settingsList{}
	err := list.client.Call(http.MethodGet, list.nextPagePath, nil, resources)
	if err != nil {
		return err
	}
	// copy over properties from the response
	list.nextPagePath = resources.Next
	list.HasMore = resources.HasMore
	list.Data = resources.Data
	return nil
}

// Count returns the count of items on the server that match this pager
func (list *SettingsList) Count() (*int64, error) {
	resources := &settingsList{}
	err := list.client.Call(http.MethodHead, list.nextPagePath, nil, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}
