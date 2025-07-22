// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type PlanName struct {
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *PlanName) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *PlanName) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type planNameList struct {
	ListMetadata
	Data            []PlanName `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *planNameList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *planNameList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// PlanNameList allows you to paginate PlanName objects
type PlanNameList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []PlanName
}

func NewPlanNameList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *PlanNameList {
	return &PlanNameList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type PlanNameLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []PlanName
	HasMore() bool
	Next() string
}

func (list *PlanNameList) HasMore() bool {
	return list.hasMore
}

func (list *PlanNameList) Next() string {
	return list.nextPagePath
}

func (list *PlanNameList) Data() []PlanName {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *PlanNameList) FetchWithContext(ctx context.Context) error {
	resources := &planNameList{}
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
func (list *PlanNameList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *PlanNameList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &planNameList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *PlanNameList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
