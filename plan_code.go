// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type PlanCode struct {
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *PlanCode) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *PlanCode) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type planCodeList struct {
	ListMetadata
	Data            []PlanCode `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *planCodeList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *planCodeList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// PlanCodeList allows you to paginate PlanCode objects
type PlanCodeList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []PlanCode
}

func NewPlanCodeList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *PlanCodeList {
	return &PlanCodeList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type PlanCodeLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []PlanCode
	HasMore() bool
	Next() string
}

func (list *PlanCodeList) HasMore() bool {
	return list.hasMore
}

func (list *PlanCodeList) Next() string {
	return list.nextPagePath
}

func (list *PlanCodeList) Data() []PlanCode {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *PlanCodeList) FetchWithContext(ctx context.Context) error {
	resources := &planCodeList{}
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
func (list *PlanCodeList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *PlanCodeList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &planCodeList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *PlanCodeList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
