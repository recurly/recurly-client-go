// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type Pricing struct {
	recurlyResponse *ResponseMetadata

	// 3-letter ISO 4217 currency code.
	Currency string `json:"currency,omitempty"`

	// Unit price
	UnitAmount float64 `json:"unit_amount,omitempty"`

	// Determines whether or not tax is included in the unit amount. The Tax Inclusive Pricing feature (separate from the Mixed Tax Pricing feature) must be enabled to use this flag.
	TaxInclusive bool `json:"tax_inclusive,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *Pricing) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *Pricing) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type pricingList struct {
	ListMetadata
	Data            []Pricing `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *pricingList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *pricingList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// PricingList allows you to paginate Pricing objects
type PricingList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []Pricing
}

func NewPricingList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *PricingList {
	return &PricingList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type PricingLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []Pricing
	HasMore() bool
	Next() string
}

func (list *PricingList) HasMore() bool {
	return list.hasMore
}

func (list *PricingList) Next() string {
	return list.nextPagePath
}

func (list *PricingList) Data() []Pricing {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *PricingList) FetchWithContext(ctx context.Context) error {
	resources := &pricingList{}
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
func (list *PricingList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *PricingList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &pricingList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *PricingList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
