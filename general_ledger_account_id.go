// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type GeneralLedgerAccountId struct {
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *GeneralLedgerAccountId) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *GeneralLedgerAccountId) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type generalLedgerAccountIdList struct {
	ListMetadata
	Data            []GeneralLedgerAccountId `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *generalLedgerAccountIdList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *generalLedgerAccountIdList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// GeneralLedgerAccountIdList allows you to paginate GeneralLedgerAccountId objects
type GeneralLedgerAccountIdList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []GeneralLedgerAccountId
}

func NewGeneralLedgerAccountIdList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *GeneralLedgerAccountIdList {
	return &GeneralLedgerAccountIdList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type GeneralLedgerAccountIdLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []GeneralLedgerAccountId
	HasMore() bool
	Next() string
}

func (list *GeneralLedgerAccountIdList) HasMore() bool {
	return list.hasMore
}

func (list *GeneralLedgerAccountIdList) Next() string {
	return list.nextPagePath
}

func (list *GeneralLedgerAccountIdList) Data() []GeneralLedgerAccountId {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *GeneralLedgerAccountIdList) FetchWithContext(ctx context.Context) error {
	resources := &generalLedgerAccountIdList{}
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
func (list *GeneralLedgerAccountIdList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *GeneralLedgerAccountIdList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &generalLedgerAccountIdList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *GeneralLedgerAccountIdList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
