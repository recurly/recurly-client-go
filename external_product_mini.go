// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type ExternalProductMini struct {
	recurlyResponse *ResponseMetadata

	// External product ID
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// The merchant-facing name of the external product.
	Name string `json:"name,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *ExternalProductMini) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *ExternalProductMini) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type externalProductMiniList struct {
	ListMetadata
	Data            []ExternalProductMini `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *externalProductMiniList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *externalProductMiniList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// ExternalProductMiniList allows you to paginate ExternalProductMini objects
type ExternalProductMiniList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []ExternalProductMini
}

func NewExternalProductMiniList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *ExternalProductMiniList {
	return &ExternalProductMiniList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type ExternalProductMiniLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []ExternalProductMini
	HasMore() bool
	Next() string
}

func (list *ExternalProductMiniList) HasMore() bool {
	return list.hasMore
}

func (list *ExternalProductMiniList) Next() string {
	return list.nextPagePath
}

func (list *ExternalProductMiniList) Data() []ExternalProductMini {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *ExternalProductMiniList) FetchWithContext(ctx context.Context) error {
	resources := &externalProductMiniList{}
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
func (list *ExternalProductMiniList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *ExternalProductMiniList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &externalProductMiniList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *ExternalProductMiniList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
