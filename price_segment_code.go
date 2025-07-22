// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type PriceSegmentCode struct {
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *PriceSegmentCode) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *PriceSegmentCode) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type priceSegmentCodeList struct {
	ListMetadata
	Data            []PriceSegmentCode `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *priceSegmentCodeList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *priceSegmentCodeList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// PriceSegmentCodeList allows you to paginate PriceSegmentCode objects
type PriceSegmentCodeList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []PriceSegmentCode
}

func NewPriceSegmentCodeList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *PriceSegmentCodeList {
	return &PriceSegmentCodeList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type PriceSegmentCodeLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []PriceSegmentCode
	HasMore() bool
	Next() string
}

func (list *PriceSegmentCodeList) HasMore() bool {
	return list.hasMore
}

func (list *PriceSegmentCodeList) Next() string {
	return list.nextPagePath
}

func (list *PriceSegmentCodeList) Data() []PriceSegmentCode {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *PriceSegmentCodeList) FetchWithContext(ctx context.Context) error {
	resources := &priceSegmentCodeList{}
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
func (list *PriceSegmentCodeList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *PriceSegmentCodeList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &priceSegmentCodeList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *PriceSegmentCodeList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
