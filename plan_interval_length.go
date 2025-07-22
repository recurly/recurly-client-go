// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type PlanIntervalLength struct {
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *PlanIntervalLength) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *PlanIntervalLength) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type planIntervalLengthList struct {
	ListMetadata
	Data            []PlanIntervalLength `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *planIntervalLengthList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *planIntervalLengthList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// PlanIntervalLengthList allows you to paginate PlanIntervalLength objects
type PlanIntervalLengthList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []PlanIntervalLength
}

func NewPlanIntervalLengthList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *PlanIntervalLengthList {
	return &PlanIntervalLengthList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type PlanIntervalLengthLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []PlanIntervalLength
	HasMore() bool
	Next() string
}

func (list *PlanIntervalLengthList) HasMore() bool {
	return list.hasMore
}

func (list *PlanIntervalLengthList) Next() string {
	return list.nextPagePath
}

func (list *PlanIntervalLengthList) Data() []PlanIntervalLength {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *PlanIntervalLengthList) FetchWithContext(ctx context.Context) error {
	resources := &planIntervalLengthList{}
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
func (list *PlanIntervalLengthList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *PlanIntervalLengthList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &planIntervalLengthList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *PlanIntervalLengthList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
