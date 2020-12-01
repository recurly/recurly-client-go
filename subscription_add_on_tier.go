// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type SubscriptionAddOnTier struct {
	recurlyResponse *ResponseMetadata

	// Ending quantity
	EndingQuantity int `json:"ending_quantity,omitempty"`

	// Unit amount
	UnitAmount float64 `json:"unit_amount,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *SubscriptionAddOnTier) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *SubscriptionAddOnTier) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type subscriptionAddOnTierList struct {
	ListMetadata
	Data            []SubscriptionAddOnTier `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *subscriptionAddOnTierList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *subscriptionAddOnTierList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// SubscriptionAddOnTierList allows you to paginate SubscriptionAddOnTier objects
type SubscriptionAddOnTierList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string

	HasMore bool
	Data    []SubscriptionAddOnTier
}

func NewSubscriptionAddOnTierList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *SubscriptionAddOnTierList {
	return &SubscriptionAddOnTierList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		HasMore:        true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *SubscriptionAddOnTierList) FetchWithContext(ctx context.Context) error {
	resources := &subscriptionAddOnTierList{}
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
func (list *SubscriptionAddOnTierList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *SubscriptionAddOnTierList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &subscriptionAddOnTierList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *SubscriptionAddOnTierList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
