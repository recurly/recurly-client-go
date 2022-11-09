// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type ExternalResourceMini struct {
	recurlyResponse *ResponseMetadata

	// System-generated unique identifier for an external resource ID, e.g. `e28zov4fw0v2`.
	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	// Identifier or URL reference where the resource is canonically available in the external platform.
	ExternalObjectReference string `json:"external_object_reference,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *ExternalResourceMini) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *ExternalResourceMini) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type externalResourceMiniList struct {
	ListMetadata
	Data            []ExternalResourceMini `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *externalResourceMiniList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *externalResourceMiniList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// ExternalResourceMiniList allows you to paginate ExternalResourceMini objects
type ExternalResourceMiniList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []ExternalResourceMini
}

func NewExternalResourceMiniList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *ExternalResourceMiniList {
	return &ExternalResourceMiniList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type ExternalResourceMiniLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []ExternalResourceMini
	HasMore() bool
	Next() string
}

func (list *ExternalResourceMiniList) HasMore() bool {
	return list.hasMore
}

func (list *ExternalResourceMiniList) Next() string {
	return list.nextPagePath
}

func (list *ExternalResourceMiniList) Data() []ExternalResourceMini {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *ExternalResourceMiniList) FetchWithContext(ctx context.Context) error {
	resources := &externalResourceMiniList{}
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
func (list *ExternalResourceMiniList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *ExternalResourceMiniList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &externalResourceMiniList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *ExternalResourceMiniList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
