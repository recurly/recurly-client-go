// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type SubscriptionShipping struct {
	recurlyResponse *ResponseMetadata

	// Object type
	Object string `json:"object,omitempty"`

	Address ShippingAddress `json:"address,omitempty"`

	Method ShippingMethodMini `json:"method,omitempty"`

	// Subscription's shipping cost
	Amount float64 `json:"amount,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *SubscriptionShipping) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *SubscriptionShipping) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type subscriptionShippingList struct {
	ListMetadata
	Data            []SubscriptionShipping `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *subscriptionShippingList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *subscriptionShippingList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// SubscriptionShippingList allows you to paginate SubscriptionShipping objects
type SubscriptionShippingList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string

	HasMore bool
	Data    []SubscriptionShipping
}

func NewSubscriptionShippingList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *SubscriptionShippingList {
	return &SubscriptionShippingList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		HasMore:        true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *SubscriptionShippingList) FetchWithContext(ctx context.Context) error {
	resources := &subscriptionShippingList{}
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
func (list *SubscriptionShippingList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *SubscriptionShippingList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &subscriptionShippingList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *SubscriptionShippingList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
