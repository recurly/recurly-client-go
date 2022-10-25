// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type Entitlements struct {
	recurlyResponse *ResponseMetadata

	// Object Type
	Object string `json:"object,omitempty"`

	// Indicates there are more results on subsequent pages.
	HasMore bool `json:"has_more,omitempty"`

	// Path to subsequent page of results.
	Next string `json:"next,omitempty"`

	Data []Entitlement `json:"data,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *Entitlements) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *Entitlements) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type entitlementsList struct {
	ListMetadata
	Data            []Entitlements `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *entitlementsList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *entitlementsList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// EntitlementsList allows you to paginate Entitlements objects
type EntitlementsList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []Entitlements
}

func NewEntitlementsList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *EntitlementsList {
	return &EntitlementsList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type EntitlementsLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []Entitlements
	HasMore() bool
	Next() string
}

func (list *EntitlementsList) HasMore() bool {
	return list.hasMore
}

func (list *EntitlementsList) Next() string {
	return list.nextPagePath
}

func (list *EntitlementsList) Data() []Entitlements {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *EntitlementsList) FetchWithContext(ctx context.Context) error {
	resources := &entitlementsList{}
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
func (list *EntitlementsList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *EntitlementsList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &entitlementsList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *EntitlementsList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
