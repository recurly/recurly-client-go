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

type ExternalAccount struct {
	recurlyResponse *ResponseMetadata

	Object string `json:"object,omitempty"`

	// UUID of the external_account .
	Id string `json:"id,omitempty"`

	// Represents the account code for the external account.
	ExternalAccountCode string `json:"external_account_code,omitempty"`

	// Represents the connection type. `AppleAppStore` or `GooglePlayStore`
	ExternalConnectionType string `json:"external_connection_type,omitempty"`

	// Created at
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Last updated at
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *ExternalAccount) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *ExternalAccount) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type externalAccountList struct {
	ListMetadata
	Data            []ExternalAccount `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *externalAccountList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *externalAccountList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// ExternalAccountList allows you to paginate ExternalAccount objects
type ExternalAccountList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []ExternalAccount
}

func NewExternalAccountList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *ExternalAccountList {
	return &ExternalAccountList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type ExternalAccountLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []ExternalAccount
	HasMore() bool
	Next() string
}

func (list *ExternalAccountList) HasMore() bool {
	return list.hasMore
}

func (list *ExternalAccountList) Next() string {
	return list.nextPagePath
}

func (list *ExternalAccountList) Data() []ExternalAccount {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *ExternalAccountList) FetchWithContext(ctx context.Context) error {
	resources := &externalAccountList{}
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
func (list *ExternalAccountList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *ExternalAccountList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &externalAccountList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *ExternalAccountList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
