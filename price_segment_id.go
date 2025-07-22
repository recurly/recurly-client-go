// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type PriceSegmentId struct {
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *PriceSegmentId) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *PriceSegmentId) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type priceSegmentIdList struct {
	ListMetadata
	Data            []PriceSegmentId `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *priceSegmentIdList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *priceSegmentIdList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// PriceSegmentIdList allows you to paginate PriceSegmentId objects
type PriceSegmentIdList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []PriceSegmentId
}

func NewPriceSegmentIdList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *PriceSegmentIdList {
	return &PriceSegmentIdList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type PriceSegmentIdLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []PriceSegmentId
	HasMore() bool
	Next() string
}

func (list *PriceSegmentIdList) HasMore() bool {
	return list.hasMore
}

func (list *PriceSegmentIdList) Next() string {
	return list.nextPagePath
}

func (list *PriceSegmentIdList) Data() []PriceSegmentId {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *PriceSegmentIdList) FetchWithContext(ctx context.Context) error {
	resources := &priceSegmentIdList{}
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
func (list *PriceSegmentIdList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *PriceSegmentIdList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &priceSegmentIdList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *PriceSegmentIdList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
