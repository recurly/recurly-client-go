// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type PriceSegment struct {
	recurlyResponse *ResponseMetadata

	// Object type
	Object string `json:"object,omitempty"`

	// The price segment ID, e.g. `e28zov4fw0v2`.
	Id string `json:"id,omitempty"`

	// The price segment code, e.g. `my-price-segment`.
	Code string `json:"code,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *PriceSegment) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *PriceSegment) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type priceSegmentList struct {
	ListMetadata
	Data            []PriceSegment `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *priceSegmentList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *priceSegmentList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// PriceSegmentList allows you to paginate PriceSegment objects
type PriceSegmentList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []PriceSegment
}

func NewPriceSegmentList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *PriceSegmentList {
	return &PriceSegmentList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type PriceSegmentLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []PriceSegment
	HasMore() bool
	Next() string
}

func (list *PriceSegmentList) HasMore() bool {
	return list.hasMore
}

func (list *PriceSegmentList) Next() string {
	return list.nextPagePath
}

func (list *PriceSegmentList) Data() []PriceSegment {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *PriceSegmentList) FetchWithContext(ctx context.Context) error {
	resources := &priceSegmentList{}
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
func (list *PriceSegmentList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *PriceSegmentList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &priceSegmentList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *PriceSegmentList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
