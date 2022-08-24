// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type ErrorMayHaveCVV struct {
	recurlyResponse *ResponseMetadata

	// Type
	Type string `json:"type,omitempty"`

	// The security code you entered does not match. Please update the CVV and try again.
	Message string `json:"message,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *ErrorMayHaveCVV) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *ErrorMayHaveCVV) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type errorMayHaveCVVList struct {
	ListMetadata
	Data            []ErrorMayHaveCVV `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *errorMayHaveCVVList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *errorMayHaveCVVList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// ErrorMayHaveCVVList allows you to paginate ErrorMayHaveCVV objects
type ErrorMayHaveCVVList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []ErrorMayHaveCVV
}

func NewErrorMayHaveCVVList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *ErrorMayHaveCVVList {
	return &ErrorMayHaveCVVList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type ErrorMayHaveCVVLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []ErrorMayHaveCVV
	HasMore() bool
	Next() string
}

func (list *ErrorMayHaveCVVList) HasMore() bool {
	return list.hasMore
}

func (list *ErrorMayHaveCVVList) Next() string {
	return list.nextPagePath
}

func (list *ErrorMayHaveCVVList) Data() []ErrorMayHaveCVV {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *ErrorMayHaveCVVList) FetchWithContext(ctx context.Context) error {
	resources := &errorMayHaveCVVList{}
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
func (list *ErrorMayHaveCVVList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *ErrorMayHaveCVVList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &errorMayHaveCVVList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *ErrorMayHaveCVVList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
