// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type BillingInfoUpdatedBy struct {
	recurlyResponse *ResponseMetadata

	// Customer's IP address when updating their billing information.
	Ip string `json:"ip,omitempty"`

	// Country of IP address, if known by Recurly.
	Country string `json:"country,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *BillingInfoUpdatedBy) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *BillingInfoUpdatedBy) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type billingInfoUpdatedByList struct {
	ListMetadata
	Data            []BillingInfoUpdatedBy `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *billingInfoUpdatedByList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *billingInfoUpdatedByList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// BillingInfoUpdatedByList allows you to paginate BillingInfoUpdatedBy objects
type BillingInfoUpdatedByList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []BillingInfoUpdatedBy
}

func NewBillingInfoUpdatedByList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *BillingInfoUpdatedByList {
	return &BillingInfoUpdatedByList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type BillingInfoUpdatedByLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []BillingInfoUpdatedBy
	HasMore() bool
	Next() string
}

func (list *BillingInfoUpdatedByList) HasMore() bool {
	return list.hasMore
}

func (list *BillingInfoUpdatedByList) Next() string {
	return list.nextPagePath
}

func (list *BillingInfoUpdatedByList) Data() []BillingInfoUpdatedBy {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *BillingInfoUpdatedByList) FetchWithContext(ctx context.Context) error {
	resources := &billingInfoUpdatedByList{}
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
func (list *BillingInfoUpdatedByList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *BillingInfoUpdatedByList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &billingInfoUpdatedByList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *BillingInfoUpdatedByList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
