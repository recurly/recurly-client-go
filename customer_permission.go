// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type CustomerPermission struct {
	recurlyResponse *ResponseMetadata

	// Customer permission ID.
	Id string `json:"id,omitempty"`

	// Customer permission code.
	Code string `json:"code,omitempty"`

	// Customer permission name.
	Name string `json:"name,omitempty"`

	// Description of customer permission.
	Description string `json:"description,omitempty"`

	// It will always be "customer_permission".
	Object string `json:"object,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *CustomerPermission) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *CustomerPermission) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type customerPermissionList struct {
	ListMetadata
	Data            []CustomerPermission `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *customerPermissionList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *customerPermissionList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// CustomerPermissionList allows you to paginate CustomerPermission objects
type CustomerPermissionList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []CustomerPermission
}

func NewCustomerPermissionList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *CustomerPermissionList {
	return &CustomerPermissionList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type CustomerPermissionLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []CustomerPermission
	HasMore() bool
	Next() string
}

func (list *CustomerPermissionList) HasMore() bool {
	return list.hasMore
}

func (list *CustomerPermissionList) Next() string {
	return list.nextPagePath
}

func (list *CustomerPermissionList) Data() []CustomerPermission {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *CustomerPermissionList) FetchWithContext(ctx context.Context) error {
	resources := &customerPermissionList{}
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
func (list *CustomerPermissionList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *CustomerPermissionList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &customerPermissionList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *CustomerPermissionList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
