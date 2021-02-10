// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type AccountBalanceAmount struct {
	recurlyResponse *ResponseMetadata

	// 3-letter ISO 4217 currency code.
	Currency string `json:"currency,omitempty"`

	// Total amount the account is past due.
	Amount float64 `json:"amount,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *AccountBalanceAmount) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *AccountBalanceAmount) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type accountBalanceAmountList struct {
	ListMetadata
	Data            []AccountBalanceAmount `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *accountBalanceAmountList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *accountBalanceAmountList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// AccountBalanceAmountList allows you to paginate AccountBalanceAmount objects
type AccountBalanceAmountList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []AccountBalanceAmount
}

type AccountBalanceAmountLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []AccountBalanceAmount
	HasMore() bool
	Next() string
}

func NewAccountBalanceAmountList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *AccountBalanceAmountList {
	return &AccountBalanceAmountList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

func (list *AccountBalanceAmountList) HasMore() bool {
	return list.hasMore
}

func (list *AccountBalanceAmountList) Next() string {
	return list.nextPagePath
}

func (list *AccountBalanceAmountList) Data() []AccountBalanceAmount {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *AccountBalanceAmountList) FetchWithContext(ctx context.Context) error {
	resources := &accountBalanceAmountList{}
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
func (list *AccountBalanceAmountList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *AccountBalanceAmountList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &accountBalanceAmountList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *AccountBalanceAmountList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
