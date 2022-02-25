// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type PercentageTiersByCurrency struct {
	recurlyResponse *ResponseMetadata

	// 3-letter ISO 4217 currency code.
	Currency string `json:"currency,omitempty"`

	// Tiers
	Tiers []PercentageTier `json:"tiers,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *PercentageTiersByCurrency) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *PercentageTiersByCurrency) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type percentageTiersByCurrencyList struct {
	ListMetadata
	Data            []PercentageTiersByCurrency `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *percentageTiersByCurrencyList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *percentageTiersByCurrencyList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// PercentageTiersByCurrencyList allows you to paginate PercentageTiersByCurrency objects
type PercentageTiersByCurrencyList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []PercentageTiersByCurrency
}

func NewPercentageTiersByCurrencyList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *PercentageTiersByCurrencyList {
	return &PercentageTiersByCurrencyList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type PercentageTiersByCurrencyLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []PercentageTiersByCurrency
	HasMore() bool
	Next() string
}

func (list *PercentageTiersByCurrencyList) HasMore() bool {
	return list.hasMore
}

func (list *PercentageTiersByCurrencyList) Next() string {
	return list.nextPagePath
}

func (list *PercentageTiersByCurrencyList) Data() []PercentageTiersByCurrency {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *PercentageTiersByCurrencyList) FetchWithContext(ctx context.Context) error {
	resources := &percentageTiersByCurrencyList{}
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
func (list *PercentageTiersByCurrencyList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *PercentageTiersByCurrencyList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &percentageTiersByCurrencyList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *PercentageTiersByCurrencyList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
