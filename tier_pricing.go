// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type TierPricing struct {
	recurlyResponse *ResponseMetadata

	// 3-letter ISO 4217 currency code.
	Currency string `json:"currency,omitempty"`

	// Allows up to 2 decimal places. Required unless `unit_amount_decimal` is provided.
	UnitAmount float64 `json:"unit_amount,omitempty"`

	// Allows up to 9 decimal places. Only supported when `add_on_type` = `usage`.
	// If `unit_amount_decimal` is provided, `unit_amount` cannot be provided.
	UnitAmountDecimal string `json:"unit_amount_decimal,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *TierPricing) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *TierPricing) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type tierPricingList struct {
	ListMetadata
	Data            []TierPricing `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *tierPricingList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *tierPricingList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// TierPricingList allows you to paginate TierPricing objects
type TierPricingList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []TierPricing
}

func NewTierPricingList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *TierPricingList {
	return &TierPricingList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type TierPricingLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []TierPricing
	HasMore() bool
	Next() string
}

func (list *TierPricingList) HasMore() bool {
	return list.hasMore
}

func (list *TierPricingList) Next() string {
	return list.nextPagePath
}

func (list *TierPricingList) Data() []TierPricing {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *TierPricingList) FetchWithContext(ctx context.Context) error {
	resources := &tierPricingList{}
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
func (list *TierPricingList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *TierPricingList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &tierPricingList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *TierPricingList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
