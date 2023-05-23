// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type ExternalProductReferenceCollection struct {
	recurlyResponse *ResponseMetadata

	// Will always be List.
	Object string `json:"object,omitempty"`

	// Indicates there are more results on subsequent pages.
	HasMore bool `json:"has_more,omitempty"`

	// Path to subsequent page of results.
	Next string `json:"next,omitempty"`

	Data []ExternalProductReferenceMini `json:"data,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *ExternalProductReferenceCollection) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *ExternalProductReferenceCollection) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type externalProductReferenceCollectionList struct {
	ListMetadata
	Data            []ExternalProductReferenceCollection `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *externalProductReferenceCollectionList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *externalProductReferenceCollectionList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// ExternalProductReferenceCollectionList allows you to paginate ExternalProductReferenceCollection objects
type ExternalProductReferenceCollectionList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []ExternalProductReferenceCollection
}

func NewExternalProductReferenceCollectionList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *ExternalProductReferenceCollectionList {
	return &ExternalProductReferenceCollectionList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type ExternalProductReferenceCollectionLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []ExternalProductReferenceCollection
	HasMore() bool
	Next() string
}

func (list *ExternalProductReferenceCollectionList) HasMore() bool {
	return list.hasMore
}

func (list *ExternalProductReferenceCollectionList) Next() string {
	return list.nextPagePath
}

func (list *ExternalProductReferenceCollectionList) Data() []ExternalProductReferenceCollection {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *ExternalProductReferenceCollectionList) FetchWithContext(ctx context.Context) error {
	resources := &externalProductReferenceCollectionList{}
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
func (list *ExternalProductReferenceCollectionList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *ExternalProductReferenceCollectionList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &externalProductReferenceCollectionList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *ExternalProductReferenceCollectionList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
