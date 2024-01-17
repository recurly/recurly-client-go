// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
	"context"
	"net/http"
)

type TransactionFraudInfo struct {
	recurlyResponse *ResponseMetadata

	// Object type
	Object string `json:"object,omitempty"`

	// Kount score
	Score int `json:"score,omitempty"`

	// Kount decision
	Decision string `json:"decision,omitempty"`

	// Kount transaction reference ID
	Reference string `json:"reference,omitempty"`

	// A list of fraud risk rules that were triggered for the transaction.
	RiskRulesTriggered []FraudRiskRule `json:"risk_rules_triggered,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *TransactionFraudInfo) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *TransactionFraudInfo) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type transactionFraudInfoList struct {
	ListMetadata
	Data            []TransactionFraudInfo `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *transactionFraudInfoList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *transactionFraudInfoList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// TransactionFraudInfoList allows you to paginate TransactionFraudInfo objects
type TransactionFraudInfoList struct {
	client         HTTPCaller
	requestOptions *RequestOptions
	nextPagePath   string
	hasMore        bool
	data           []TransactionFraudInfo
}

func NewTransactionFraudInfoList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *TransactionFraudInfoList {
	return &TransactionFraudInfoList{
		client:         client,
		requestOptions: requestOptions,
		nextPagePath:   nextPagePath,
		hasMore:        true,
	}
}

type TransactionFraudInfoLister interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Count() (*int64, error)
	CountWithContext(ctx context.Context) (*int64, error)
	Data() []TransactionFraudInfo
	HasMore() bool
	Next() string
}

func (list *TransactionFraudInfoList) HasMore() bool {
	return list.hasMore
}

func (list *TransactionFraudInfoList) Next() string {
	return list.nextPagePath
}

func (list *TransactionFraudInfoList) Data() []TransactionFraudInfo {
	return list.data
}

// Fetch fetches the next page of data into the `Data` property
func (list *TransactionFraudInfoList) FetchWithContext(ctx context.Context) error {
	resources := &transactionFraudInfoList{}
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
func (list *TransactionFraudInfoList) Fetch() error {
	return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *TransactionFraudInfoList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &transactionFraudInfoList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *TransactionFraudInfoList) Count() (*int64, error) {
	return list.CountWithContext(context.Background())
}
