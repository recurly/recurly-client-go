// This file is automatically created by Recurly's OpenAPI generation process
// and thus any edits you make by hand will be lost. If you wish to make a
// change to this file, please create a Github issue explaining the changes you
// need and we will usher them to the appropriate places.
package recurly

import (
        "net/http"
        "context"
)

type AccountBalance struct {
  recurlyResponse *ResponseMetadata

  
        // Object type
        Object string `json:"object,omitempty"`

  
        // Account mini details
        Account AccountMini `json:"account,omitempty"`

  
        PastDue bool `json:"past_due,omitempty"`

  
        Balances []AccountBalanceAmount `json:"balances,omitempty"`

  
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *AccountBalance) GetResponse() *ResponseMetadata {
  return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *AccountBalance) setResponse(res *ResponseMetadata) {
  resource.recurlyResponse = res
}





// internal struct for deserializing accounts
type accountBalanceList struct {
	ListMetadata
  Data []AccountBalance `json:"data"`
  recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *accountBalanceList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *accountBalanceList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}




// AccountBalanceList allows you to paginate AccountBalance objects
type AccountBalanceList struct {
  client         HTTPCaller
  requestOptions *RequestOptions
  nextPagePath   string
  hasMore bool
  data    []AccountBalance
}

func NewAccountBalanceList(client HTTPCaller, nextPagePath string, requestOptions *RequestOptions) *AccountBalanceList {
  return &AccountBalanceList{
    client:       client,
    requestOptions: requestOptions,
    nextPagePath: nextPagePath,
    hasMore:      true,
  }
}


type AccountBalanceLister interface {
  Fetch() error
  FetchWithContext(ctx context.Context) error
  Count() (*int64, error)
  CountWithContext(ctx context.Context) (*int64, error)
  Data()    []AccountBalance
  HasMore() bool
  Next() string
}

func (list  *AccountBalanceList) HasMore() bool {
    return list.hasMore
}

func (list  *AccountBalanceList) Next() string {
    return list.nextPagePath
}

func (list *AccountBalanceList) Data() []AccountBalance {
    return list.data
}


// Fetch fetches the next page of data into the `Data` property
func (list *AccountBalanceList) FetchWithContext(ctx context.Context) error {
	resources := &accountBalanceList{}
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
func (list *AccountBalanceList) Fetch() error {
  return list.FetchWithContext(context.Background())
}

// Count returns the count of items on the server that match this pager
func (list *AccountBalanceList) CountWithContext(ctx context.Context) (*int64, error) {
	resources := &accountBalanceList{}
	err := list.client.Call(ctx, http.MethodHead, list.nextPagePath, nil, nil, list.requestOptions, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}

// Count returns the count of items on the server that match this pager
func (list *AccountBalanceList) Count() (*int64, error) {
  return list.CountWithContext(context.Background())
}
