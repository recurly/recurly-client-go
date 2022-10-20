// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
        "net/http"
        "context"
)

type TransactionPaymentGateway struct {
  recurlyResponse *ResponseMetadata

  
        Id string `json:"id,omitempty"`

  
        // Object type
        Object string `json:"object,omitempty"`

  
        Type string `json:"type,omitempty"`

  
        Name string `json:"name,omitempty"`

  
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *TransactionPaymentGateway) GetResponse() *ResponseMetadata {
  return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *TransactionPaymentGateway) setResponse(res *ResponseMetadata) {
  resource.recurlyResponse = res
}





// internal struct for deserializing accounts
type transactionPaymentGatewayList struct {
	ListMetadata
  Data []TransactionPaymentGateway `json:"data"`
  recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *transactionPaymentGatewayList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *transactionPaymentGatewayList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}




// TransactionPaymentGatewayList allows you to paginate TransactionPaymentGateway objects
type TransactionPaymentGatewayList struct {
  client         HTTPCaller
  requestOptions *RequestOptions
  nextPagePath   string
  hasMore bool
  data    []TransactionPaymentGateway
}

func NewTransactionPaymentGatewayList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *TransactionPaymentGatewayList {
  return &TransactionPaymentGatewayList{
    client:       client,
    requestOptions: requestOptions,
    nextPagePath: nextPagePath,
    hasMore:      true,
  }
}


type TransactionPaymentGatewayLister interface {
  Fetch() error
  FetchWithContext(ctx context.Context) error
  Count() (*int64, error)
  CountWithContext(ctx context.Context) (*int64, error)
  Data()    []TransactionPaymentGateway
  HasMore() bool
  Next() string
}

func (list  *TransactionPaymentGatewayList) HasMore() bool {
    return list.hasMore
}

func (list  *TransactionPaymentGatewayList) Next() string {
    return list.nextPagePath
}

func (list *TransactionPaymentGatewayList) Data() []TransactionPaymentGateway {
    return list.data
}


// Fetch fetches the next page of data into the `Data` property
func (list *TransactionPaymentGatewayList) FetchWithContext(ctx context.Context) error {
	resources := &transactionPaymentGatewayList{}
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
func (list *TransactionPaymentGatewayList) Fetch() error {
  return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *TransactionPaymentGatewayList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &transactionPaymentGatewayList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *TransactionPaymentGatewayList) Count() (*int64, error) {
  return list.CountWithContext(context.Background())
}
