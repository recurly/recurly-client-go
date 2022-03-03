// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type SubscriptionAddOnPercentageTier struct {
	recurlyResponse *ResponseMetadata

	// Ending amount
	EndingAmount float64 `json:"ending_amount,omitempty"`

	// The percentage taken of the monetary amount of usage tracked.
	// This can be up to 4 decimal places represented as a string. A value between
	// 0.0 and 100.0.
	UsagePercentage string `json:"usage_percentage,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *SubscriptionAddOnPercentageTier) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *SubscriptionAddOnPercentageTier) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type subscriptionAddOnPercentageTierList struct {
	ListMetadata
	Data            []SubscriptionAddOnPercentageTier `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *subscriptionAddOnPercentageTierList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *subscriptionAddOnPercentageTierList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// SubscriptionAddOnPercentageTierList allows you to paginate SubscriptionAddOnPercentageTier objects
type SubscriptionAddOnPercentageTierList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []SubscriptionAddOnPercentageTier
}

func NewSubscriptionAddOnPercentageTierList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *SubscriptionAddOnPercentageTierList {
	return &SubscriptionAddOnPercentageTierList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type SubscriptionAddOnPercentageTierLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []SubscriptionAddOnPercentageTier
	HasMore() bool
	Next() string
}

func (list *SubscriptionAddOnPercentageTierList) HasMore() bool {
	return list.hasMore
}

func (list *SubscriptionAddOnPercentageTierList) Next() string {
	return list.nextPagePath
}

func (list *SubscriptionAddOnPercentageTierList) Data() []SubscriptionAddOnPercentageTier {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *SubscriptionAddOnPercentageTierList) FetchWithContext(ctx context.Context) error {
	resources := &subscriptionAddOnPercentageTierList{}
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
func (list *SubscriptionAddOnPercentageTierList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *SubscriptionAddOnPercentageTierList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &subscriptionAddOnPercentageTierList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *SubscriptionAddOnPercentageTierList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
