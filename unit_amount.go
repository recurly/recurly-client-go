// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type UnitAmount struct {
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *UnitAmount) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *UnitAmount) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type unitAmountList struct {
	ListMetadata
	Data            []UnitAmount `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *unitAmountList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *unitAmountList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// UnitAmountList allows you to paginate UnitAmount objects
type UnitAmountList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []UnitAmount
}

func NewUnitAmountList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *UnitAmountList {
	return &UnitAmountList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type UnitAmountLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []UnitAmount
	HasMore() bool
	Next() string
}

func (list *UnitAmountList) HasMore() bool {
	return list.hasMore
}

func (list *UnitAmountList) Next() string {
	return list.nextPagePath
}

func (list *UnitAmountList) Data() []UnitAmount {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *UnitAmountList) FetchWithContext(ctx context.Context) error {
	resources := &unitAmountList{}
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
func (list *UnitAmountList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *UnitAmountList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &unitAmountList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *UnitAmountList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
