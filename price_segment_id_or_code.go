// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type PriceSegmentIdOrCode struct {
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *PriceSegmentIdOrCode) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *PriceSegmentIdOrCode) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type priceSegmentIdOrCodeList struct {
	ListMetadata
	Data            []PriceSegmentIdOrCode `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *priceSegmentIdOrCodeList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *priceSegmentIdOrCodeList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// PriceSegmentIdOrCodeList allows you to paginate PriceSegmentIdOrCode objects
type PriceSegmentIdOrCodeList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []PriceSegmentIdOrCode
}

func NewPriceSegmentIdOrCodeList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *PriceSegmentIdOrCodeList {
	return &PriceSegmentIdOrCodeList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type PriceSegmentIdOrCodeLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []PriceSegmentIdOrCode
	HasMore() bool
	Next() string
}

func (list *PriceSegmentIdOrCodeList) HasMore() bool {
	return list.hasMore
}

func (list *PriceSegmentIdOrCodeList) Next() string {
	return list.nextPagePath
}

func (list *PriceSegmentIdOrCodeList) Data() []PriceSegmentIdOrCode {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *PriceSegmentIdOrCodeList) FetchWithContext(ctx context.Context) error {
	resources := &priceSegmentIdOrCodeList{}
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
func (list *PriceSegmentIdOrCodeList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *PriceSegmentIdOrCodeList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &priceSegmentIdOrCodeList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *PriceSegmentIdOrCodeList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
