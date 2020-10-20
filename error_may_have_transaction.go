package recurly

import (
	"net/http"
)

type ErrorMayHaveTransaction struct {
	recurlyResponse *ResponseMetadata

	// Type
	Type string `json:"type,omitempty"`

	// Message
	Message string `json:"message,omitempty"`

	// Parameter specific errors
	Params []map[string]interface{} `json:"params,omitempty"`

	// This is only included on errors with `type=transaction`.
	TransactionError TransactionError `json:"transaction_error,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *ErrorMayHaveTransaction) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *ErrorMayHaveTransaction) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type errorMayHaveTransactionList struct {
	ListMetadata
	Data            []ErrorMayHaveTransaction `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *errorMayHaveTransactionList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *errorMayHaveTransactionList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// ErrorMayHaveTransactionList allows you to paginate ErrorMayHaveTransaction objects
type ErrorMayHaveTransactionList struct {
	client       HttpCaller
	nextPagePath string

	HasMore bool
	Data    []ErrorMayHaveTransaction
}

func NewErrorMayHaveTransactionList(client HttpCaller, nextPagePath string) *ErrorMayHaveTransactionList {
	return &ErrorMayHaveTransactionList{
		client:       client,
		nextPagePath: nextPagePath,
		HasMore:      true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *ErrorMayHaveTransactionList) Fetch() error {
	resources := &errorMayHaveTransactionList{}
	err := list.client.Call(http.MethodGet, list.nextPagePath, nil, resources)
	if err != nil {
		return err
	}
	// copy over properties from the response
	list.nextPagePath = resources.Next
	list.HasMore = resources.HasMore
	list.Data = resources.Data
	return nil
}

// Count returns the count of items on the server that match this pager
func (list *ErrorMayHaveTransactionList) Count() (*int64, error) {
	resources := &errorMayHaveTransactionList{}
	err := list.client.Call(http.MethodHead, list.nextPagePath, nil, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}
