// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type PerformanceObligationId struct {
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *PerformanceObligationId) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *PerformanceObligationId) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type performanceObligationIdList struct {
	ListMetadata
	Data            []PerformanceObligationId `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *performanceObligationIdList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *performanceObligationIdList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// PerformanceObligationIdList allows you to paginate PerformanceObligationId objects
type PerformanceObligationIdList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []PerformanceObligationId
}

func NewPerformanceObligationIdList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *PerformanceObligationIdList {
	return &PerformanceObligationIdList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type PerformanceObligationIdLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []PerformanceObligationId
	HasMore() bool
	Next() string
}

func (list *PerformanceObligationIdList) HasMore() bool {
	return list.hasMore
}

func (list *PerformanceObligationIdList) Next() string {
	return list.nextPagePath
}

func (list *PerformanceObligationIdList) Data() []PerformanceObligationId {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *PerformanceObligationIdList) FetchWithContext(ctx context.Context) error {
	resources := &performanceObligationIdList{}
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
func (list *PerformanceObligationIdList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *PerformanceObligationIdList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &performanceObligationIdList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *PerformanceObligationIdList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
