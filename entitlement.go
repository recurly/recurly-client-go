// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
	"time"
)

type Entitlement struct {
	recurlyResponse *ResponseMetadata

	// Entitlement
	Object string `json:"object,omitempty"`

	CustomerPermission CustomerPermission `json:"customer_permission,omitempty"`

	// Subscription or item that granted the customer permission.
	GrantedBy []GrantedBy `json:"granted_by,omitempty"`

	// Time object was created.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Time the object was last updated
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *Entitlement) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *Entitlement) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type entitlementList struct {
	ListMetadata
	Data            []Entitlement `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *entitlementList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *entitlementList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// EntitlementList allows you to paginate Entitlement objects
type EntitlementList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []Entitlement
}

func NewEntitlementList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *EntitlementList {
	return &EntitlementList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type EntitlementLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []Entitlement
	HasMore() bool
	Next() string
}

func (list *EntitlementList) HasMore() bool {
	return list.hasMore
}

func (list *EntitlementList) Next() string {
	return list.nextPagePath
}

func (list *EntitlementList) Data() []Entitlement {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *EntitlementList) FetchWithContext(ctx context.Context) error {
	resources := &entitlementList{}
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
func (list *EntitlementList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *EntitlementList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &entitlementList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *EntitlementList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
