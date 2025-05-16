// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
	"time"
)

type PerformanceObligation struct {
	recurlyResponse *ResponseMetadata

	// The ID of a performance obligation. Performance obligations are
	// only accessible as a part of the Recurly RevRec Standard and
	// Recurly RevRec Advanced features.
	Id string `json:"id,omitempty"`

	// Performance Obligation Name
	Name string `json:"name,omitempty"`

	// Created At
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// Last updated at
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *PerformanceObligation) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *PerformanceObligation) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type performanceObligationList struct {
	ListMetadata
	Data            []PerformanceObligation `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *performanceObligationList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *performanceObligationList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// PerformanceObligationList allows you to paginate PerformanceObligation objects
type PerformanceObligationList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []PerformanceObligation
}

func NewPerformanceObligationList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *PerformanceObligationList {
	return &PerformanceObligationList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type PerformanceObligationLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []PerformanceObligation
	HasMore() bool
	Next() string
}

func (list *PerformanceObligationList) HasMore() bool {
	return list.hasMore
}

func (list *PerformanceObligationList) Next() string {
	return list.nextPagePath
}

func (list *PerformanceObligationList) Data() []PerformanceObligation {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *PerformanceObligationList) FetchWithContext(ctx context.Context) error {
	resources := &performanceObligationList{}
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
func (list *PerformanceObligationList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *PerformanceObligationList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &performanceObligationList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *PerformanceObligationList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
