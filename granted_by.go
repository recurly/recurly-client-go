// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type GrantedBy struct {
	recurlyResponse *ResponseMetadata

	// Object Type
	Object string `json:"object,omitempty"`

	// The ID of the subscription or external subscription that grants the permission to the account.
	Id string `json:"id,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *GrantedBy) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *GrantedBy) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type grantedByList struct {
	ListMetadata
	Data            []GrantedBy `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *grantedByList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *grantedByList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// GrantedByList allows you to paginate GrantedBy objects
type GrantedByList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []GrantedBy
}

func NewGrantedByList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *GrantedByList {
	return &GrantedByList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type GrantedByLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []GrantedBy
	HasMore() bool
	Next() string
}

func (list *GrantedByList) HasMore() bool {
	return list.hasMore
}

func (list *GrantedByList) Next() string {
	return list.nextPagePath
}

func (list *GrantedByList) Data() []GrantedBy {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *GrantedByList) FetchWithContext(ctx context.Context) error {
	resources := &grantedByList{}
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
func (list *GrantedByList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *GrantedByList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &grantedByList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *GrantedByList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
