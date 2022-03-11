// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type SubscriptionRampIntervalResponse struct {
	recurlyResponse *ResponseMetadata

	// Represents how many billing cycles are included in a ramp interval.
	StartingBillingCycle int `json:"starting_billing_cycle,omitempty"`

	// Represents how many billing cycles are left in a ramp interval.
	RemainingBillingCycles int `json:"remaining_billing_cycles,omitempty"`

	// Represents the price for the ramp interval.
	UnitAmount int `json:"unit_amount,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *SubscriptionRampIntervalResponse) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *SubscriptionRampIntervalResponse) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type subscriptionRampIntervalResponseList struct {
	ListMetadata
	Data            []SubscriptionRampIntervalResponse `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *subscriptionRampIntervalResponseList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *subscriptionRampIntervalResponseList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// SubscriptionRampIntervalResponseList allows you to paginate SubscriptionRampIntervalResponse objects
type SubscriptionRampIntervalResponseList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []SubscriptionRampIntervalResponse
}

func NewSubscriptionRampIntervalResponseList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *SubscriptionRampIntervalResponseList {
	return &SubscriptionRampIntervalResponseList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type SubscriptionRampIntervalResponseLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []SubscriptionRampIntervalResponse
	HasMore() bool
	Next() string
}

func (list *SubscriptionRampIntervalResponseList) HasMore() bool {
	return list.hasMore
}

func (list *SubscriptionRampIntervalResponseList) Next() string {
	return list.nextPagePath
}

func (list *SubscriptionRampIntervalResponseList) Data() []SubscriptionRampIntervalResponse {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *SubscriptionRampIntervalResponseList) FetchWithContext(ctx context.Context) error {
	resources := &subscriptionRampIntervalResponseList{}
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
func (list *SubscriptionRampIntervalResponseList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *SubscriptionRampIntervalResponseList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &subscriptionRampIntervalResponseList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *SubscriptionRampIntervalResponseList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
