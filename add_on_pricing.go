// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type AddOnPricing struct {
	recurlyResponse *ResponseMetadata

	// 3-letter ISO 4217 currency code.
	Currency string `json:"currency,omitempty"`

	// Unit price
	UnitAmount float64 `json:"unit_amount,omitempty"`

	// This field is deprecated. Please do not use it.
	TaxInclusive bool `json:"tax_inclusive,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *AddOnPricing) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *AddOnPricing) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type addOnPricingList struct {
	ListMetadata
	Data            []AddOnPricing `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *addOnPricingList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *addOnPricingList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// AddOnPricingList allows you to paginate AddOnPricing objects
type AddOnPricingList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	HasMore        bool
	Data           []AddOnPricing
}

func NewAddOnPricingList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *AddOnPricingList {
	return &AddOnPricingList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		HasMore:        true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *AddOnPricingList) FetchWithContext(ctx context.Context) error {
	resources := &addOnPricingList{}
	err := list.client.Call(ctx, http.MethodGet, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return err
	}
	// copy over properties from the response
	list.nextPagePath = resources.Next
	list.HasMore = resources.HasMore
	list.Data = resources.Data
	return nil
}

// Fetch fetches the next page of data into the `Data` property
func (list *AddOnPricingList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *AddOnPricingList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &addOnPricingList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *AddOnPricingList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
