// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type DunningInterval struct {
	recurlyResponse *ResponseMetadata

	// Number of days before sending the next email.
	Days int `json:"days,omitempty"`

	// Email template being used.
	EmailTemplate string `json:"email_template,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *DunningInterval) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *DunningInterval) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type dunningIntervalList struct {
	ListMetadata
	Data            []DunningInterval `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *dunningIntervalList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *dunningIntervalList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// DunningIntervalList allows you to paginate DunningInterval objects
type DunningIntervalList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []DunningInterval
}

func NewDunningIntervalList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *DunningIntervalList {
	return &DunningIntervalList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type DunningIntervalLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []DunningInterval
	HasMore() bool
	Next() string
}

func (list *DunningIntervalList) HasMore() bool {
	return list.hasMore
}

func (list *DunningIntervalList) Next() string {
	return list.nextPagePath
}

func (list *DunningIntervalList) Data() []DunningInterval {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *DunningIntervalList) FetchWithContext(ctx context.Context) error {
	resources := &dunningIntervalList{}
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
func (list *DunningIntervalList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *DunningIntervalList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &dunningIntervalList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *DunningIntervalList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
