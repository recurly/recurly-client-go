// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type CurrencyCode struct {
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *CurrencyCode) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *CurrencyCode) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type currencyCodeList struct {
	ListMetadata
	Data            []CurrencyCode `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *currencyCodeList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *currencyCodeList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// CurrencyCodeList allows you to paginate CurrencyCode objects
type CurrencyCodeList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []CurrencyCode
}

func NewCurrencyCodeList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *CurrencyCodeList {
	return &CurrencyCodeList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type CurrencyCodeLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []CurrencyCode
	HasMore() bool
	Next() string
}

func (list *CurrencyCodeList) HasMore() bool {
	return list.hasMore
}

func (list *CurrencyCodeList) Next() string {
	return list.nextPagePath
}

func (list *CurrencyCodeList) Data() []CurrencyCode {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *CurrencyCodeList) FetchWithContext(ctx context.Context) error {
	resources := &currencyCodeList{}
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
func (list *CurrencyCodeList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *CurrencyCodeList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &currencyCodeList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *CurrencyCodeList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
