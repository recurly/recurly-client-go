// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type TransactionNextAction struct {
	recurlyResponse *ResponseMetadata

	// The type of next action required.
	Type string `json:"type,omitempty"`

	// The value associated with the next action type.
	Value string `json:"value,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *TransactionNextAction) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *TransactionNextAction) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type transactionNextActionList struct {
	ListMetadata
	Data            []TransactionNextAction `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *transactionNextActionList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *transactionNextActionList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// TransactionNextActionList allows you to paginate TransactionNextAction objects
type TransactionNextActionList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []TransactionNextAction
}

func NewTransactionNextActionList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *TransactionNextActionList {
	return &TransactionNextActionList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type TransactionNextActionLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []TransactionNextAction
	HasMore() bool
	Next() string
}

func (list *TransactionNextActionList) HasMore() bool {
	return list.hasMore
}

func (list *TransactionNextActionList) Next() string {
	return list.nextPagePath
}

func (list *TransactionNextActionList) Data() []TransactionNextAction {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *TransactionNextActionList) FetchWithContext(ctx context.Context) error {
	resources := &transactionNextActionList{}
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
func (list *TransactionNextActionList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *TransactionNextActionList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &transactionNextActionList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *TransactionNextActionList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
