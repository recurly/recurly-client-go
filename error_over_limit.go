// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type ErrorOverLimit struct {
	recurlyResponse *ResponseMetadata

	// Type
	Type string `json:"type,omitempty"`

	// This credit card has too many cvv check attempts.
	Message string `json:"message,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *ErrorOverLimit) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *ErrorOverLimit) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type errorOverLimitList struct {
	ListMetadata
	Data            []ErrorOverLimit `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *errorOverLimitList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *errorOverLimitList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// ErrorOverLimitList allows you to paginate ErrorOverLimit objects
type ErrorOverLimitList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []ErrorOverLimit
}

func NewErrorOverLimitList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *ErrorOverLimitList {
	return &ErrorOverLimitList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type ErrorOverLimitLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []ErrorOverLimit
	HasMore() bool
	Next() string
}

func (list *ErrorOverLimitList) HasMore() bool {
	return list.hasMore
}

func (list *ErrorOverLimitList) Next() string {
	return list.nextPagePath
}

func (list *ErrorOverLimitList) Data() []ErrorOverLimit {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *ErrorOverLimitList) FetchWithContext(ctx context.Context) error {
	resources := &errorOverLimitList{}
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
func (list *ErrorOverLimitList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *ErrorOverLimitList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &errorOverLimitList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *ErrorOverLimitList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
