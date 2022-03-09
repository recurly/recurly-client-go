// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type Tier struct {
	recurlyResponse *ResponseMetadata

	// Ending quantity for the tier.  This represents a unit amount for unit-priced add ons.
	EndingQuantity int `json:"ending_quantity,omitempty"`

	// (deprecated) -- Use the percentage_tiers object instead.
	UsagePercentage string `json:"usage_percentage,omitempty"`

	// Tier pricing
	Currencies []TierPricing `json:"currencies,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *Tier) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *Tier) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type tierList struct {
	ListMetadata
	Data            []Tier `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *tierList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *tierList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// TierList allows you to paginate Tier objects
type TierList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []Tier
}

func NewTierList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *TierList {
	return &TierList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type TierLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []Tier
	HasMore() bool
	Next() string
}

func (list *TierList) HasMore() bool {
	return list.hasMore
}

func (list *TierList) Next() string {
	return list.nextPagePath
}

func (list *TierList) Data() []Tier {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *TierList) FetchWithContext(ctx context.Context) error {
	resources := &tierList{}
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
func (list *TierList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *TierList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &tierList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *TierList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
